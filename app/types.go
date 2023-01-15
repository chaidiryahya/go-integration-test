package app

// Scenarios represent slice of scenario
type ScenarioInfo struct {
	ID        int
	Scenarios []Scenario
}

// Scenarion repsresent the detail of each scenario
type Scenario struct {
	Name             string      `json:"name"`
	Method           string      `json:"method"`
	Host             string      `json:"host"`
	Path             string      `json:"path"`
	ExpectedResponse interface{} `json:"expected_response"`
}

type CallAPIParam struct {
	Host   string
	Path   string
	Method string
}
