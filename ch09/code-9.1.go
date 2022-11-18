package main

import (
	"log"
	"os"
)

type Jobs struct {
	Cmd string
	*log.Logger
}

func CreateJob(cmd string) *Jobs {
	return &Jobs{cmd, log.New(os.Stderr, "New Job ", log.Ldate)}
}

func main() {
	CreateJob("SampleRun").Print("Initiating . . .")
}

/*
OUTPUT
New Job 2022/06/24 Initiating . . .
*/
