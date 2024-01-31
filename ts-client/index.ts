import { Hello, Task, TaskService } from '@genezio-sdk/go-server_eu-central-1';

(async () => {
    const response = await Hello.SayHello();
    console.log(response);
    const response2 = await Hello.SayHelloTo('Genezio');
    console.log(response2);
    const tasks: Task[] = await TaskService.GetTasks();
    for (const task of tasks) {
        console.log('Task ' + task.Id + ': ' + task.Description);
    }
})();
