package entity

// RunnerResult wraps a test runner result
type RunnerResult struct {
	Name      string         `json:"name"`
	Error     error          `json:"error"`
	Responses []*FetchResult `json:"responses"`
}
