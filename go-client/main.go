package main

import (
	"fmt"

	genezioSdk "go-client/sdk"
)

func main() {
	helloService := genezioSdk.NewHello()
	response, err := helloService.SayHello()
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	response2, err := helloService.SayHelloTo("Genezio")
	if err != nil {
		panic(err)
	}
	fmt.Println(response2)

	taskService := genezioSdk.NewTaskService()
	tasks, err := taskService.GetTasks()
	for _, task := range tasks {
		fmt.Printf("Task %d: %s\n", task.Id, task.Description)
	}
}
