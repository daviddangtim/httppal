/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type methodHandler func (url, body string) error 


// rootCmd represents the base command when called without any subcommands
var (
	// Flag variables
	includeHeaders bool
	method string
	body string
	outputFile string

	rootCmd = &cobra.Command{
	Use:   "httppal <url>",
	Short: "A CLI application that makes http requests and prints the response body.",
	Args: cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		return dispatchRequest(method,url, body)
	},
	}

	methodHandlers = map[string]methodHandler{
		http.MethodGet: func(url, _ string) error {
			return executeRequest(http.MethodGet, url, "")
		},

		http.MethodPost: func(url, body string) error {
			return executeRequest(http.MethodPost, url, body)
		},

		http.MethodPut: func(url, body string) error {
			return executeRequest(http.MethodPut, url, body)
		},

		http.MethodDelete: func(url, _ string) error {
			return executeRequest(http.MethodDelete, url, "")
		},
		}
)

func dispatchRequest(method, url, body string) error  {
	method = normalizeMethod(method, body)
	handler, ok := methodHandlers[method]
	if !ok {
		return fmt.Errorf("unsupported HTTP method %s", method)
	}

	return handler(url, body)
}

func normalizeMethod(method, body string) string {
	if method == ""{
		if body != ""{
			return http.MethodPost
		}
		return http.MethodGet
	}
	return strings.ToUpper(method)
}

func executeRequest(method, url, body string) error {
	request, err := requestBuilder(method, url, body)
	if err != nil{
		return err
	}

	response, err := sendRequest(request)
	if err != nil{
		return err
	}
	defer response.Body.Close()

	return handleResponse(response)
}

func requestBuilder(method, url, body string) (*http.Request, error)  {
	var bodyReader io.Reader
	if body != "" {
		bodyReader = strings.NewReader(body)
	}

	return http.NewRequest(method, url, bodyReader)
}

func sendRequest(req *http.Request) (*http.Response, error)  {
	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}

func handleResponse(resp *http.Response) error  {
	if resp.StatusCode < 200 || resp.StatusCode >= 300{
		return fmt.Errorf("request failed: %s", resp.Status)
	}
	
	if outputFile !=""{
		return writeToFile(outputFile, resp.Body)
	}

	_, err := io.Copy(os.Stdout, resp.Body)
	return err
}

func writeToFile(filename string, r io.Reader) error  {
	fi, err := os.Stat(filename)
	if err == nil && fi.IsDir(){
		return fmt.Errorf("%q is a directory", filename)
	}

	file, err := os.Create(filename)
	if err != nil{
		return err
	}
	defer file.Close()

	_,err = io.Copy(file, r)
	return err
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Write response body to file instead of stdout")
	rootCmd.Flags().StringVarP(&method, "request", "X", "GET", "HTTP method to use",)
	rootCmd.Flags().StringVarP(&body, "data", "d", "","Request body")
	rootCmd.Flags().BoolVarP(&includeHeaders, "include", "i", false, "include response headers")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


