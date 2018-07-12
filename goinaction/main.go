package main

import (
	"fmt"
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	// search.Run("abc")

	var a int = 10 //指定变量类型
	var b = 10
	c := 10
	fmt.Println(a, b, c)

}
