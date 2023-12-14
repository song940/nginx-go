package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/song940/nginx-go/nginx"
)

func main() {
	// filePath := "/Users/Lsong/Projects/confbook/nginx/nginx.conf"
	var nginxConfigPath = "/Users/Lsong/Projects/confbook/nginx/sites-available"

	files, _ := os.ReadDir(nginxConfigPath)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".conf") {
			continue
		}
		fullpath := filepath.Join(nginxConfigPath, file.Name())
		f, _ := os.Open(fullpath)
		items := nginx.ParseNginxConfig(f)
		log.Println(items[0].(*nginx.Block).GetServerNames())
	}
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// nginxConfig := nginx.ParseNginxConfig(file)
	// printNginxConfig(nginxConfig)
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
