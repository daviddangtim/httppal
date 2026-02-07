# httppal

httppal is a minimal curl-like HTTP client written in Go.  
It allows you to send HTTP requests from the command line and print or save the response body.

This project is primarily a learning exercise focused on:
- Building CLIs with Cobra
- Making HTTP requests in Go
- Streaming data to stdout or files
- Structuring clean and extensible Go code

## Features

- Send HTTP requests to a URL
- Supports multiple HTTP methods
- Optional request body
- Stream response to stdout or write to a file
- Simple and predictable CLI interface

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/yourname/httppal.git
cd httppal
go build -o httppal
```

Optionally move the binary into your PATH:
```bash
mv httppal /usr/local/bin/
```
## Usage
httppal <url> [flags]


<url> is required and specifies the target URL.

[flags] are optional modifiers to control HTTP method, request body, and output.

## Flags
Flag	Long form	Description
-X	--request	HTTP method to use (GET, POST, PUT, DELETE, etc.)
-d	--data	Request body to send with POST or PUT requests
-o	--output	Write the response body to a file instead of stdout
-i	--include	Include response headers in the output
-h	--help	Show help information
Examples
```bash
Basic GET request
httppal https://example.com
```

Prints the response body to stdout.

POST request with a body
```bash
httppal https://api.example.com/users -X POST -d '{"name":"example"}'
```

Sends a POST request with the specified JSON body.

Save response to a file
```bash
httppal https://example.com/file.zip -o file.zip
```
Writes the response to file.zip instead of printing to stdout.

Include response headers
```bash
httppal https://example.com -i
```
Prints both the response headers and the body.

Using automatic method promotion

If a request body is provided without specifying a method, httppal may automatically use POST:
```bash
httppal https://api.example.com/users -d '{"name":"example"}'
```
## Error handling

Network or HTTP errors are printed to stderr.

Non-success HTTP status codes (not 2xx) will return an error.

Runtime errors do not print usage, keeping the CLI behavior clean.

```bash
This version is **focused on usage**, showing commands, flags, and practical examples.  

```
