# Go、C#与源码：按同一个问题对照学习

## 学习目标与用法

每题练习册已指定当前要读的入口和C#对照对象。主项目继续用Go；C#先用小片段、测试客户端或单个模块对照，逐步积累第二种表达方式。若两种语言都不熟，先把Go当前步骤跑通，再对照同一个行为，不要求同时重写整套系统。

对照时固定四件事：输入、输出、失败、资源/生命周期。相似名称不等于相同机制；框架没有一一对应时，对照接口合同和业务行为。

## 01 · 先看一组函数参数与返回值

下面两个程序都把设备名称交给函数，接住结果并打印。它们用来逐行对照；2026-10-24已在临时目录用Go 1.27.0与.NET SDK 10.0.402分别编译运行，均输出“你好，sw-core”。HTTP方法是局部示意片段，未作为完整Web程序运行。

~~~go
package main

import "fmt"

// greet只生成文字；调用者负责输出，方便独立验证返回值。
func greet(name string) string {
    return "你好，" + name
}

func main() {
    message := greet("sw-core")
    fmt.Println(message)
}
~~~

~~~csharp
using System;

class Program
{
    // Greet只生成文字；调用者决定怎样输出。
    static string Greet(string name)
    {
        return "你好，" + name;
    }

    static void Main()
    {
        var message = Greet("sw-core");
        Console.WriteLine(message);
    }
}
~~~

Go参数写“名称 类型”，C#写“类型 名称”；Go返回类型在参数括号后，C#写在方法名前。Go的message :=和C#的var message都可以从右侧推断类型，但语法规则与作用域细节仍按各自语言。两种语言都要区分定义与调用，return把结果交给调用方，打印则是副作用。

本题变式：名称改为参数、加入空名称分支、让函数只生成文字而不打印，最后换一个设备名称预测结果。小改动后再运行，不靠记住原样例输出。

## 02 · 把原HTTP方法换一种语言表达

原Go片段：

~~~go
func (h *httpHandler) answer(w http.ResponseWriter, r *http.Request) {
    requestID := normalizeRequestID(r.Header.Get("X-Request-ID"))
    w.Header().Set("X-Request-ID", requestID)
}
~~~

对应ASP.NET Core概念的示意片段，NormalizeRequestID仍是需要你定义合同的函数：

~~~csharp
void Answer(HttpContext context)
{
    string rawID = context.Request.Headers["X-Request-ID"].ToString();
    string requestID = NormalizeRequestID(rawID);
    context.Response.Headers["X-Request-ID"] = requestID;
}
~~~

这两个片段都没有写正文。Go通过r字段读头、w.Header()方法取得响应头；C#通过HttpContext访问Request/Response属性。变量h、w、r和context都不是必须的关键字；类型、调用位置和框架合同才决定用法。

多值请求头压成单个字符串的行为可能不同，所以这个例子只比较单值头的流程。真实实现应明确重复头如何处理，不能因示意代码看起来相似就断言所有请求都等价。

Go方法接收者不是把C#的this机械换成一个*号。Go显式声明接收者；C#实例方法隐含this。Go此处没有返回值，C#示意为void；若C#实际用异步写响应，则常见返回Task并await写入。Task用于表示异步操作完成，不是直接返回给浏览器的正文。

