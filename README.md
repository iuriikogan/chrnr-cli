## CHRNR-CLI is a command-line interface (CLI) application that creates a specified number of files with random data, and then "churns" them, or modifies them, at a specified interval this is useful to test the efficiency of change block tracking and consistency in backup and recovery of an application which simulates a stateful workload in Kubernetes.

* Here's a breakdown of the code:

- The `runCommand` function is the main function that gets executed when the CLI command is run. It first loads the configuration, creates a directory for the test files, and logs some information about the configuration. 

- It then calculates the number of files to create based on the size of the PVC and the size of each file. It creates a WaitGroup and a cron job that logs a message every 2 minutes.

- It then launches a goroutine for each file to be created, using the `CreateFileWithRandomData` function from the `utils` package. It waits for all the goroutines to finish, then logs the number of files created, their size, and the time it took.

- After all files are created, it starts a ticker to churn the files at the specified interval. It launches a goroutine that waits for the ticker to tick, then churns a percentage of the files using the `ChurnFiles` function from the `utils` package. It also waits for 120 seconds before logging a message. This continues until the churn duration has elapsed.

- The `main` function simply executes the root command and logs any errors that occur.

## Makefile: 

* Here's a breakdown of the directives:

- GO := go: This sets the variable GO to the string go, which is the command to run the Go compiler.

- BINARY := chrnr-cli: This sets the variable BINARY to the string chrnr-cli, which is the name of the binary that will be produced.

- build:: This is a target that builds the binary. It runs the go build command, which compiles the main.go file and outputs the binary to the ./bin directory.

- run:: This is a target that runs the main.go file using the go run command.

- test:: This is a target that runs all the tests in the project with verbose output and a coverage report.

- benchmark:: This is a target that runs all the benchmark tests in the project and generates CPU and memory profiles.

- .PHONY: build run test benchmark: This line declares build, run, test, and benchmark as phony targets. Phony targets are ones that do not correspond to actual files; instead, they are just a name for a recipe to be executed.

- clean:: This is a target that removes the generated binary, memory profile, and test files. It uses the rm -rf command to remove files and directories.
