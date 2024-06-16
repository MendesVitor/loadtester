package cmd

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
}

var url string
var totalRequests int
var concurrency int

func worker(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Request failed: %v", err)
		results <- Result{StatusCode: 0, Duration: time.Since(start)}
		return
	}
	defer resp.Body.Close()
	duration := time.Since(start)
	results <- Result{StatusCode: resp.StatusCode, Duration: duration}
}

var loadTestCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "A simple load testing tool",
	Run: func(cmd *cobra.Command, args []string) {
		results := make(chan Result, totalRequests)
		var wg sync.WaitGroup

		start := time.Now()
		for i := 0; i < totalRequests; i++ {
			if i%concurrency == 0 && i != 0 {
				wg.Wait()
			}
			wg.Add(1)
			go worker(url, &wg, results)
		}
		wg.Wait()
		close(results)
		duration := time.Since(start)

		status200Count := 0
		otherStatusCount := 0
		var totalDuration time.Duration

		for result := range results {
			if result.StatusCode == http.StatusOK {
				status200Count++
			} else {
				otherStatusCount++
			}
			totalDuration += result.Duration
		}

		fmt.Printf("Total time taken: %v\n", duration)
		fmt.Printf("Total requests made: %d\n", totalRequests)
		fmt.Printf("Requests with status 200 (OK): %d\n", status200Count)
		fmt.Printf("Requests with other statuses: %d\n", otherStatusCount)
	},
}

func init() {
	rootCmd.AddCommand(loadTestCmd)
	loadTestCmd.Flags().StringVar(&url, "url", "", "URL to test")
	loadTestCmd.Flags().IntVar(&totalRequests, "requests", 1, "Number of total requests")
	loadTestCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Number of concurrent requests")

	if err := loadTestCmd.MarkFlagRequired("url"); err != nil {
		log.Fatalf("url flag is required: %v", err)
	}
	if err := loadTestCmd.MarkFlagRequired("requests"); err != nil {
		log.Fatalf("requests flag is required: %v", err)
	}
	if err := loadTestCmd.MarkFlagRequired("concurrency"); err != nil {
		log.Fatalf("concurrency flag is required: %v", err)
	}
}
