package main

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
)

type modelFunc func(context.Context, []Message) (Reply, error)

func (f modelFunc) Next(ctx context.Context, h []Message) (Reply, error) { return f(ctx, h) }

// 若循环不把工具结果交回模型，正常任务无法完成。
func TestRunReadsThenAnswers(t *testing.T) {
	answer, h, err := Run(context.Background(), ScriptedModel{}, "下一步练什么", Config{3, 2})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(answer, "多返回值") {
		t.Fatalf("回答缺少资料依据: %q", answer)
	}
	var ids []string
	for _, m := range h {
		if m.Role == "tool" {
			if m.IsError {
				t.Fatal(m.Text)
			}
			ids = append(ids, m.CallID)
		}
	}
	if !reflect.DeepEqual(ids, []string{"c1", "c2"}) {
		t.Fatalf("工具配对错误: %v", ids)
	}
	if len(h) != 6 {
		t.Fatalf("历史不完整: %d", len(h))
	}
}
func TestUnknownToolBecomesVisibleErrorResult(t *testing.T) {
	model := modelFunc(func(ctx context.Context, h []Message) (Reply, error) {
		if len(h) == 1 {
			return Reply{Call: &Call{"x", "shell", "{}"}}, nil
		}
		last := h[len(h)-1]
		if last.Role != "tool" || !last.IsError || last.CallID != "x" || !strings.Contains(last.Text, "未知工具") {
			t.Fatalf("模型未收到错误结果: %+v", last)
		}
		return Reply{Text: "无法使用该工具"}, nil
	})
	_, _, err := Run(context.Background(), model, "test", Config{2, 1})
	if err != nil {
		t.Fatal(err)
	}
}
func TestReadNoteRejectsInvalidRequests(t *testing.T) {
	cases := []string{`{"name":"../secret"}`, `{"name":"other.md"}`, `{"name":""}`, `{"name":3}`, `{"name":"functions.md","admin":true}`, `{"name":"functions.md"} {}`, `null`, `[]`, `{`}
	for _, args := range cases {
		t.Run(args, func(t *testing.T) {
			if _, err := executeTool(context.Background(), Call{"x", "read_note", args}); err == nil {
				t.Fatalf("接受了无效输入: %s", args)
			}
		})
	}
}
func TestListNotesIsStableAndStrict(t *testing.T) {
	got, err := executeTool(context.Background(), Call{"x", "list_notes", "{}"})
	if err != nil {
		t.Fatal(err)
	}
	if got != `["functions.md","variables.md"]` {
		t.Fatalf("列表: %q", got)
	}
	if _, err = executeTool(context.Background(), Call{"x", "list_notes", `{"extra":1}`}); err == nil {
		t.Fatal("未拒绝多余参数")
	}
}
func TestModelStepBudget(t *testing.T) {
	_, h, err := Run(context.Background(), ScriptedModel{}, "test", Config{2, 2})
	if !errors.Is(err, ErrStepLimit) {
		t.Fatalf("错误: %v", err)
	}
	if len(h) != 5 {
		t.Fatalf("应保留已完成过程: %d", len(h))
	}
}
func TestToolBudgetBlocksNextDispatch(t *testing.T) {
	_, h, err := Run(context.Background(), ScriptedModel{}, "test", Config{3, 1})
	if !errors.Is(err, ErrToolLimit) {
		t.Fatalf("错误: %v", err)
	}
	count := 0
	for _, m := range h {
		if m.Role == "tool" {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("越过工具预算: %d", count)
	}
}
func TestCancellationPreventsFirstModelRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := 0
	model := modelFunc(func(context.Context, []Message) (Reply, error) { calls++; return Reply{Text: "late"}, nil })
	_, _, err := Run(ctx, model, "test", Config{3, 2})
	if !errors.Is(err, context.Canceled) || calls != 0 {
		t.Fatalf("err=%v calls=%d", err, calls)
	}
}
func TestLateModelReplyDoesNotOverrideCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	model := modelFunc(func(context.Context, []Message) (Reply, error) { cancel(); return Reply{Text: "late"}, nil })
	answer, _, err := Run(ctx, model, "test", Config{3, 2})
	if !errors.Is(err, context.Canceled) || answer != "" {
		t.Fatalf("answer=%q err=%v", answer, err)
	}
}
func TestModelFailureIsNotSuccess(t *testing.T) {
	boom := errors.New("provider unavailable")
	model := modelFunc(func(context.Context, []Message) (Reply, error) { return Reply{}, boom })
	answer, _, err := Run(context.Background(), model, "test", Config{3, 2})
	if !errors.Is(err, boom) || answer != "" {
		t.Fatalf("answer=%q err=%v", answer, err)
	}
}
func TestModelCannotMutateRecordedHistory(t *testing.T) {
	model := modelFunc(func(ctx context.Context, h []Message) (Reply, error) {
		h[0].Text = "rewritten"
		return Reply{Text: "ok"}, nil
	})
	_, h, err := Run(context.Background(), model, "original", Config{1, 0})
	if err != nil {
		t.Fatal(err)
	}
	if h[0].Text != "original" {
		t.Fatalf("历史被修改: %+v", h)
	}
}
func TestRejectsMalformedModelReplyAndCallID(t *testing.T) {
	cases := []Reply{{}, {Text: "both", Call: &Call{"x", "list_notes", "{}"}}, {Call: &Call{"", "list_notes", "{}"}}}
	for i, reply := range cases {
		model := modelFunc(func(context.Context, []Message) (Reply, error) { return reply, nil })
		if _, _, err := Run(context.Background(), model, "test", Config{2, 2}); err == nil {
			t.Fatalf("第%d种错误回复被接受", i)
		}
	}
}
func TestDuplicateCallIDIsRejected(t *testing.T) {
	model := modelFunc(func(context.Context, []Message) (Reply, error) {
		return Reply{Call: &Call{"same", "list_notes", "{}"}}, nil
	})
	_, _, err := Run(context.Background(), model, "test", Config{3, 3})
	if err == nil || !strings.Contains(err.Error(), "重复") {
		t.Fatalf("未发现重复标识: %v", err)
	}
}
func TestCancelledToolDoesNotRead(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := executeTool(ctx, Call{"x", "read_note", `{"name":"functions.md"}`})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("错误: %v", err)
	}
}
