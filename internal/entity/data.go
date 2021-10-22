package entity

type TestData struct {
	Name      string     `json:"query_name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Host    string                 `json:"host"`
	Method  string                 `json:"method"`
	Params  map[string]string      `json:"params"`
	Query   map[string]interface{} `json:"query"`
	Body    map[string]interface{} `json:"body"`
	Headers map[string]string      `json:"headers"`
	Timeout int                    `json:"timeout"`
}
