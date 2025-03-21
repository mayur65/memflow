package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	serverAddress = "localhost:8082"
	numRequests   = 1000
	concurrency   = 1
)

var failCount int
var mu sync.Mutex

func main() {

	var wg sync.WaitGroup

	startTime := time.Now()

	failCount = 0

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		go func(workedId int) {
			defer wg.Done()

			for j := 0; j < numRequests/concurrency; j++ {

				conn, err := net.Dial("tcp", serverAddress)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Fprintf(conn, "GET key\n")
				reader := bufio.NewReader(conn)
				_, err = reader.ReadString('\n') // Read response until newline
				if err != nil {
					fmt.Println("Error reading response:", err)
					failCount++
					return
				}

				// Print the server's response
				//fmt.Println("Server response:", response)
				_ = conn.Close()
			}

		}(i)

	}

	wg.Wait()
	elapsed := time.Since(startTime)

	fmt.Printf("Time elapsed: %s\n", elapsed)

	fmt.Printf("Number of failed requests: %d\n", failCount)
}

func incrementFailCount() {
	mu.Lock()
	defer mu.Unlock()
	failCount++
}
