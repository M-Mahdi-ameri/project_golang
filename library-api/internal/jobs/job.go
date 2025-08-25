package jobs

type Job struct {
	ID      int
	Payload string
}

type ReportJob struct {
	RequestID int
}
