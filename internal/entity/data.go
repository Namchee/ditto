package entity

// TestData defines services to be tested
type TestData struct {
	Name      string     `json:"name" validate:"required"`
	Endpoints []Endpoint `json:"endpoints" validate:"required,gt=1,dive"`
}

// Endpoint contains endpoint information to be tested
type Endpoint struct {
	Host    string                 `json:"host" validate:"required,ip|url|hostname"`
	Method  string                 `json:"method,omitempty" validate:"oneof=GET POST PUT PATCH DELETE"`
	Query   map[string]interface{} `json:"query"`
	Body    map[string]interface{} `json:"body"`
	Headers map[string]string      `json:"headers"`
	Timeout int                    `json:"timeout" validate:"number"`
}
