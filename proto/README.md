These are the protocol buffer structure files that define the methods you can call on the server, the inputs and outputs. 
The files are at the root directory, so that they can be shared with the client application and kept in sync.

When you extend these, you need to run `make gen-proto` at root to update the Go files.