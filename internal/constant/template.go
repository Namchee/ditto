package constant

const (
	// TestTemplate is a template to use when automatically creating a new test file
	TestTemplate = `{
	"name": "%s",
	"endpoints": [
		{
			"host": "https://jsonplaceholder.typicode.com/comments",
			"method": "GET",
			"query": {
				"postId": 1
			},
			"headers": {
				"Accept": "application/json; charset=UTF-8"
			},
			"timeout": 3
		},
		{
			"host": "https://jsonplaceholder.typicode.com/comments",
			"method": "GET",
			"query": {
				"postId": 1
			},
			"headers": {
				"Accept": "application/json; charset=UTF-8"
			},
			"timeout": 3
		}
	]
}`
	DefaultTestName = "SampledDittoTest"
)
