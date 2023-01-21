package app

// TestCaseInfo represent slice of scenario
type TestCaseInfo struct {
	ID        int
	TestCases []TestCase
}

// TestCase repsresent the detail of each TestCase
type TestCase struct {
	Name             string           `json:"name"`
	Method           string           `json:"method"`
	Host             string           `json:"host"`
	Header           []TestCaseHeader `json:"header"`
	Path             string           `json:"path"`
	Payload          string           `json:"payload"`
	ExpectedResponse interface{}      `json:"expected_response"`
}

type TestCaseHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CallAPIParam struct {
	Host    string
	Path    string
	Method  string
	Payload string
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}
