package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to specify the URL and amount of queries/second in the command line.\n   flood-test <ip/route> <reqs/s>")
		return
	}

	url := os.Args[1]
	amount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid 2nd argument. Number of workers expected")
		return
	}

	var counter int64

	for range amount {
		ticker := time.NewTicker(time.Second)
		go func() {
			for {
				<-ticker.C
				resp, err := http.Get(url)

				if err != nil || resp.StatusCode != 200 {
					fmt.Println("Error on a request!")
				}
				atomic.AddInt64(&counter, 1)
			}
		}()
		time.Sleep(time.Second / time.Duration(amount))
	}

	fmt.Println("Doing", amount, "requests per second.")

	totalFailed := 0
	totalDone := 0

	ticker2 := time.NewTicker(time.Second)
	for {
		<-ticker2.C

		count := atomic.LoadInt64(&counter)
		atomic.StoreInt64(&counter, 0)

		totalDone += amount
		failed := amount - int(count)
		totalFailed += failed

		fmt.Print("\rTotal failed: ", totalFailed, " | Errors/s:", failed, " | Total:", totalDone)
	}
}
