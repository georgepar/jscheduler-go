package main

import (
	"./jscheduler"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"syscall"
)

func main() {
	// Get command line args

	threadCount := make(map[string]int)

	// Print thread occurrence count on CTRL-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		keys := make([]string, 0, len(threadCount))
		for k := range threadCount {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("%s: %d", k, threadCount[k])
		}
		os.Exit(1)
	}()

	for {
		// Get thread dump

		// Parse thread dump

		// Filter thread groups

		// Assign thread groups to CPU pool

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
