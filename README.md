# dieGo

dieGo is a simple command-line task management tool written in Golang. It allows you to add, delete, list, complete, and view details of tasks.

## Getting Started

To use dieGo, you need to have Golang installed on your system.

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/diegoquinfa/dieGo.git
   cd dieGo
   ```

2. Build the application:

   ```bash
   go install
   ```

3. Run dieGo:

   ```bash
   dieGo
   ```

## Usage

Usage: dieGo [-u] \<command\> [arg]

Commands:
  - `add:`        Add a new task
  - `delete:`     Delete a task
  - `list:`      List all tasks
  - `complete:`   Mark a task as complete
  - `details:`    View task details

Options:
  `-u`          Mark the task as urgent (optional)

Example:
  - Add a new urgent task:
    ```bash
    dieGo add -u "Buy groceries"
    ```
  - Add a new normal task:
    ```bash
    dieGo add "Drink water"
    ```

  - List all tasks:
    ```bash
    dieGo list
    ```

  - Complete a task:
    ```bash
    dieGo complete 1
    ```
    
  - task details:
    ```bash
    dieGo details 1
    ```
    
## Add task:
<div align="center">
<img src="https://res.cloudinary.com/drvoywub5/image/upload/v1697928712/github/dieGo/dhfgohtbhyuiwq10ohtq.png" width=80%>
</div>

## Complete taks:
<div align="center">
<img src="https://res.cloudinary.com/drvoywub5/image/upload/v1697928712/github/dieGo/krrhyuovjk9jzbjudev5.png" width=80%>
</div>

## Task details:
<div align="center">
<img src="https://res.cloudinary.com/drvoywub5/image/upload/v1697929664/github/dieGo/ezbptwva5akl4qo4x1sr.png" width=80%>
</div>


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
