package constant

const (
	// TestTemplate is a template to use when automatically creating a new test file
	TestTemplate = `{
	"name": "SampleTest",
	"endpoints": [
		{
			"host": "https://www.google.com",
			"method": "GET",
			"query": {
				"q": "github"
			},
			"headers": {
				"Accept": "application/json"
			},
			"timeout": 3
		},
		{
			"host": "https://www.github.com"
		}
	]
}`
)
