package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (cli *cli) createModel() {
	model := cli.command[1]
	file := fmt.Sprintf("%s.go", strings.ToLower(model))
	modelStruct := strings.ToUpper(string(model[0])) + model[1:]
	createNewStruct := fmt.Sprintf(`package models\n\ntype %s struct {}`, modelStruct)
	cmd := exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" > ./entities/models/%s", createNewStruct, file))

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(out))
		return
	}
	fmt.Println("Model created successfully")
}
