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
	
	includeHeaders bool
	method string
	body string

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
	if method == ""{
		method = http.MethodGet
	}

	if body != "" && method == ""{
		method = http.MethodPost
	}

	method = strings.ToUpper(method)

	handler, ok := methodHandlers[method]
	if !ok {
		return fmt.Errorf("unsupported HTTP method %s", method)
	}

	return handler(url, body)
}

func executeRequest(method, url, body string) error {
	var bodyReader io.Reader
	if body != ""{
		bodyReader = strings.NewReader(body)
	}

	request, err := http.NewRequest(method, url, bodyReader)
	if err!= nil{
		return err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if  err != nil{
		return err
	}

	if includeHeaders {
		for k, v := range response.Header{
			fmt.Printf("%s: %s\n", k, strings.Join(v,","))
		}
		fmt.Println()
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300{
		return fmt.Errorf("request failed :%s", response.Status)
	}

	_,err = io.Copy(os.Stdout, response.Body)

	return err
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&method, "request", "X", "GET", "HTTP method to use",)
	rootCmd.Flags().StringVarP(&body, "data", "d", "","Request body")
	rootCmd.Flags().BoolVarP(&includeHeaders, "include", "i", false, "include response headers")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


