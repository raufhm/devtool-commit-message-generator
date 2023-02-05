### Commit Message Generator
This Go program generates a commit message by combining the current branch name and a user-entered string.

### Usage
1. Clone this repository to your local machine
2. Run the program by executing go run main.go
3. Enter the desired commit message when prompted
4. The program will print the complete commit message, which will be in the format branch_name: message

### Functions
The program consists of the following functions:

- getCommitMsg: prompts the user for a commit message and returns it
- cmdCurrentBranch: retrieves the current branch name and returns it, up until the ticket number
- main: combines the branch name and commit message, then runs the git commit command with the combined message

### Requirements
- Go
- Git

### Note
If there's an error, the program will panic and display an error message.