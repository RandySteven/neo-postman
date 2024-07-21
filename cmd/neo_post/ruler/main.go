package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const templateCommand = "create models %s -p -r -u -h -q\n"

/**exp
-p = payload
-r = repositories
-u = usecases
-h = handlers
-q = queries
*/

func main() {
	cliV2()
}

func cliV2() {
	commandArgs := os.Args[1:]
	if len(commandArgs) == 0 {
		fmt.Println("No command provided")
		return
	}

	switch commandArgs[0] {
	case "model":
		if len(commandArgs) < 2 {
			fmt.Println("Model name is required")
			return
		}
		createModel(commandArgs[1])
	case "-r", "--repository":
		if len(commandArgs) < 2 {
			fmt.Println("Repository name is required")
			return
		}
		createRepository(commandArgs[1])
	default:
		fmt.Println("Unknown command")
	}
}

func createModel(model string) {
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

func createRepository(repository string) {
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
