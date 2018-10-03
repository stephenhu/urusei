package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	DEFAULT_OUTPUT 		= "output.csv"
)

type Job struct {
	Url					string
	Origin			string
}

var jobs = make(chan Job, 2000)
var completed = make(chan bool)
var wg sync.WaitGroup

func loadJobs() {

	f, err := os.Open(DEFAULT_OUTPUT)

	if err != nil {
		log.Fatal(err)
	}

  scan := bufio.NewScanner(f)
	
	for scan.Scan() {

		line := scan.Text()

		tuple := strings.Split(line, ",")

		job := Job{
			Url: tuple[0],
			Origin: tuple[1],
		}

		jobs <- job

		//m[tuple[0]] = tuple[1]

		//download(tuple[0], tuple[1])
		
	}

} // loadJobs

/*
func dispatchJobs() {

	for {
		
		select {
		case j <-jobs:
			go download(j.Url, j.Origin)
		case <-completed:

		}

	}

} // dispatchJobs
*/

func main() {

	loadJobs()

	log.Println(len(jobs))
	wg.Add(2)

	go era()
	go era()

	wg.Wait()

} // main
