package main

import (
	"goHello/src/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("start work")

	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("due to interrupt")
			os.Exit(1)
		}
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
