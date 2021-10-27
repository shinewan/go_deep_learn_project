package mock

import "fmt"

type Retriever struct{
	Contents string
}

// r Retriever 这种方式是值传递
//r *Retriever 这种方式可以同时改变
func (r *Retriever) Post(url string,form map[string] string) string{
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string{
	return r.Contents
}

//相当于toString，打印类的结构信息
func (r *Retriever) String() string{
	return fmt.Sprintf("Retriever:{Contents=%s}",r.Contents)
}


