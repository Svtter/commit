package pkg

import "log"

type RunResult struct {
	errout string
	out    string
	err    error
}

// OutputError function to print formatted errors.
func (r *RunResult) OutputError() {
	log.Printf("Command finished with error: %v", r.err)
	log.Printf("stdout: %v", r.out)
	log.Fatalf("stderr: %v", r.errout)
}

// Output Just another normal output
func (r *RunResult) Output() {
	if r.err != nil {
		r.OutputError()
	} else {
		log.Print(r.out)
	}
}
