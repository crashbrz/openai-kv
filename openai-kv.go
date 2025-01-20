package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func validateOpenAIAPIKey(apiKey string) bool {
	url := "https://api.openai.com/v1/models"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return false
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func main() {
	apiKey := flag.String("k", "", "Single OpenAI API key")
	filePath := flag.String("f", "", "Path to a file containing OpenAI API keys")
	showAll := flag.Bool("d", false, "Debug Mode - Display both valid and invalid keys")
	numThreads := flag.Int("t", 1, "Number of goroutines to use for validation")
	flag.Parse()

	if *apiKey == "" && *filePath == "" {
		fmt.Println("Usage: -k <Single OpenAI API key> or -f <path to file containing API keys>")
		return
	}

	validCount := 0
	invalidCount := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	validateKey := func(key string) {
		isValid := validateOpenAIAPIKey(key)
		mu.Lock()
		if isValid {
			fmt.Printf("\033[32m%s is valid.\033[0m\n", key)
			validCount++
		} else {
			if *showAll {
				fmt.Printf("\033[31m%s is invalid.\033[0m\n", key)
			}
			invalidCount++
		}
		mu.Unlock()
	}

	if *apiKey != "" {
		if validateOpenAIAPIKey(*apiKey) {
			fmt.Printf("\033[32m%s is valid.\033[0m\n", *apiKey)
			validCount++
		} else {
			fmt.Printf("\033[31m%s is invalid.\033[0m\n", *apiKey)
			invalidCount++
		}
		fmt.Printf("Valid keys: %d\n", validCount)
		if *showAll {
			fmt.Printf("Invalid keys: %d\n", invalidCount)
		}
		return
	}

	if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		keys := []string{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			key := strings.TrimSpace(scanner.Text())
			if key != "" {
				keys = append(keys, key)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		keyChan := make(chan string, len(keys))
		for _, key := range keys {
			keyChan <- key
		}
		close(keyChan)

		worker := func() {
			for key := range keyChan {
				validateKey(key)
			}
			wg.Done()
		}

		for i := 0; i < *numThreads; i++ {
			wg.Add(1)
			go worker()
		}
		wg.Wait()

		fmt.Printf("Valid keys: %d\n", validCount)
		if *showAll {
			fmt.Printf("Invalid keys: %d\n", invalidCount)
		}
	}
}
