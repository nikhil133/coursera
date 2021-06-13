package models

type CreateRequest struct {
	Query  string `json:"query"`
	Limit  string `json:"limit,omitempty"`
	Offset string `json:"start,omitempty"`
}
