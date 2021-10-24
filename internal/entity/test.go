package entity

// TestResult wraps a test runner result
type TestResult struct {
	Name  string
	Error error
	Diff  []int
}
