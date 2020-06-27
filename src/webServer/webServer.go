package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "io/ioutil"
)

func main1() {
    mux := http.NewServeMux()
    var htmlPath="/home/tiemuhua/mongoWithGo/src/webServer/register.html"
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
            s, e := ioutil.ReadFile(htmlPath) //这个就是读取你的html
            if e != nil {
               panic(e)
            }
            fmt.Fprintf(w, string(s)) //这个就是response
    })
	server := &http.Server{
		Addr     : ":8080", //监听服务器端口
		Handler  : mux,
	}
	if e := server.ListenAndServe(); e != http.ErrServerClosed {
		fmt.Printf("error was happened: %v.", e)
	}
	fmt.Println("server was shut down.")
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    var htmlPath="/home/tiemuhua/mongoWithGo/src/webServer/register.html"
    requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles(htmlPath)
        t.Execute(w, "")
    })
    
    server := &http.Server{
        Addr:           ":8080",
        Handler:        requestHandler,
        ReadTimeout:    10 ,
        WriteTimeout:   10 ,
        MaxHeaderBytes: 1 << 20,
    }
    
    log.Fatal(server.ListenAndServe())
    // 解析url传递的参数
    fmt.Fprintf(w,"hello!")
    s, e := ioutil.ReadFile(htmlPath)
    if e!=nil{panic(e)}
    fmt.Fprintf(w, string(s))
    r.ParseForm()
    for k, v := range r.Form {
        fmt.Println("key:", k)
        // join() 方法用于把数组中的所有元素放入一个字符串。元素是通过指定的分隔符进行分隔的
        fmt.Println("val:", strings.Join(v, ""))
    }
    // 输出到客户端
    name :=r.Form["username"]
    pass :=r.Form["password"]
    for i,v :=range name{
        fmt.Println(i)
        fmt.Fprintf(w,"用户名:%v\n",v)
    }
    for k,n :=range pass{
        fmt.Println(k)
        fmt.Fprintf(w,"密码:%v\n",n)
    }
}
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        // func (t *Template) Execute(wr io.Writer, data interface{}) error {
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}
func main() {
    http.HandleFunc("/", sayHelloName)
    /*http.HandleFunc("/login", login)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndserve:", err)
    }*/
    var htmlPath="/home/tiemuhua/mongoWithGo/src/webServer/register.html"
    requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        t, _ := template.ParseFiles(htmlPath)
        t.Execute(w, "")
    })
    
    server := &http.Server{
        Addr:           ":8080",
        Handler:        requestHandler,
        ReadTimeout:    10 ,
        WriteTimeout:   10 ,
        MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()
}
/*
func sayhelloName(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(request.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//var name string
	fmt.Fprintf(writer, "Hello Wrold!") //这个写入到w的是输出到客户端的
}
func main1() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func In_slice(val string, slice []string) bool {
    for _, v := range slice {
        if v == val {
            return true
        }
    }
    return false
}
func Slice_diff(slice1, slice2 []string) (diffslice []string) {
    for _, v := range slice1 {
        if !In_slice(v, slice2) {
            diffslice = append(diffslice, v)
        }
    }
    return
}
func sayhelloname(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //解析参数,默认是不会解析的。
    fmt.Println(r.Form)
    fmt.Println("path:", r.URL.Path)
    fmt.Println("scheme:", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("value:", strings.Join(v, ","))
    }
    fmt.Fprintf(w, "hello, welcome you!") //这个字符串写入到w中，用于返回给客户端。
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username: ", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func register(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("register.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        //r.Form对不通类型的表单元素的留空的处理：
        //对于空文本框、空文本区域及文件上传，元素的值为空值
        //如果是未选中的复选框和单选按钮，不会在r.Form中产生相应的条目，通过r.Form获取不存在的元素会报错。
        //r.Form.get()可以返回不存在元素的值为空值，不会报错，但只能获取单个的值，如果是map的值，必须通过r.Form来获取。
        if len(r.Form["username"][0]) == 0 {
            //用户名为空的处理
        }
        //期望用户名为中文。
        if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("username")); !m {
            //username不是中文的或不完全是中文的
        }
        //期望用户名是英文。
        if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
            //username不是英文的或不完全是英文的
        }
        //strconv包提供了字符串与简单数据类型之间的类型转换功能。可以将简单类型转换为字符串，也可以将字符串转换为其它简单类型
        //这个包里提供的函数有一个是，字符串转int：Atoi()
        getint, err := strconv.Atoi(r.Form.Get("age"))
        if err != nil {
            //转化出错了，age可能不是int类型。
        }
        if getint <= 0 && getint > 100 {
            //年龄不在正常范围内
        }
        //使用正则表达式来判断,regexp.MatchString函数用于将字符串与正则表达式进行匹配，如果能匹配成功，返回true
        if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
            //age不是int类型
        }
        //验证电子邮件是否满足格式的要求
        if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
            //不满足要求
        }
        //验证手机号是否满足格式要求
        if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
            //不满足要求
        }
        //验证下拉菜单
        option_slice := []string{"apple", "pear", "banane"}
        for _, v := range option_slice {
            if v == r.Form.Get("fruit") {
                //下拉选项正确的情况
            }
            //下拉选项不正确的情况
        }
        //验证单选按钮
        radio_slice := []int{1, 2}
        for _, v := range radio_slice {
            if strconv.Itoa(v) == r.Form.Get("gender") {
                //单选按钮正确的情况
            }
            //单选按钮错误的情况
        }
        //验证复选框
        checkbox_slice := []string{"football", "basketball", "tennis"}
        a := Slice_diff(r.Form["interest"], checkbox_slice)
        if a == nil {
            //复选框按钮正确的情况
        } else {
            //复选框按钮错误的情况
        }
        //验证15位身份证，15位的是全部数据
        if m, _ := regexp.MatchString(`^(\d{15})`, r.Form.Get("usercard")); !m {
            //不满足的情况
        }
        //验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或者字符X
        if m, _ := regexp.MatchString(`^(\d{17}([0-9]|X))$`, r.Form.Get("usercard")); !m {
            //不满足的情况
        }
    }
}

func giveHTML(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w,r,"login.html")
}

func main1() {
    http.HandleFunc("/",giveHTML)
    //http.HandleFunc("/", sayhelloname)       //设置访问的路由
    //http.HandleFunc("/login", login)         //设置访问的路由
    //http.HandleFunc("/register", register)   //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
*/