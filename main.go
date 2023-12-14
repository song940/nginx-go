package main

import (
	"fmt"
	"log"
	"os"

	"github.com/song940/nginx-go/nginx"
)

func main() {
	filePath := "/Users/Lsong/Projects/confbook/nginx/nginx.conf"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	nginxConfig := nginx.ParseNginxConfig(file)
	printNginxConfig(nginxConfig)
}

func printNginxConfig(items []nginx.Items) {
	for _, item := range items {
		switch v := item.(type) {
		case *nginx.Block:
			printBlock(v)
		case *nginx.Directive:
			printDirective(v)
		}
	}
}

func printBlock(block *nginx.Block) {
	log.Println(block.Name)
	printNginxConfig(block.Items)
}

func printDirective(directive *nginx.Directive) {
	log.Println(directive.Name, "===>", directive.Value)
}
