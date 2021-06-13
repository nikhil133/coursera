package services

//service package handles the business logic of an app
import (
	crsConfig "coursera/config"
	"coursera/models"
	"coursera/repository"
	"strings"
	"sync"
)

func CreateCourse(c crsConfig.Config, courseraData *models.Coursera) error {
	courseSize := len(courseraData.Elements)
	coursesData := make([]models.Course, courseSize)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, value := range courseraData.Elements {
			coursesData[i].CourseName = strings.TrimSpace(strings.ToLower(value.Name))
			coursesData[i].CourseDescription = strings.TrimSpace(strings.ToLower(value.Description))
		}

	}()

	go func() {
		defer wg.Done()
		for i := 0; i < courseSize; i++ {
			if i < len(courseraData.Linked.Instructors) {
				coursesData[i].AuthorFirstName = strings.TrimSpace(strings.ToLower(courseraData.Linked.Instructors[i].FirstName + courseraData.Linked.Instructors[i].LastName))
				coursesData[i].AuthorLastName = strings.TrimSpace(strings.ToLower(courseraData.Linked.Instructors[i].FirstName))
				coursesData[i].AuthorFullName = strings.TrimSpace(strings.ToLower(courseraData.Linked.Instructors[i].LastName))
			}

		}
	}()
	wg.Wait()

	return repository.CreateCourse(c, coursesData)
}

func FetchCourses(c crsConfig.Config, start int, limit int) ([]models.Course, error) {
	return repository.FetchCourses(c, start, limit)
}

func FetchCoursesByName(c crsConfig.Config, cname string, start int, limit int) ([]models.Course, error) {
	return repository.FetchCoursesByDesc(c, cname, start, limit)
}

func FetchCoursesByAuthor(c crsConfig.Config, aname string, start int, limit int) ([]models.Course, error) {
	return repository.FetchCoursesByAuthor(c, aname, start, limit)
}

func FetchCoursesByDesc(c crsConfig.Config, desc string, start int, limit int) ([]models.Course, error) {
	return repository.FetchCoursesByDesc(c, desc, start, limit)
}
