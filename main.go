package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const filename = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]

	taskList := loadTasks()

	switch command {
	case "add":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli add <description>")
			os.Exit(1)
		}
		description := os.Args[2]
		addTask(&taskList, description)
	case "update":
		if len(os.Args) != 4 {
			fmt.Println("Usage: task-cli update <id> <new description>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		newDescription := os.Args[3]
		updateTask(&taskList, id, newDescription)
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli delete <id>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(&taskList, id)
	case "mark-in-progress":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		markTaskStatus(&taskList, id, "in-progress")
	case "mark-done":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			os.Exit(1)
		}
		id, _ := strconv.Atoi(os.Args[2])
		markTaskStatus(&taskList, id, "done")
	case "list":
		if len(os.Args) == 2 {
			listAllTasks(taskList)
		} else if len(os.Args) == 3 {
			status := os.Args[2]
			listTasksByStatus(taskList, status)
		} else {
			fmt.Println("Usage: task-cli list [status]")
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}

	saveTasks(taskList)
}

func loadTasks() TaskList {
	var taskList TaskList

	data, err := os.ReadFile(filename)
	if err != nil {
		return taskList
	}

	err = json.Unmarshal(data, &taskList)
	if err != nil {
		fmt.Println("Error parsing tasks file:", err)
		os.Exit(1)
	}

	return taskList
}

func saveTasks(taskList TaskList) {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
		os.Exit(1)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}

func addTask(taskList *TaskList, description string) {
	newID := 1
	if len(taskList.Tasks) > 0 {
		newID = taskList.Tasks[len(taskList.Tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	taskList.Tasks = append(taskList.Tasks, newTask)
	fmt.Printf("Task added successfully (ID: %d)\n", newID)
}

func updateTask(taskList *TaskList, id int, newDescription string) {
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Description = newDescription
			taskList.Tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Task updated successfully (ID: %d)\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func deleteTask(taskList *TaskList, id int) {
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)
			fmt.Printf("Task deleted successfully (ID: %d)\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func markTaskStatus(taskList *TaskList, id int, status string) {
	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Status = status
			taskList.Tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Task marked as %s (ID: %d)\n", status, id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func listAllTasks(taskList TaskList) {
	for _, task := range taskList.Tasks {
		fmt.Printf("[%d] %s (Status: %s)\n", task.ID, task.Description, task.Status)
	}
}

func listTasksByStatus(taskList TaskList, status string) {
	for _, task := range taskList.Tasks {
		if task.Status == status {
			fmt.Printf("[%d] %s\n", task.ID, task.Description)
		}
	}
}
