package main

//controllers helps to use different service package and handles response
import (
	crsConfig "coursera/config"
	"coursera/models"
	"coursera/services"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Create: Create(config) it is an handler function which pulls data from coursera site and create resources inside DB
func Create(c crsConfig.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestParameter := models.CreateRequest{}
		var resp models.Response
		resp.Status = http.StatusOK
		resp.Message = "Resource creation sucessfull"
		if r.Method == "POST" {
			err := json.NewDecoder(r.Body).Decode(&requestParameter)
			if err != nil {
				resp.Status = http.StatusInternalServerError
				resp.Message = err.Error()
				Data, _ := json.Marshal(resp)
				JSON(c, string(Data))(w, r)
				return
			}
		}

		uri := "https://api.coursera.org/api/courses.v1?&query=" + requestParameter.Query + "&fields=description,domainTypes,instructors.v1(firstName,lastName,suffix)&includes=instructorIds&start=" + requestParameter.Offset + "&limit=" + requestParameter.Limit
		response, err := http.Get(uri)
		if err != nil {
			resp.Status = http.StatusInternalServerError
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			resp.Status = http.StatusInternalServerError
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		courseraData := models.Coursera{}
		err = json.Unmarshal(responseData, &courseraData)
		if err != nil {
			resp.Status = http.StatusInternalServerError
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		err = services.CreateCourse(c, &courseraData)
		if err != nil {
			resp.Status = http.StatusBadRequest
			resp.Message = err.Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		Data, _ := json.Marshal(resp)
		JSON(c, string(Data))(w, r)
		return

	}
}

//FetchData: FetchData(config) it is a handler function which fetch data on queriying from db
func FetchData(c crsConfig.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		link := r.RequestURI
		domain := r.Host
		query := r.URL.Query().Get("q")
		queryMethod := r.URL.Query().Get("qm")
		start, err := strconv.Atoi(r.URL.Query().Get("start"))
		var resp models.Response
		resp.Status = http.StatusOK

		if err != nil {
			start = 0
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 10
		}
		var courses []models.Course
		switch queryMethod {
		case "course-name":
			courses, err = services.FetchCoursesByName(c, query, start, limit)
		case "author-name":
			courses, err = services.FetchCoursesByAuthor(c, query, start, limit)
		case "course-desc":
			courses, err = services.FetchCoursesByDesc(c, query, start, limit)
		default:
			courses, err = services.FetchCourses(c, start, limit)
		}
		if err != nil || len(courses) == 0 {
			resp.Status = http.StatusBadRequest
			resp.Message = errors.New("No record found").Error()
			Data, _ := json.Marshal(resp)
			JSON(c, string(Data))(w, r)
			return

		}
		l := len(courses) - 1
		resp.Size = len(courses)
		resp.Limit = limit
		if queryMethod != "" {

			if courses[0].CourseNo+1 > limit {
				resp.Links.PreviousPage = "&start=" + strconv.Itoa(courses[0].CourseNo) + "&limit=" + strconv.Itoa(limit)
			} else if len(courses) > 1 {
				resp.Links.NextPage = "&start=" + strconv.Itoa(courses[l].CourseNo+1) + "&limit=" + strconv.Itoa(limit)
			}
		}

		if courses[0].CourseNo+1 > limit {
			resp.Links.PreviousPage = domain + link + "?start=" + strconv.Itoa(courses[0].CourseNo-limit) + "&limit=" + strconv.Itoa(limit)
		} else if len(courses) > 1 {
			resp.Links.NextPage = domain + link + "?start=" + strconv.Itoa(limit-courses[l].CourseNo+1) + "&limit=" + strconv.Itoa(limit)
		}

		resp.Data = courses
		Data, _ := json.Marshal(resp)
		JSON(c, string(Data))(w, r)
		return
	}
}
