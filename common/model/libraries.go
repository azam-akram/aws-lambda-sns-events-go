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

/*
{
    "libraries": [
        {
            "id": 1,
            "name": "The central library",
            "city": "London",
            "books": [
                {
                    "id": 1,
                    "title": "C++ for professional",
                    "units": 5
                },
                {
                    "id": 2,
                    "title": "Introduction to Go lang",
                    "units": 10,
                    "attributes": {
                        "condition": "new"
                    }
                }

            ]
        },
        {
            "id": 2,
            "name": "The central library",
            "city": "Stockholm",
            "books": [
                {
                    "id": 1,
                    "title": "C++ for professional",
                    "units": 10
                },
                {
                    "id": 2,
                    "title": "Effective Java",
                    "units": 15,
                    "attributes": {
                        "demand": "medium"
                    }
                }

            ]
        }
    ]
}
*/
