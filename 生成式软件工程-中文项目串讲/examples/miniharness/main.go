// 这是可重复的教学运行器：只使用内存笔记和固定脚本，不访问真实模型或文件。
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Message是本样例的消息格式，不是任何供应商的协议结构。
// 所有字段都是值类型；给模型复制切片后，模型不能改写已保存的字段。
type Message struct {
	Role, Text, CallID, Tool, Arguments string
	IsError                             bool
}

// Call把一次请求的标识、工具名和JSON参数放在一起。
type Call struct{ ID, Name, Arguments string }

// Reply在本样例中只能是最终文字或一个工具请求，不能同时存在。
type Reply struct {
	Text string
	Call *Call
}

// Model描述“根据当前历史取得下一次回复”的能力。
type Model interface {
	Next(context.Context, []Message) (Reply, error)
}
type Config struct{ MaxSteps, MaxTools int }

var ErrStepLimit = errors.New("达到模型请求次数上限")
var ErrToolLimit = errors.New("达到工具调用次数上限")

// Run执行完整教学任务；失败也把已经记录的历史交回调用者。
// 它只保存内存历史，不承诺持久化、并行、恢复、审批或完整事件日志。
func Run(ctx context.Context, model Model, question string, cfg Config) (string, []Message, error) {
	history := []Message{{Role: "user", Text: question}}
	if ctx == nil || model == nil || cfg.MaxSteps <= 0 || cfg.MaxTools < 0 {
		return "", history, errors.New("运行参数无效")
	}
	toolCount := 0
	seenCalls := make(map[string]bool)
	for step := 0; step < cfg.MaxSteps; step++ {
		if err := ctx.Err(); err != nil {
			return "", history, err
		}
		// 副本防止模型实现直接改写已记录的历史。
		snapshot := append([]Message(nil), history...)
		reply, err := model.Next(ctx, snapshot)
		if err != nil {
			return "", history, fmt.Errorf("模型调用失败: %w", err)
		}
		// 一个不配合取消的模型可能迟到返回，仍不能把它当作成功。
		if err := ctx.Err(); err != nil {
			return "", history, err
		}
		if reply.Call == nil {
			if strings.TrimSpace(reply.Text) == "" {
				return "", history, errors.New("模型返回了空回复")
			}
			history = append(history, Message{Role: "assistant", Text: reply.Text})
			return reply.Text, history, nil
		}
		if reply.Text != "" {
			return "", history, errors.New("教学协议不接受文字与工具请求并存")
		}
		call := *reply.Call
		if call.ID == "" || seenCalls[call.ID] {
			return "", history, errors.New("工具调用标识为空或重复")
		}
		seenCalls[call.ID] = true
		history = append(history, Message{Role: "assistant", CallID: call.ID, Tool: call.Name, Arguments: call.Arguments})
		if toolCount >= cfg.MaxTools {
			return "", history, ErrToolLimit
		}
		if err := ctx.Err(); err != nil {
			return "", history, err
		}
		toolCount++
		result, toolErr := executeTool(ctx, call)
		// 取消后停止任务，不让迟到的结果被解释成任务正常完成。
		if err := ctx.Err(); err != nil {
			return "", history, err
		}
		message := Message{Role: "tool", CallID: call.ID, Tool: call.Name, Text: result}
		if toolErr != nil {
			message.Text = toolErr.Error()
			message.IsError = true
		}
		history = append(history, message)
	}
	return "", history, ErrStepLimit
}

// decodeObject同时拒绝非对象、未知字段以及多个连续JSON值。
func decodeObject(raw string, target any) error {
	text := strings.TrimSpace(raw)
	if !strings.HasPrefix(text, "{") {
		return errors.New("参数必须是JSON对象")
	}
	decoder := json.NewDecoder(strings.NewReader(text))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("参数不合法: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return errors.New("参数后存在多余内容")
	}
	return nil
}

// executeTool只允许这里明确登记的两个工具，且不读取操作系统文件。
func executeTool(ctx context.Context, call Call) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	switch call.Name {
	case "list_notes":
		var args struct{}
		if err := decodeObject(call.Arguments, &args); err != nil {
			return "", err
		}
		return `["functions.md","variables.md"]`, nil
	case "read_note":
		var args struct {
			Name string `json:"name"`
		}
		if err := decodeObject(call.Arguments, &args); err != nil {
			return "", err
		}
		notes := map[string]string{
			"variables.md": "变量会声明。",
			"functions.md": "多返回值还不懂。",
		}
		text, allowed := notes[args.Name]
		if !allowed {
			return "", errors.New("拒绝访问：仅允许固定教学笔记")
		}
		return text, nil
	default:
		return "", fmt.Errorf("未知工具: %s", call.Name)
	}
}

// ScriptedModel按已取得的工具结果推进，用于观察接线，不代表真实智能。
type ScriptedModel struct{}

func (ScriptedModel) Next(ctx context.Context, history []Message) (Reply, error) {
	if err := ctx.Err(); err != nil {
		return Reply{}, err
	}
	var listResult, readResult *Message
	for i := range history {
		if history[i].Role != "tool" {
			continue
		}
		switch history[i].Tool {
		case "list_notes":
			listResult = &history[i]
		case "read_note":
			readResult = &history[i]
		}
	}
	if readResult != nil {
		if readResult.IsError {
			return Reply{Text: "读取失败，尚不能给出笔记结论：" + readResult.Text}, nil
		}
		return Reply{Text: "依据 functions.md：" + readResult.Text + "下一步先练习接收两个返回值，并检查error。"}, nil
	}
	if listResult != nil {
		if listResult.IsError {
			return Reply{Text: "列目录失败：" + listResult.Text}, nil
		}
		var names []string
		if err := json.Unmarshal([]byte(listResult.Text), &names); err != nil {
			return Reply{}, err
		}
		found := false
		for _, name := range names {
			if name == "functions.md" {
				found = true
			}
		}
		if !found {
			return Reply{Text: "没有找到函数笔记，需要补充资料。"}, nil
		}
		return Reply{Call: &Call{"c2", "read_note", `{"name":"functions.md"}`}}, nil
	}
	return Reply{Call: &Call{"c1", "list_notes", "{}"}}, nil
}

func main() {
	steps := flag.Int("steps", 3, "最多模型请求次数")
	tools := flag.Int("tools", 2, "最多工具调用次数")
	cancelled := flag.Bool("cancel", false, "在开始前取消")
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if *cancelled {
		cancel()
	}
	answer, history, err := Run(ctx, ScriptedModel{}, "我下一步练什么？", Config{*steps, *tools})
	for _, message := range history {
		if message.CallID != "" {
			fmt.Printf("[%s] %s %s %s %s error=%v\n", message.Role, message.CallID, message.Tool, message.Arguments, message.Text, message.IsError)
		} else {
			fmt.Printf("[%s] %s\n", message.Role, message.Text)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "停止原因：", err)
		os.Exit(1)
	}
	fmt.Println("最终回答：", answer)
}
