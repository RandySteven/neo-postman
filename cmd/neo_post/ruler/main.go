package main

import (
	"fmt"
	"strings"
)

func main() {
	i := -1
	file := ""
	model := ""
	for i != 0 {
		menu()
		fmt.Print(">>")
		fmt.Scanf("%d", &i)
		switch i {
		case 0:
			fmt.Println("Byee")
		case 1:
			fmt.Print("Insert model name : ")
			fmt.Scanf("%s", &model)
			file = fmt.Sprintf("%s.go", strings.ToLower(model))
			cli := fmt.Sprintf("nano ./entities/models/%s", file)
			fmt.Println(cli)
		}
	}
}

func menu() {
	fmt.Println("1. Create model")
	fmt.Println("2. Create mock")
	fmt.Println("3. Create repository")
	fmt.Println("4. Create payload")
	fmt.Println("5. Create usecase")
}
