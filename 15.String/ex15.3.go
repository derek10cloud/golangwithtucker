package main

import "fmt"

func main() {
	var char rune = '๐'

	fmt.Printf("%T\n", char) //rune ํ์์ int32์ ๊ฐ์์ int32์ถ๋ ฅ
	fmt.Println(char)        //char๊ฐ ์ถ๋ ฅ, int32๋ผ ์ซ์๋ก ์ถ๋ ฅ
	fmt.Printf("%c\n", char) //%cํฌ๋งท ์ฌ์ฉํ๋ฉด ๋ฌธ์ ์ถ๋ ฅ
}
