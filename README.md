# Task Tracker CLI

Task Tracker CLI is a simple command-line application for managing tasks. It allows users to add, update, delete, and list tasks, as well as mark tasks as in-progress or done.

## Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as in-progress or done
- List all tasks
- List tasks by status

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/mrdankuta/rdmpsh-task-tracker-cli.git
   ```
3. Navigate to the project directory:
   ```
   cd task-tracker-cli
   ```
4. Build the application:
   ```
   go build -o task-cli
   ```

## Usage

The general syntax for using the CLI is:

```
./task-cli <command> [arguments]
```

Available commands:

1. Add a task:
   ```
   ./task-cli add <description>
   ```

2. Update a task:
   ```
   ./task-cli update <id> <new description>
   ```

3. Delete a task:
   ```
   ./task-cli delete <id>
   ```

4. Mark a task as in-progress:
   ```
   ./task-cli mark-in-progress <id>
   ```

5. Mark a task as done:
   ```
   ./task-cli mark-done <id>
   ```

6. List all tasks:
   ```
   ./task-cli list
   ```

7. List tasks by status:
   ```
   ./task-cli list <status>
   ```
   Where `<status>` can be "todo", "in-progress", or "done".

## Data Storage

Tasks are stored in a JSON file named `tasks.json` in the same directory as the executable. This file is created automatically when you add your first task.

## Error Handling

The application includes basic error handling for file operations and invalid commands. Error messages will be displayed in the console.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements

This project was created as part of the [Roadmap.sh Backend Project series](https://roadmap.sh/projects/task-tracker).
