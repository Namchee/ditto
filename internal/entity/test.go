package entity

// TestResult wraps a test runner result
type TestResult struct {
	Name  string
	Error error
	Diff  []int
}

// TestResult wraps a fetcher result
type FetcherResult struct {
	Status int
	Body   string
}
