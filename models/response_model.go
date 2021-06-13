package models

type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Course `json:"data"`
	Links   Link     `json:"links"`
	Size    int      `json:"size"`
	Limit   int      `json:"limit"`
}

type Link struct {
	NextPage     string `json:"next_page"`
	PreviousPage string `json:"previous_page"`
}
