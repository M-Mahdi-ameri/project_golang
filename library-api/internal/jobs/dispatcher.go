package jobs

import (
	"fmt"
)

var JobQueue chan Job

var ReportQueue chan ReportJob

func InitDispatcher(bufferSize int) {
	JobQueue = make(chan Job, bufferSize)
}

func StartWorker(id int) {
	go func() {
		for job := range JobQueue {
			fmt.Printf("worker %d started job %d with payload: %s\n", id, job.ID, job.Payload)

			fmt.Printf("worker %d finished job %d\n", id, job.ID)
		}
	}()

}
func InitReportDispatcher(bufferSize int) {
	ReportQueue = make(chan ReportJob, bufferSize)

}
func StartReportWorker(id int) {
	go func() {
		for job := range ReportQueue {
			fmt.Printf("Report worker %d started request %d\n", id, job.RequestID)

			fmt.Printf("Report worker %d finished request %d (generated report)\n", id, job.RequestID)
		}
	}()

}
