# Go Example

This folder contains a working example API client for [Minim's API](https://my.minim.co/api_doc) written
using the Go programming language.

Use this project as a starting point for building a Go application that interacts with Minim's API. The command in
`cmd/minim-example/main.go` serves as a starting point with a few example API calls. 

Check the [Minim API Documentation](https://my.minim.co/api_doc) for a complete reference of available API calls. 

## Getting Started

1. Install a [recent version of Go](https://golang.org/doc/install)
2. Generate an [Application ID and Secret Key for your account on Minim](https://my.minim.co/api_keys)
3. Clone the repository or [download the latest code as a zip](https://github.com/MinimSecure/minim-api-examples/archive/main.zip)
4. Open a terminal and enter the directory you cloned or extracted the project to, then enter the `go/` directory
7. Run the example command with go build
   
   Replace `<APPLICATION ID>` with your Application ID and `<SECRET>` with your Secret from the Minim portal.
   ```
   go run cmd/minim-example/main.go <APPLICATION ID> <SECRET>
   ```
