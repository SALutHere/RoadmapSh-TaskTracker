```
Usage: task-cli [ACTION] [OPTIONS]...
Allows you to manage your tasks.

Actions with their options:
	help						gives you a syntax-hint

	add DESCRIPTION				adds a new task with specified description
	update ID DESCRIPTION		updates a description of task by the specified id
	delete ID					deletes a task by the specified id
	
	mark-todo ID				sets status "todo" to the task by specified id
	mark-in-progress ID			sets status "in-progress" to the task by specified id
	mark-done ID				sets status "done" to the task by specified id
	
	list [STATUS]				lists your tasks. If you specified a status, lists
								lists your tasks only with specified status
									allowed statuses: "todo", "in-progress", "done"
									
	clear						clears your tasks list
```
