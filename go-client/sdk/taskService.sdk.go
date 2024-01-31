/**
* This is an auto generated code. This code should not be modified since the file can be overwritten
* if new genezio commands are executed.
*/

package genezioSdk

import "encoding/json"

type Task struct {
    Id int
    Description string
    Completed bool
}


type TaskService struct {
	remote *Remote
}

func NewTaskService() *TaskService {
	return &TaskService{remote: &Remote{URL: "https://seea5ouurawwh4v4oetk6bygcm0wctua.lambda-url.eu-central-1.on.aws/" }}
}
func (f *TaskService) GetTasks() ([]Task, error) {
    result, err := f.remote.Call("TaskService.GetTasks")
    var castResult []Task
    if err != nil {
        return castResult, err
    }
    jsonMap, err := json.Marshal(result)
    if err != nil {
        return castResult, err
    }
    err = json.Unmarshal(jsonMap, &castResult)
    if err != nil {
        return castResult, err
    }
	return castResult, err
}

