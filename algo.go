package main

import (
	"fmt"
	"github.com/jokerl/go_study/algo"
	"rsc.io/quote"
	"time"
)

func main() {
	fmt.Println(quote.Hello())
	//abcabcbb -> abc 3
	//bbbbb -> b 1
	//pwwkew -> wke 3
	// abc -> abc 3
	// "" -> ""  0
	time.Now()
	algo.SearchLengthSonString("abcabcbb")
	algo.SearchLengthSonString("bbbbb")
	algo.SearchLengthSonString("pwwkew")
	algo.SearchLengthSonString("abc")
	algo.SearchLengthSonString("h;l;l朝陈")
	algo.SearchLengthSonString("")
	algo.SearchLengthSonString("pwkew1234")

}
