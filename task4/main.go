package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type OutputResult struct {
	totalTimeReq    int64
	averageTimeReq  int64
	maxResponseTime int64
	minResponseTime int64
	countMissResp   int
}

type InputParameters struct {
	url        string
	countOfReq int
	timeOut    int
}

func sendReq() (out OutputResult, err error) {
	parameters := initParameters()
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var totalReqTime int64
	var countReq int64
	var allTimeResp []int64
	var countMissResp int

	for i := 0; i < parameters.countOfReq; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			client := http.Client{
				Timeout: time.Duration(parameters.timeOut),
			}
			_, err := client.Get(parameters.url)
			checkErr(err)

			timeReq := time.Since(start).Nanoseconds()
			mutex.Lock()
			if err, ok := err.(net.Error); ok && err.Timeout() {
				countMissResp++
			} else if err != nil {
				return
			} else {
				totalReqTime += timeReq
				countReq++
			}
			allTimeResp = append(allTimeResp, timeReq)
			mutex.Unlock()
		}()
	}
	wg.Wait()

	if countReq == 0 {
		return out, errors.New("all requests failed")
	}

	out = OutputResult{totalReqTime, totalReqTime / countReq,
		getMaxTimeResp(allTimeResp), getMinTimeResp(allTimeResp), countMissResp}
	return
}

func initParameters() InputParameters {
	urlFlag := flag.String("url", "", "url")
	countOfReqFlag := flag.String("count", "", "count of request")
	timeOutFlag := flag.String("timeOut", "", "timeOut")
	flag.Parse()

	checkParameters(*urlFlag, *countOfReqFlag, *timeOutFlag)

	url := *urlFlag
	countOfReq, err := strconv.Atoi(*countOfReqFlag)
	checkErr(err)
	timeOut, err := strconv.Atoi(*timeOutFlag)
	checkErr(err)

	return InputParameters{url: url, countOfReq: countOfReq, timeOut: timeOut}
}

func checkParameters(url, countOfReq, timeOut string) {
	if url == "" {
		panic("-url flag has to be specified")
	} else if countOfReq == "" {
		panic("-count flag has to be specified")
	} else if timeOut == "" {
		panic("-timeOut flag has to be specified")
	}

}

func getMaxTimeResp(allTimeResp []int64) (maxTimeResp int64) {
	for i := 0; i < len(allTimeResp); i++ {
		if allTimeResp[i] > maxTimeResp {
			maxTimeResp = allTimeResp[i]
		}
	}
	return
}

func getMinTimeResp(allTimeResp []int64) (minTimeResp int64) {
	minTimeResp = allTimeResp[0]
	for i := 0; i < len(allTimeResp); i++ {
		if allTimeResp[i] < minTimeResp {
			minTimeResp = allTimeResp[i]
		}
	}
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printRespInConsole(response OutputResult) {
	fmt.Println("Time during which all requests worked:", response.totalTimeReq)
	fmt.Println("Average request time:", response.averageTimeReq)
	fmt.Println("Maximum response time:", response.maxResponseTime)
	fmt.Println("Minimum response time:", response.minResponseTime)
	fmt.Println("Number of missed responses:", response.countMissResp)
}

func main() {
	response, err := sendReq()
	printRespInConsole(response)
	checkErr(err)
}
