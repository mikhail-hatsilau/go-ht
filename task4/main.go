package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"
)

const (
	request_num             = 1
	max_request_num         = 1000
	time_out        float64 = 3
	max_time_out    float64 = 30
	nanoToMilli             = float64(1000000)
)

var (
	average, min, max float64
	tardyResponses    int
	total             time.Duration
	responseDurations []time.Duration
	parameters        struct {
		urls            flags
		requestQuantity int
		timeOut         float64
	}
)

type flags []string

func (flag *flags) String() string {
	return "Empty URL"
}

func (flag *flags) Set(url string) error {
	*flag = append(*flag, url)
	return nil
}

func main() {

	flag.Var(&parameters.urls, "url", "URL (e.g. google.com), it's necessary")
	flag.IntVar(&parameters.requestQuantity, "num", request_num, "requests quantity")
	flag.Float64Var(&parameters.timeOut, "timeout", time_out, "time out value, second")
	flag.Parse()
	if flag.NFlag() == 0 {
		printUsage()
	}

	urls := parameters.urls
	requestQuantity := parameters.requestQuantity
	timeOut := parameters.timeOut

	if len(parameters.urls) == 0 {
		fmt.Println("Unable to proceed: empty URL.")
		os.Exit(1)
	}
	if requestQuantity > max_request_num || timeOut > max_time_out {
		fmt.Println("You ain't need so much )")
		os.Exit(1)
	}
	if requestQuantity <= 0 || timeOut <= 0 {
		fmt.Println("The value gotta be positive.")
		os.Exit(1)
	}

	for _, url := range urls {
		for i := 0; i < requestQuantity; i++ {
			go letsGo(url)
		}
	}
	time.Sleep(time.Duration(timeOut * float64(len(urls)) * math.Pow10(9)))

	responseQuantity := len(responseDurations)
	average = round(float64(total.Nanoseconds())/float64(responseQuantity)/
		nanoToMilli, 0.001)
	min, max = getMinMax()
	min = round(min/nanoToMilli, 0.001)
	max = round(max/nanoToMilli, 0.001)
	tardyResponses = requestQuantity*len(urls) - responseQuantity
	printResults()
}

func printUsage() {
	fmt.Println("Available flags are:")
	flag.PrintDefaults()
	os.Exit(1)
}

func letsGo(url string) {
	uri := "https://" + url
	req, err := http.NewRequest(
		"POST", uri, bytes.NewBuffer([]byte(`{"title":"Hello there!"}`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Custom-Header", "MyHeader")
	client := &http.Client{}
	responseStart := time.Now()
	resp, err := client.Do(req)
	responseDuration := time.Since(responseStart)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("Can't process the request '%s': status %d \n",
			req.Host, resp.StatusCode)
	}
	time.Sleep(time.Duration(parameters.timeOut * math.Pow10(9)))
	total += responseDuration
	responseDurations = append(responseDurations, responseDuration)
}

func round(x, unit float64) float64 {
	if x > 0 {
		return float64(int64(x/unit+0.5)) * unit
	}
	return float64(int64(x/unit-0.5)) * unit
}

func getMinMax() (float64, float64) {
	var min, max int64
	for _, v := range responseDurations {
		min = v.Nanoseconds()
	}
	for _, v := range responseDurations {
		if v.Nanoseconds() < min {
			min = int64(v)
		}
		if v.Nanoseconds() > max {
			max = int64(v)
		}
	}
	return float64(min), float64(max)
}

func printResults() {
	fmt.Printf("----------------Initial data---------------\n"+
		"-> URLs:\t\t\t%s\n"+
		"-> request quantity:\t\t%8d\n"+
		"-> time out:\t\t\t%8.1f s",
		parameters.urls, parameters.requestQuantity, parameters.timeOut)

	fmt.Printf("\n------------------Results------------------\n"+
		"-> total work time:\t\t%8.3f ms\n"+
		"-> average response duration:\t%8v ms\n"+
		"-> max response duration:\t%8v ms\n"+
		"-> min response duration:\t%8v ms\n"+
		"-> tardy responses amount:\t%8v\n",
		round(float64(total)/nanoToMilli, 0.001),
		round(float64(average), 0.001),
		round(float64(max), 0.001),
		round(float64(min), 0.001),
		tardyResponses)
}
