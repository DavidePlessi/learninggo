# What cli command go have ?
- `go build`: compile a bunch of go source code files
- `go run`: compile and executes one or two files
- `go fmt`: formats all the code in each file in the current directory
- `go install`: compile and 'installs' a package
- `go get`: downloads the raw source code of someone else's package
- `go test`: runs any test associated with the current project
# What does 'package main' mean?
Package in go is like a project or a workspace.  
So package can have many related file inside of it each file ending with a file extension of go.  
The requirement is that at the very first line of each file must declare the package that it belongs to.
## Why do we call it main
Inside of go there are two types of packages:
  - **Executable**: generate a file that we can run
  - **Reusable**: Code used as 'helpers'. Good place to put reusable logic.  
 
 How do we know that a package is executable or reusable? It's actually the name of the package that you use that 
 determines whether you are making one or another, so specifically the word main.  
 So the word package main is one that we use only when we are making a package that we want to spit out some runnable 
 file.  
 The file must have a func called **main**.
# What does 'import' means?
 The import statement is used to give our package access to some code is written inside another package.  
 **Packages ->** https://golang.org/pkg/
# What's that 'func' thing?
Is just like function in other programming language.
# How is the main.go file organized?
The pattern is :
- _Package declaration_
- _Import other packages that we need_
- _Declare functions, tell Go to do things_