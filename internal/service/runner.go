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
	ch chan<- *entity.RunnerResult,
) {
	defer wg.Done()
	var result []*entity.FetchResult

	err := make(chan error)
	rch := make(chan *entity.FetchResult, len(r.data.Endpoints))

	rwg := &sync.WaitGroup{}

	for _, ep := range r.data.Endpoints {
		fetcher := NewFetcher(ep)

		rwg.Add(1)
		go r.wrapFetcher(rwg, fetcher, rch, err)
	}

	go r.cleanup(rwg, rch, err)

	if e := <-err; e != nil {
		ch <- &entity.RunnerResult{
			Name:  r.data.Name,
			Error: e,
		}
		return
	}

	for cha := range rch {
		result = append(result, cha)
	}

	ch <- &entity.RunnerResult{
		Name:   r.data.Name,
		Result: result,
	}
}

func (r *TestRunner) wrapFetcher(
	wg *sync.WaitGroup,
	f *Fetcher,
	ch chan<- *entity.FetchResult,
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
	ch chan *entity.FetchResult,
	err chan error,
) {
	wg.Wait()
	close(ch)
	close(err)
}
