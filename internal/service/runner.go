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
	var result []string

	err := make(chan error)
	rch := make(chan string, len(r.data.Endpoints))

	rwg := &sync.WaitGroup{}

	for _, ep := range r.data.Endpoints {
		fetcher := NewFetcher(&ep)

		rwg.Add(1)
		go r.wrapFetcher(rwg, fetcher, rch, err)
	}

	go r.cleanup(rwg, rch, err)

	if e := <-err; err != nil {
		ch <- &entity.TestResult{
			Name:  r.data.Name,
			Error: e,
			Diff:  nil,
		}
		return
	}

	for cha := range rch {
		result = append(result, cha)
	}

	diff := GetDiff(result)

	ch <- &entity.TestResult{
		Name:  r.data.Name,
		Error: nil,
		Diff:  diff,
	}
}

func (r *TestRunner) wrapFetcher(
	wg *sync.WaitGroup,
	f *Fetcher,
	ch chan<- string,
	errC chan<- error,
) {
	defer wg.Done()
	result, err := f.Fetch()

	if err != nil {
		errC <- err
		return
	}

	ch <- result
}

func (r *TestRunner) cleanup(
	wg *sync.WaitGroup,
	ch chan string,
	err chan error,
) {
	wg.Wait()
	close(ch)
	close(err)
}
