package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	cache := base.NewLruCache(3)

	cache.Put("hello", "world")
	res := cache.Get("hello")

	if res != nil {
		fmt.Println(*res)
	}
}
