package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func main() {
	job := &Job{"demo", log.New(os.Stdout, "Job: ", log.Ldate)}
	job.Print("starting now...")
}