[Go net/http](https://pkg.go.dev/net/http)、[ASP.NET Core中间件](https://learn.microsoft.com/en-us/aspnet/core/fundamentals/middleware/?view=aspnetcore-10.0)。

## 03 · 容易类比错的地方

| 主题 | Go | C# | 学习时要验证 |
|---|---|---|---|
| 参数传递 | 参数按值传递，指针值也会复制 | 默认按值；class复制引用，struct复制值；ref/out等另有合同 | 改对象内容与重新赋参数是否影响调用方 |
| 缺失值 | 指针、额外bool、可空类型等 | nullable值类型、可空引用标注等 | 编译器警告、运行时null与业务缺失不是一件事 |
| 方法 | 显式接收者，值/指针方法集合 | 实例/静态方法和this | 调用者传了什么、是否修改共享对象 |
| 接口 | 方法集合满足即可隐式实现 | 通常在类型声明中明确实现接口 | 不从“接口名相同”推断合同一样 |
| 错误 | 常用结果和error多返回 | 常用异常，也有Try模式与结果类型 | 哪些失败是预期业务结果，在哪里映射 |
| 多个结果 | 原生多返回值 | 元组、out或专用结果类型 | 结构相似不代表错误模型相同 |
| 集合 | slice/map，可能共享底层数据 | List/Dictionary/数组也涉及共享引用 | 拷贝容器不一定深拷贝元素 |
| 资源 | defer配合Close/Unlock等 | using/await using/try-finally与Dispose | 释放发生在哪个作用域、异常路径是否覆盖 |
| 并发/异步 | goroutine、channel、锁、context | Task/async-await、Channel、锁、CancellationToken | Task不是线程，也不是goroutine的等号替代 |
| 可选更新 | 指针或显式存在性，ORM各有零值规则 | DTO、nullable、EF跟踪/选择更新策略 | 缺失、false、0、空字符串分别怎样写库 |
| 数据访问 | sql.DB通常管理连接池 | DbConnection表示连接，provider可池化 | 不把sql.DB直接类比成一条DbConnection |
| 依赖注入 | 可以直接构造并传依赖 | 常用DI容器，也可直接构造 | 生命周期与拥有者由工程设计决定 |
| 泛型 | 类型参数与约束 | 泛型与约束体系 | 不追求翻译所有特性，先读懂当前签名 |
| AI框架 | Eino组件与Graph | 选定SDK或自定义服务接口 | 没有强制的一一对应库，比较输入输出与授权 |

参考：[C#参数](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/method-parameters)、[C#异步](https://learn.microsoft.com/en-us/dotnet/csharp/asynchronous-programming/)、[Dispose](https://learn.microsoft.com/en-us/dotnet/standard/garbage-collection/implementing-dispose)、[EF Core并发](https://learn.microsoft.com/en-us/ef/core/saving/concurrency)。

## 04 · 源码从调用点读，不从仓库第一页读

一次源码阅读只解决一个具体问题，例如“Header.Set为什么没有返回值”“QueryRow的错误为什么在Scan出现”“中间件next究竟在哪调用”。

1. 先写自己的预测：传入什么、返回什么、会改哪里。
2. 看当前使用版本的签名和官方合同，记录Go/.NET/依赖版本。
3. 从自己的调用点跳到实现，先读直接函数及一个相关分支。
4. 用3–5个框画调用链，圈出错误、锁、网络/磁盘、共享状态。
5. 用一个最小实验验证预测。暂时不理解的运行时/汇编边界注明，先停止追踪。
6. 再看相关测试怎样证明合同；不要只背一张调用图。

Go标准库入口示例（在已有Go环境运行，只读源码）：

~~~text
go version
go env GOROOT
go doc -src fmt.Println
go doc -src net/http.HandlerFunc.ServeHTTP
go doc -src net/http.Header.Set
go doc -src database/sql.Row.Scan
~~~

第三方库从当前go.mod/go.sum确定版本，用编辑器跳转定义或模块源码定位。仓库main分支可能不同于本机依赖，不直接抄当前网页中的行号。首次记录符号和版本即可，以后要定位性能或并发时再向下追。

C#从实际SDK/包版本和编辑器跳转/Source Link进入实现；核心库看 [dotnet/runtime](https://github.com/dotnet/runtime)，Web管道看 [dotnet/aspnetcore](https://github.com/dotnet/aspnetcore)。反编译可辅助理解，但要区分反编译视图与对应版本原始源码。

## 05 · 源码阅读也分层重复

- 第一次：自己的函数与测试，能说出控制流。
- 第二次：标准库入口和直接调用，验证一个边界。
- 第三次：框架适配与中间件，比较合同如何保持。
- 第四次：为当前故障追锁、重试、池或资源释放，不为了“读完源码”无限下钻。

记录只需：问题、版本、入口、两三个关键分支、实验结果、仍未确认之处。对照C#时补“相同合同、机制差异、一个会类比错的反例”。

[回项目路线](README.md) · [总讲义](Go开发与运维知识串讲.md)
