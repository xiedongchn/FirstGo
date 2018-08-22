package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("你输入的是：", input.Text())
}