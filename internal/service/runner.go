package service

import "github.com/Namchee/ditto/internal/entity"

type testRunner struct {
	data *entity.TestData
}

// NewTestRunner creates a new test runner that executes endpoint concurrently
func NewTestRunner(data *entity.TestData) *testRunner {
	return &testRunner{
		data: data,
	}
}

// RunTest executes the test and returns the test result
func (r *testRunner) RunTest() *entity.TestResult {
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
		return &entity.TestResult{
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

	return &entity.TestResult{
		Name:  r.data.Name,
		Error: nil,
		Diff:  diff,
	}
}

func (r *testRunner) wrapFetcher(
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
