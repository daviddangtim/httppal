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
