package model

type Event struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Source    string  `json:"source,omitempty"`
	EventTime string  `json:"eventTime,omitempty"`
	Payload   Payload `json:"payload,omitempty"`
}

type Payload struct {
	Numbers []int `json:"numbers,omitempty"`
	Sum     int   `json:"sum,omitempty"`
}
