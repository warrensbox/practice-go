package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

type Host struct {
	ipaddr  string
	updated bool
	status  bool
	date    string
}

var wg = sync.WaitGroup{}

func main() {

	dirname := os.Args[1]

	percentage := 1

	if len(os.Args) == 0 {
		fmt.Println("Arg cannot be empty")
		os.Exit(1)
	}

	//open file
	file, errOpen := os.Open(dirname)
	if errOpen != nil {
		fmt.Printf("Unable to open file: %s\n", errOpen)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hostsAddr []Host

	for scanner.Scan() {
		var host Host
		fmt.Println(scanner.Text())
		host.ipaddr = scanner.Text()
		host.status = false
		host.updated = false
		hostsAddr = append(hostsAddr, host)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Unable to read lines on file")
		os.Exit(1)
	}

	ch := make(chan *Host, len(hostsAddr))

	for i := 0; i < len(hostsAddr)*percentage; i++ {
		wg.Add(1)
		go update(hostsAddr[i], ch)
	}

	go func(ch chan<- *Host) {
		defer close(ch)
		wg.Wait()
	}(ch)

	for val := range ch {
		fmt.Println("Host : ", val.ipaddr)
		fmt.Println("Status : ", val.status)
		fmt.Println("Updated : ", val.updated)
		fmt.Println("Date : ", val.date)
		fmt.Println("-----")
	}

	select {
	case <-ch:
		fmt.Println("Done", <-ch)
	}

}

func update(host Host, ch chan<- *Host) {
	defer wg.Done()

	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("The date is %s\n", out)

	host.ipaddr = host.ipaddr + "X"
	host.status = true
	host.updated = true
	host.date = string(out)

	ch <- &host
}
