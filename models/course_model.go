package models

type Course struct {
	TableName         struct{} `sql:"course" json:"-"`
	CourseName        string   `sql:"course_name" json:"course_name"`
	CourseDescription string   `sql:"course_description" json:"course_description"`
	AuthorFullName    string   `sql:"author_fullname" json:"author_name"`
	AuthorFirstName   string   `sql:"author_firstname" json:"author_firstname"`
	AuthorLastName    string   `sql:"author_lastname" json:"author_lastname"`
	CourseNo          int      `sql:"course_no"`
}
