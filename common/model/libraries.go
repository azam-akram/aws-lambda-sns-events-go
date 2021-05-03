package model

type Libraries struct {
	Libraries []Library `json:libraries,omitempty`
}

type Library struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty`
	City  string `json:"city,omitempty"`
	Books []Book `json:book,omitempty`
}

type Book struct {
	ID         int         `json:"id,omitempty"`
	Title      string      `json:"title,omitempty`
	Units      int         `json:"units,omitempty"`
	Attributes interface{} `json:attributes`
}
