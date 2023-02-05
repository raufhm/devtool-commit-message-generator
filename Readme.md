### Commit Message Generator
This Go program generates a commit message by combining the current branch name and a user-entered string.

### Usage
1. Clone this repository to your local machine
2. Run the program by executing go run main.go
3. Enter the desired commit message when prompted
4. The program will print the complete commit message, which will be in the format `branch_name: message`

### Functions

- Generates a commit message by reading from the standard input.
- Retrieves the current branch in the Git repository.
- Commits to the current branch with the generated commit message.

### Requirements
- Go
- Git

### Contributions
Contributions are welcome. Feel free to open an issue or create a pull request.
