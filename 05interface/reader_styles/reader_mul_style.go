package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

/**
GoLang提供了很多读文件的方式，一般来说常用的有三种。使用Read加上buffer，使用bufio库和ioutil库

运行命令go run read.go filename
 */

func read1(path string)string{
	fi,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte,1024,1024)
	buf := make([]byte,1024)
	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func read2(path string)string{
	fi,err := os.Open(path)
	if err != nil{panic(err)}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte,1024,1024)

	buf := make([]byte,1024)
	for{
		n,err := r.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {
			break
		}
		chunks=append(chunks,buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func read3(path string)string{
	fi,err := os.Open(path)
	if err != nil{panic(err)}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func main(){

	flag.Parse()
	for idx,args := range os.Args{
		fmt.Println("参数" + strconv.Itoa(idx) + ":",args)
	}
	/**
	fi,err := os.Open("05interface/loop/abc.txt")
	if err != nil{panic(err)}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	fmt.Println(string(fd))
	*/
	file := flag.Arg(1)
	f,err := ioutil.ReadFile(file)
	if err != nil{
		fmt.Printf("%s\n",err)
		panic(err)
	}
	fmt.Println(string(f))
	start := time.Now()
	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n",t1.Sub(start))
	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n",t2.Sub(t1))
	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n",t3.Sub(t2))

}
