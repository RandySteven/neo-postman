package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"
)

func (cli *cli) createRepository() {
	repository := cli.command[1]
	repositoryIntrContent, err := ioutil.ReadFile("./files/txt/template/repository_interfaces.txt")
	if err != nil {
		fmt.Println("Error reading interface template:", err)
		return
	}

	repositoryContent, err := ioutil.ReadFile("./files/txt/template/repository.txt")
	if err != nil {
		fmt.Println("Error reading repository template:", err)
		return
	}
	repositoryStruct := string(repositoryContent)
	repositoryIntrface := string(repositoryIntrContent)

	modelName := strings.ToUpper(string(repository[0])) + repository[1:]
	structName := strings.ToLower(string(repository[0])) + repository[1:] + "Repository"
	interfaceName := strings.ToUpper(string(repository[0])) + repository[1:] + "Repository"

	file := fmt.Sprintf("%s_repository.go", strings.ToLower(repository))

	createIntrfaceCmd := fmt.Sprintf(repositoryIntrface, interfaceName, modelName)
	createStrctCmd := fmt.Sprintf(repositoryStruct, structName, interfaceName, structName)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" > ./interfaces/repositories/%s", createIntrfaceCmd, file))
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error creating interface:", err)
			fmt.Println("Output:", string(out))
			return
		}
	}()

	go func() {
		defer wg.Done()
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" > ./repositories/%s", createStrctCmd, file))
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error creating repository:", err)
			fmt.Println("Output:", string(out))
			return
		}
	}()

	wg.Wait()
	fmt.Println("Repository created successfully")
}
