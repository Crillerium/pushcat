package main

import (
	"github.com/electricbubble/go-toast"
	"flag"
)

func main() {
	var title string
	var content string
	flag.StringVar(&title,"t","未定义の标题","标题")
	flag.StringVar(&content,"c","还没有设置内容呢~","内容")
	flag.Parse()
	_ = toast.Push(content, toast.WithTitle(title))
}
