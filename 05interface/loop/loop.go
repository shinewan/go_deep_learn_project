package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string{
	result := ""
	for ; n>0 ; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	//file.Read() 实现了read方法，因此，file实现了io.Reader
	//file.Write()
	/**
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	 */
	printFileContents(file)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



func main() {
	//底层关于读写相关的
	//fmt.Fprintf()
	//fmt.Fscanf()
	fmt.Println(
		convertToBin(5),//101
		convertToBin(13),//1101
		convertToBin(72387885),
		convertToBin(0),

		)
	printFile("05interface/loop/abc.txt")

	s := `
		abcfs"dd"
		kkkk
		12345
	
		hr
    `
	printFileContents(strings.NewReader(s))

}
