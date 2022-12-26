package model

type Event struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Source    string  `json:"source,omitempty"`
	EventTime string  `json:"eventTime,omitempty"`
	Payload   Payload `json:"payload,omitempty"`
}

type Payload struct {
	Number1 int `json:"number1,omitempty"`
	Number2 int `json:"number2,omitempty"`
	Answer  int `json:"Num1,omitempty"`
}
