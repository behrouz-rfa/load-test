package logger

import (
	"fmt"
	"github.com/apoorvam/goterminal"
	clr "github.com/fatih/color"
	"load-test/internal/domain"
	"os"
	"time"
)

// ShowTimeProgress is going to show a countdown timer base on requests duration
func ShowTimeProgress(requestDurationInSeconds int) {
	// get an instance of writer
	writer := goterminal.New(os.Stdout)
	counter := requestDurationInSeconds
	go func(rd int, bar *goterminal.Writer) {
		d1, _ := time.ParseDuration(fmt.Sprintf("%vs", rd))
		timeout := time.After(d1)
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-timeout:

				return
			case <-ticker.C:
				if counter > 0 {
					clr.Set(clr.FgGreen)
					counter -= 1
					fmt.Fprintf(writer, "Sending requests in ... %d second\n", counter)
					// write to terminal
					writer.Print()
					time.Sleep(time.Second)
					// clear the text written by previous write, so that it can be re-written.
					writer.Clear()
					clr.Unset()
				} else {
					fmt.Println("please Wait...")
				}
			}
		}
	}(requestDurationInSeconds, writer)

	writer.Reset()

}

// PrintResult out the result into console
func PrintResult(status domain.Status, counter *domain.Counter) {
	if status.RequestsCounter == 0 {
		status.RequestsCounter = 1
	}
	status.AverageTime = status.TotalDuration / time.Duration(status.RequestsCounter)
	fmt.Printf("=========================== Results ========================================\n")
	fmt.Printf("\n")
	fmt.Printf(`Total   Request`)
	fmt.Printf("\t\t")
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", status.RequestsCounter)
	clr.Unset()
	fmt.Printf("Fastest Request\t\t")
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", status.MinTime)
	clr.Unset()
	fmt.Printf("Average Request\t\t")
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", status.AverageTime)
	clr.Unset()
	fmt.Printf("Slowest Request\t\t")
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", status.MaxTime)
	clr.Unset()
	fmt.Printf("Total   Error\t\t")
	clr.Set(clr.FgRed)
	fmt.Printf(" %v \n", status.ErrorCount)
	clr.Unset()
	fmt.Printf("=========================== Status ========================================\n")
	fmt.Printf("\n")
	fmt.Printf(`Status Code         Count     `)
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf(`1XX`)
	fmt.Printf(`                  `)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", counter.Get("1xx"))
	clr.Unset()
	fmt.Printf("2XX")
	fmt.Printf(`                  `)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", counter.Get("2xx"))
	clr.Unset()
	fmt.Printf("3XX")
	fmt.Printf(`                  `)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", counter.Get("3xx"))
	clr.Unset()
	fmt.Printf("4XX")
	fmt.Printf(`                  `)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", counter.Get("4xx"))
	clr.Unset()
	fmt.Printf("5XX")
	fmt.Printf(`                  `)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", counter.Get("5xx"))
	clr.Unset()
	fmt.Printf("Others")
	fmt.Printf(`               `)
	clr.Set(clr.FgGreen)
	fmt.Printf(" %v \n", counter.Get("0xx"))
	clr.Unset()
	fmt.Printf(`― ― ― ― ― ― ― ― ― ― ―― ― ― ― ― ― ― ― ―`)
	fmt.Printf("\n")

}
