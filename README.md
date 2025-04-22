# Task CLI

Task CLI is a simple command-line tool for managing tasks. It allows you to add, list, mark as done, and delete tasks, with tasks stored in a JSON file in your home directory.

## Installation:

1. Clone the repository:

```
git clone https://github.com/mihabgit/go-todo-cli.git
```


2. Navigate to the project directory:

```
cd task-cli
```


3. Build the project:

```
go build
```

Note: You need to have Go installed and set up properly on your system. If you don't have Go installed, you can download it from golang.org.

## Usage

1. After building the project, you can use the following commands to manage your tasks:

add: Add a new task

```
./todo-cli add "Buy milk"
```
This will add a new task with the description "Buy milk" and display the task ID.


2. list: List all tasks

```
./todo-cli list
```
This will show the list of todos

3. complete: Mark a task as done

```
./todo-cli complete 1
```
This will mark the task with ID 1 as done.


Tasks are stored in a JSON file located at  ```~/.todo-cli/todo.json.```

## Project Structure

- ```cmd/```: Contains the Cobra command definitions for add, list, complete.

- ```todo/```: Contains the task management logic, including functions for adding, listing, marking as done.

- ```main.go```: The entry point of the application.

## License
This project is licensed under the MIT License.
