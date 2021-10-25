package service

import (
	"sync"

	"github.com/Namchee/ditto/internal/entity"
)

// TestRunner runs and executes test concurrently
type TestRunner struct {
	data *entity.TestData
}

// NewTestRunner creates a new test runner that executes test concurrently
func NewTestRunner(data *entity.TestData) *TestRunner {
	return &TestRunner{
		data: data,
	}
}

// RunTest executes the test and returns the test result
func (r *TestRunner) RunTest(
	wg *sync.WaitGroup,
	ch chan<- *entity.TestResult,
) {
	defer wg.Done()
	var errChannel = make(chan error)
	var channels [](chan string)
	var result []string

	for _, ep := range r.data.Endpoints {
		fetcher := NewFetcher(&ep)
		channel := make(chan string)

		channels = append(channels, channel)

		go r.wrapFetcher(fetcher, channel, errChannel)
	}

	if e := <-errChannel; e != nil {
		ch <- &entity.TestResult{
			Name:  r.data.Name,
			Error: e,
			Diff:  []int{},
		}
	}

	for _, cha := range channels {
		res := <-cha

		result = append(result, res)
	}

	diff := GetDiff(result)

	ch <- &entity.TestResult{
		Name:  r.data.Name,
		Error: nil,
		Diff:  diff,
	}
}

func (r *TestRunner) wrapFetcher(
	f *Fetcher,
	ch chan<- string,
	errC chan<- error,
) {
	result, err := f.Fetch()

	if err != nil {
		errC <- err
		return
	}

	ch <- result
}
