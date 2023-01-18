package main

import (
	"flag"
	"fmt"
	clr "github.com/fatih/color"
	"load-test/internal/application"
	"load-test/internal/domain"
	"load-test/internal/logger"
	"load-test/internal/utils"
	"load-test/internal/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	NUM_PARALLEL             int    //number of concurrent worker
	requestURL               string //url request
	requestMethod            string //Request method (default "GET")
	headerValues             string // header values separated with ';'
	requestDurationInSeconds int    //  Request duration in Second
	requestTimeOut           int    //Request time out in milli seconds (default 2000)
	help                     bool   // print Help
)

func init() {
	flag.IntVar(&NUM_PARALLEL, "c", 5, "Number of concurrent requests")
	flag.StringVar(&requestMethod, "m", "GET", "Request method")
	flag.StringVar(&headerValues, "h", "", "header values separated with ';'")
	flag.StringVar(&requestURL, "u", "", "URL to run against")
	flag.IntVar(&requestDurationInSeconds, "d", 5, "Request duration")
	flag.IntVar(&requestTimeOut, "to", 2000, "Request time out in milli seconds")
	flag.BoolVar(&help, "help", false, "know more about command")
}

func main() {

	// create channel for get interrupt sing from the user
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	//Pars all falg
	flag.Parse()
	if !flag.Parsed() {
		log.Fatalln("[Info] Command line flags parsing failed, Please check the input")
	}
	//pring help if user wants to show the help
	if help {
		fmt.Println("Usage: command [<flags>] <url>")
		flag.VisitAll(func(flag *flag.Flag) {
			fmt.Println("\t-"+flag.Name, "\t", flag.Usage, "(Default value = "+flag.DefValue+")")
		})
		return
	}

	//check the url flag
	if len(requestURL) == 0 {
		fmt.Println("Usage: command [<flags>] <url>")
		flag.VisitAll(func(flag *flag.Flag) {
			fmt.Println("\t-"+flag.Name, "\t", flag.Usage, "(Default value = "+flag.DefValue+")")
		})
		return
	}

	run(sigChannel)
}

// run the application
func run(sigChannel chan os.Signal) {

	//create header for the request
	requestHeader := utils.RequestHeader(headerValues)

	// create the request parameter
	requestParams := domain.CreateRequestParams(requestURL, requestMethod, requestHeader)

	//create the chana for getting all status and details from the request
	statusChan := make(chan *domain.Status, NUM_PARALLEL)

	//create config
	config := domain.NewAPIConfig(NUM_PARALLEL, requestDurationInSeconds, requestTimeOut, statusChan, requestParams)

	//print some information about the request and url
	fmt.Printf("Load Test running for %vs on this url: ", requestDurationInSeconds)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", requestURL)
	clr.Unset()
	fmt.Printf(" %v Number of Concurrent connections!\n", NUM_PARALLEL)

	//create new http client forom utils
	client := utils.NewFastClient(config.TimeOut)
	//counter i
	counter := &domain.Counter{}

	//create new repository
	request := worker.NewWorker(config, client, counter)

	app := application.New(request)

	// progressbar configuration
	logger.ShowTimeProgress(requestDurationInSeconds)

	// runt the workers here in concurrent
	for i := 0; i < NUM_PARALLEL; i++ {
		go func() {
			app.CreateRequest()
		}()
	}

	//this part noting just for shoing the results from the request
	responseCounter := 0
	status := domain.NewStatus()

	//in this for loop we wait to get data from chanel
	//also we listen to interrupt signal form use wit ctrl + c to interrupt the loading test
	for responseCounter < NUM_PARALLEL {
		select {
		case <-sigChannel:
			config.Stop()
			os.Exit(0)
		case s := <-statusChan:
			status.RequestsCounter += s.RequestsCounter
			status.ErrorCount += s.ErrorCount
			status.TotalDuration += s.TotalDuration
			status.MaxTime = s.MaxTime
			status.MinTime = s.MinTime
			status.TotalResponseSize = s.TotalResponseSize
			responseCounter++
		}
	}

	// check how many request we send if ==  0 show the message
	if status.RequestsCounter == 0 {
		fmt.Println("[Info] No request found")
		return
	}

	//finish the result when program exit
	defer func() {
		logger.PrintResult(status, counter)
	}()
}
