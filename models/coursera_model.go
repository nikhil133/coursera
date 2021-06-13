package models

type Coursera struct {
	Elements []*Element `json:"elements"`
	Linked   *Linked    `json:"linked"`
}
type Element struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Linked struct {
	Instructors []*Instructor `json:"instructors.v1"`
}

type Instructor struct {
	FullName  string `json:"fullName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
