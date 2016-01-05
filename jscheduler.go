package main

import (
	"./jscheduler"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"flag"
	"time"
)


func main() {
	// Get command line args
	var pid int
	var interval int
	threadSpecs := jscheduler.NewThreadSpecArgList()

	flag.IntVar(&pid, "pid", 0, "The pid of the monitored java process")
	flag.IntVar(&interval, "interval", 3000, "Time to wait between polling jstack in milliseconds")
	flag.Var(&threadSpecs, "thread specs", "The threads which need to be rescheduled and the scheduling options")

	if help := flag.Bool("help", false, "Display usage information"); *help {
		fmt.Println(flag.Usage())
	}

	threadCount := make(map[string]int)

	// Print thread occurrence count on CTRL-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		printThreadCount(threadCount)
		os.Exit(1)
	}()


	for {
		// Get thread dump
		threadDump, err := jscheduler.GetJstackThreadDump(os.Getenv("JAVA_HOME"), pid)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Parse thread dump
		threads, err := jscheduler.ParseThreadDump(threadDump)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Filter and adjust thread specs
		jscheduler.AdjustThreadSpecs(&threads, &threadSpecs.Get())

		// Set Thread affinities and priorities
		jscheduler.RescheduleThreadGroup(threads)

		time.Sleep(interval * time.Millisecond)
	}

	/*
		out, _ := jscheduler.GetThreadDump("34768")

		fmt.Println(out)

		parsed, _ := jscheduler.ParseThreadDump(out)

		fmt.Println(parsed)

		for _,v := range *parsed {
			fmt.Println(v.Name)
		}

		//fmt.Println(len(*parsed.threads))
	*/
}

func printThreadCount(threadCount map[string] int) {
	keys := make([]string, 0, len(threadCount))
	for k := range threadCount {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d", k, threadCount[k])
	}
}