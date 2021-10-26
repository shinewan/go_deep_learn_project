package main

/**
	接口变量里面有什么?
	接口变量
		实现者的类型 +  实现者的值（或是实现者的指针）  -----> 实现者

	1.接口变量自带指针
	2.接口变量同样采用值传递，几乎不需要使用接口的指针
	3.指针接受者实现只能以指针方式使用；值接受者都可
 */


import (
	"fmt"
	"go_deep/05interface/retriever/mock"//go_deep 这是叫go.mod里面的方法
	"go_deep/05interface/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

//在添加一个新的接口
type Poster interface {
	Post(url string,form map[string]string) string
}
func post(poster Poster){
	poster.Post("http://www.imooc.com",map[string]string{
		"name": "ccmouse",
		"course":"golang",
	})
}


//组合接口
type RetrieverPoster interface{
	Retriever
	Poster
}


const url = "http://www.imooc.com"
func session(s RetrieverPoster) string{
	s.Post(url,map[string]string{
		"contents":"another faked imooc com",
	})
	return s.Get(url)
}


func download(r Retriever) string{
	return r.Get("http://www.imooc.com") //r.Get调用接口的方法，就会调用绑定接口的类型的同一个方法
}

func inspect(r Retriever){
	fmt.Printf("%T --- %v\n",r,r)
	switch v := r.(type) {   //r.()类型断言
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever//r 可被赋值为不同类型或指针的实现接口方法的类型
	r = mock.Retriever{"this is a fake com"}
	//fmt.Printf("%T --- %v\n",r,r)
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla5.0",
		Timeout: time.Minute,

	}
	//fmt.Printf("%T --- %v\n",r,r)
	inspect(r)

	//type assertion类型断言
	mockRetriever, ok := r.(mock.Retriever)
	if ok {
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}

	//dfmt.Println(download(r))
}
