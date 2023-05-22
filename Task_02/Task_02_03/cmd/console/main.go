package main

import (
	"fmt"
	"time"

	"example.com/tykoon/pkg/shortcache"
)

func main() {
	var cache = shortcache.New(10)

	cache.Set("integer", 1024)
	fmt.Println(`Set("integer", 1024)`)
	WaitASecond(1)

	cache.Set("float", 10.24)
	fmt.Println(`Set("float", 10.24)`)
	WaitASecond(1)

	fmt.Println("Current count:", cache.Count())
	WaitASecond(1)

	fmt.Println("float: ", cache.Get("float"))
	WaitASecond(1)

	fmt.Println("Remove float: ", cache.Remove("float"))
	WaitASecond(3)

	fmt.Println("Current count:", cache.Count())
	WaitASecond(4)

	fmt.Println("Current count:", cache.Count())
}

func WaitASecond(s int) {
	time.Sleep(time.Second * time.Duration(s))
}
