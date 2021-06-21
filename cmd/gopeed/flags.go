package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// url下载地址 connections并发数 dir存放路径
type args struct {
	url         string
	connections *int
	dir         *string
}

func parse() *args {
	// 返回路径dir
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	var args args
	args.connections = flag.Int("C", 16, "Concurrent connections.")
	args.dir = flag.String("D", dir, "Save directory.")
	flag.Parse()
	// t为接受的命令行参数
	t := flag.Args()

	// t的命令行参数个数大于0
	if len(t) > 0 {
		// 把第一个参数存储在url
		args.url = t[0]
	} else {
		gPrintln("missing url parameter")
		gPrintln("try 'gopeed -h' for more information")
		os.Exit(1)
	}
	return &args
}

func gPrint(msg string) {
	fmt.Print("gopeed: " + msg)
}

func gPrintln(msg string) {
	gPrint(msg + "\n")
}
