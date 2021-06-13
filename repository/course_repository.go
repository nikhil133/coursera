package repository

//repository package helps in doing operation with db
import (
	crsConfig "coursera/config"
	"coursera/models"
	"log"
)

func CreateCourse(c crsConfig.Config, courses []models.Course) error {
	err := c.PG.Create(&courses)
	return err
}

func FetchCourses(c crsConfig.Config, start int, limit int) ([]models.Course, error) {
	var rows []models.Course
	err := c.PG.Model(&models.Course{}).
		Where("course_no>=?", start).Limit(limit).Select(&rows)
	if err != nil {
		log.Println("Course:Error couldnot retrive Values. Reasons", err.Error())

	}
	return rows, err

}

func FetchCoursesByDesc(c crsConfig.Config, desc string, start int, limit int) ([]models.Course, error) {
	var rows []models.Course
	err := c.PG.Model(&models.Course{}).
		Where("course_description=? and course_no>=?", desc, start).Limit(limit).Select(&rows)
	if err != nil {
		log.Println("Course:Error couldnot retrive Values. Reasons", err.Error())

	}
	return rows, err

}

func FetchCoursesByName(c crsConfig.Config, cname string, start int, limit int) ([]models.Course, error) {
	var rows []models.Course
	err := c.PG.Model(&models.Course{}).
		Where("course_name=? and course_no>=?", cname, start).Limit(limit).Select(&rows)
	if err != nil {
		log.Println("Course:Error couldnot retrive Values. Reasons", err.Error())

	}
	return rows, err

}

func FetchCoursesByAuthor(c crsConfig.Config, aname string, start int, limit int) ([]models.Course, error) {
	var rows []models.Course
	err := c.PG.Model(&models.Course{}).
		Where("author_fullname=? and course_no>=?", aname, start).Limit(limit).Select(&rows)
	if err != nil {
		log.Println("Course:Error couldnot retrive Values. Reasons", err.Error())

	}
	return rows, err

}
