package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	i := -1
	file := ""
	model := ""
	repository := ""
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
			modelStruct := strings.ToUpper(string(model[0])) + model[1:]
			createNewStruct := fmt.Sprintf(`package models \n\ntype %s struct {}`, modelStruct)
			cmd := exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" > ./entities/models/%s && clear", createNewStruct, file))

			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(out)
		case 2:
		//TODO
		case 3:
			repositoryIntrContent, err := ioutil.ReadFile("./files/txt/template/repository_interfaces.txt")
			if err != nil {
				fmt.Println(err)
				return
			}

			repositoryContent, err := ioutil.ReadFile("./files/txt/template/repository.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			repositoryStruct := string(repositoryContent)
			repositoryIntrface := string(repositoryIntrContent)

			fmt.Print("Insert repository name : ")
			fmt.Scanf("%s", &repository)

			modelName := strings.ToUpper(string(repository[0])) + repository[1:]
			structName := strings.ToLower(string(repository[0])) + repository[1:] + "Repository"
			interfaceName := strings.ToUpper(string(repository[0])) + repository[1:] + "Repository"

			file = fmt.Sprintf("%s_repository.go", strings.ToLower(repository))

			createIntrfaceCmd := fmt.Sprintf(repositoryIntrface, interfaceName, modelName)
			createStrctCmd := fmt.Sprintf(repositoryStruct, structName, interfaceName, structName)

			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				cmd := exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" > ./interfaces/repositories/%s && clear", createIntrfaceCmd, file))
				_, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println(err)
					return
				}
			}()

			go func() {
				defer wg.Done()
				cmd := exec.Command("sh", "-c", fmt.Sprintf("echo \"%s\" > ./repositories/%s && clear", createStrctCmd, file))
				_, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println(err)
					return
				}
			}()

			wg.Wait()
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
