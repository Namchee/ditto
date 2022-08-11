package entity

// RunnerResult wraps a test runner result
type RunnerResult struct {
	Name      string
	Error     error
	Responses []*FetchResult
}
