1.程序执行结果为( )

    func main() {
        defer_call()
    }
    
    func defer_call() {
    defer func() {
        fmt.Print("打印前 ")
    }()
     
    defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
        if err := recover();err != nil {
            fmt.Print(err) //err 就是panic传入的参数
        }
        fmt.Print("打印中 ")
    }()
     
    defer func() {
     
        fmt.Print("打印后 ")
    }()
     
    panic("触发异常 ")
    }
A. 触发异常  打印后  打印中 打印前
B.打印后 触发异常 打印中   打印前
C.打印后 触发异常 打印前   打印中
D.打印后 打印中  打印前  触发异常

答案:B

2.以下代码有什么问题，说明原因
type student struct {
    Name string
    Age  int
}
func pase_student() map[string]*student {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    for _, stu := range stus {
        m[stu.Name] = &stu
    }
    return m
}
func main() {
    students := pase_student()
    for k, v := range students {
        fmt.Printf("key=%s,value=%v \n", k, v)
    }
}

答案:每次遍历仅进行struct值拷贝，故m[stu.Name]=&stu实际上一直指向同一个地址，最终该地址的值为遍历的最后一个struct的值拷贝。


3. 下面的代码会输出什么，并说明原因

func init() {
    fmt.Println("Current Go Version:", runtime.Version())
}
func main() {
    runtime.GOMAXPROCS(1)

    count := 10
    wg := sync.WaitGroup{}
    wg.Add(count * 2)
    for i := 0; i < count; i++ {
        go func() {
            fmt.Printf("[%d]", i)
            wg.Done()
        }()
    }
    for i := 0; i < count; i++ {
        go func(i int) {
            fmt.Printf("-%d-", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}

答案：
Current Go Version: go1.10.1
-9-[10][10][10][10][10][10][10][10][10][10]-0--1--2--3--4--5--6--7--8-
Process finished with exit code 0
两个for循环内部go func 调用参数i的方式是不同的，导致结果完全不同。这也是新手容易遇到的坑。
第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10（10次遍历很快完成）。
第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

4. 下面代码会输出什么？
type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowA()
}

答案：
showA
showB

5. 下面代码会触发异常吗？请详细说明
func main() {
    runtime.GOMAXPROCS(1)
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {
        case value := <-int_chan:
            fmt.Println(value)
        case value := <-string_chan:
            panic(value)
    }
}


答案：有可能会发生异常，如果没有selct这段代码，就会出现线程阻塞，当有selct这个语句后，系统会随机抽取一个case进行判断，只有有其中一条语句正常return，此程序将立即执行。

6. 下面代码输出什么？

func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}

func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
    b = 1
}

答案：
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
不管代码顺序如何，defer calc func中参数b必须先计算，故会在运行到第三行时，执行calc("10",a,b)输出：10 1 2 3得到值3，将cal("1",1,3)存放到延后执执行函数队列中。
执行到第五行时，现行计算calc("20", a, b)即calc("20", 0, 2)输出：20 0 2 2得到值2,将cal("2",0,2)存放到延后执行函数队列中。
执行到末尾行，按队列先进后出原则依次执行：cal("2",0,2)、cal("1",1,3)，依次输出：2 0 2 2、1 1 3 4 。

7. 请写出以下输入内容

func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}

答案：[0 0 0 0 0 1 2 3]

8. 以下代码能编译过去吗？为什么？
type People interface {
    Speak(string) string
}
type Stduent struct{}
func (stu *Stduent) Speak(think string) (talk string) {
    if think == "bitch" {
        talk = "You are a good boy"
    } else {
        talk = "hi"
    }
    return
}
func main() {
    var peo People = Stduent{}
    think := "bitch"
    fmt.Println(peo.Speak(think))
}
答案：编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People 类型。


9.使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个 goroutine 打印字母， 最终效果如下 
12AB34CD56EF78GH910IJ。
答案：
chan_n := make(chan bool)
chan_c := make(chan bool, 1)
done := make(chan struct{})

go func() {
  for i := 1; i < 11; i += 2 {
    <-chan_c  
    fmt.Print(i)
    fmt.Print(i + 1)
    chan_n <- true 
  }
}()

go func() {
  char_seq := []string{"A","B","C","D","E","F","G","H","I","J","K"}
  for i := 0; i < 10; i += 2 {
    <-chan_n 
    fmt.Print(char_seq[i])
    fmt.Print(char_seq[i+1])
    chan_c <- true  
  }
  done <- struct{}{} 
}()

chan_c <- true 
<-done