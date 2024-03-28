### Spawn internal process
when you execute a command in Go using the exec package, it does create a new process internally to execute that command. The exec.Command function creates an Cmd struct, which represents a command to be run in a separate process. When you call methods like Run, Start, or CombinedOutput on this Cmd struct, Go internally creates a new process and executes the specified command in that process.

### How does it work?

Creation of Cmd Struct: You use the exec.Command function to create a Cmd struct, specifying the command to execute and any arguments.

Execution: When you call methods like Run, Start, or CombinedOutput on the Cmd struct, Go spawns a new process and executes the specified command in that process.

Communication with the Process (Optional): Depending on the method you use (Run, Start, etc.), you may interact with the running process. For example, you can wait for the process to finish (Run), start the process asynchronously (Start), or retrieve its output (CombinedOutput).

Cleanup: Once the command execution completes and the process exits, Go cleans up the resources associated with the process.