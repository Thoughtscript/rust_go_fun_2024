package models

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

type WorkerModel struct {
	Uuid    string
	Time    time.Time
	Status  string
	Command string
	Output  string
}

// ----------------------------
// Worker-Specific Helpers
// ----------------------------

func (worker WorkerModel) ExecuteCommand() {
	msg := "executing"
	worker.updateStatus(msg)
	outputChannel := make(chan string)

	go func() {
		if worker.validateCommand() {
			cmd := worker.assignCommand()
			var out bytes.Buffer
			cmd.Stdout = &out
	
			var second time.Duration = 1000000000
			var seconds time.Duration = 3
			time.Sleep(second * seconds)
	
			// Check if worker stopped
			if worker.Status != "stopped" {
				err := cmd.Run()
				if err != nil {
					msg = "failed"
				} else {
					msg = "completed"
				}
				worker.log(msg, outputChannel, out.String())
			}

		} else {
			msg = "failed"
			worker.log(msg, outputChannel, "aborting: invalid command")
			
		}
	}()

	result := <-outputChannel
	fmt.Println(result)
	close(outputChannel)
}

func (worker WorkerModel) validateCommand() bool {
	switch worker.Command {
		case "a","b":
			return true
		default:
			return false
	}
	return false
}

func (worker WorkerModel) assignCommand() *exec.Cmd {
	switch worker.Command {
		case "a":
			return exec.Command("bash", "bin/a.sh")
		case "b":
			return exec.Command("bash", "bin/b.sh")
		default:
			return exec.Command(worker.Command)
	}
	return exec.Command(worker.Command)
}

func (worker WorkerModel) log(msg string, outputChannel chan string, stdoutStr string) {
	worker.Output = stdoutStr
	worker.updateStatus(msg)
	outputChannel <- worker.Uuid + " " + worker.Status + " [" + worker.Command + "] " + stdoutStr
	DeleteFromWorkerQueue(worker.Uuid)
}

func (worker WorkerModel) updateStatus(msg string) {
	worker.Status = msg
	UpdateStatusTable(worker.Uuid, msg)
	// Update worker
	AddToWorkerTable(worker)
}
