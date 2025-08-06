# CLI Todo App

A simple and lightweight command-line todo application written in Go. Manage your tasks directly from the terminal.

## Features

  * **Add tasks:** Easily add new tasks to your list.
  * **Delete tasks:** Remove tasks you've completed or no longer need.
  * **Toggle status:** Mark tasks as done or undone.
  * **List tasks:** View all your tasks with their current status.
  * **Clear list:** Erase your entire task list.
  * **Persistent storage:** Your tasks are saved locally to a JSON file, so they're always there when you need them.

## Installation

To install this tool and use it from any terminal, follow these steps:

1.  **Build the executable:**
    Navigate to the root directory of this project and run the following command to compile your Go code into an executable named `main`.

    ```bash
    go build main.go
    ```

2.  **Move to the `PATH`:**
    Move the `todo` executable to a directory that is included in your system's `PATH`. A common choice is `/usr/local/bin/`. You may need to use `sudo` for this.

    ```bash
    sudo mv main /usr/local/bin/
    ```

3. **Rename as `todo` or however you want to name it:**

    ```bash
    sudo mv /usr/local/bin/main /usr/local/bin/todo
    ```

4.  **Verify installation:**
    Open a new terminal window and simply type `todo`. The application's usage instructions should appear.

## Usage

The application uses a simple command structure: `todo <command> [arguments]`.

### Commands

| Command                   | Alias | Description                                                | Example                                     |
| ------------------------- | ----- | ---------------------------------------------------------- | ------------------------------------------- |
| `add <task name>`         | `a`   | Adds a new task. Use quotes for multi-word tasks.          | `todo add buy milk`                         |
| `list`                    | `l`   | Displays all tasks in the list.                            | `todo list`                                 |
| `delete <id\|task name>`  | `d`   | Deletes a task by its index number or name.                | `todo delete 2` or `todo delete buy milk`   |
| `toggle <id\|task name>`  | `t`   | Marks a task as done or undone by its index number or name.| `todo toggle 1` or `todo toggle buy milk`   |
| `clear`                   | `c`   | Clears all tasks from the list.                            | `todo clear`                                |

## Data Storage

Your task list is automatically saved to a file named **`tasks.json`** inside a hidden directory called `.todo` in your user's home folder. This ensures your data is persistent across sessions and kept separate from the application's executable.

  * **Linux/macOS:** `~/.todo/tasks.json`
  * **Windows:** `%USERPROFILE%\.todo\tasks.json`
