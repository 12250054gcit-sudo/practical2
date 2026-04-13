package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	//stud to store student data send by the client
	var stud model.Student

	decoder := json.NewDecoder(r.Body)
	//storing the json data of a student in stud variable
	err := decoder.Decode(&stud)
	if err != nil {
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.Body.Close()
	// fmt.Println(stud)
	dbErr := stud.Create()
	if dbErr != nil {
		// w.Write([]byte(dbErr.Error()))
		httpResp.ResponseWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}

	//no error
	// w.Write([]byte("successfully stored"))
	httpResp.ResponseWithJson(w, http.StatusCreated, map[string]string{"Status": "Student Added"})
}

// helper function to convert string to int
func getUserId(userId string) (int64, error) {
	intId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, err
	}
	return intId, nil
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	//to store student info retrieved from db
	var stud model.Student
	myMap := mux.Vars(r)
	sid := myMap["sid"]
	stdid, idErr := getUserId(sid)

	if idErr != nil {
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	stud = model.Student{StdId: stdid}

	r.Body.Close()
	// fmt.Println(stud)
	dbErr := stud.Read() //calling read method in the model component
	if dbErr != nil {
		// w.Write([]byte(dbErr.Error()))
		httpResp.ResponseWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}

	//no error
	// w.Write([]byte("successfully stored"))
	httpResp.ResponseWithJson(w, http.StatusOK, stud)
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	//to store student info retrieved from db
	var stud model.Student
	//extract sid from old data
	myMap := mux.Vars(r)
	sid := myMap["sid"]
	stdid, idErr := getUserId(sid)

	if idErr != nil {
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	//extract new data from the request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)
	if err != nil {
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	// fmt.Println(stud)
	dbErr := stud.Update(stdid) //calling read method in the model component

	if dbErr != nil {
		switch dbErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, "student not found")
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, dbErr.Error())
		}
	} else {
		httpResp.ResponseWithJson(w, http.StatusOK, stud)

	}
}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	myMap := mux.Vars(r)
	sid := myMap["sid"]
	stdid, idErr := getUserId(sid)
	if idErr != nil {//error while converting string to int 
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdid}
	if err := s.Delete(stdid); err != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.ResponseWithJson(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	students, getErr := model.GetAllStuds()
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.ResponseWithJson(w, http.StatusOK, students)
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	//course to store course data send by the client
	var course model.Course

	decoder := json.NewDecoder(r.Body)
	//storing the json data of a course in course variable
	err := decoder.Decode(&course)
	if err != nil {
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	r.Body.Close()
	// fmt.Println(course)
	dbErr := course.Create()
	if dbErr != nil {
		// w.Write([]byte(dbErr.Error()))
		httpResp.ResponseWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}

	//no error
	// w.Write([]byte("successfully stored"))
	httpResp.ResponseWithJson(w, http.StatusCreated, map[string]string{"Status": "Course Added"})
}

// helper function to convert string to int
func getCourseId(userId string) (int64, error) {
	intId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, err
	}
	return intId, nil
}

type courseUpdater interface {
	Update(int64) error
}

func callCourseUpdate(course *model.Course, oldCid int64) error {
	if updater, ok := any(course).(courseUpdater); ok {
		return updater.Update(oldCid)
	}
	return errors.New("model.Course does not implement Update(int64)")
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	//to store course info retrieved from db
	var course model.Course
	myMap := mux.Vars(r)
	cid := myMap["cid"]
	courseId, idErr := getCourseId(cid)

	if idErr != nil {
		// w.Write([]byte(err.Error()))
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	course = model.Course{Cid: courseId}

	r.Body.Close()
	// fmt.Println(course)
	dbErr := course.Read() //calling read method in the model component
	if dbErr != nil {
		// w.Write([]byte(dbErr.Error()))
		httpResp.ResponseWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}

	//no error
	// w.Write([]byte("successfully stored"))
	httpResp.ResponseWithJson(w, http.StatusOK, course)
}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hello")
	oldCid := mux.Vars(r)["cid"]
	courseId, idErr := getCourseId(oldCid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var course model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		// fmt.Println("err  here")
		httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	// fmt.Println(course)
	updateErr := course.Update(courseId)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.ResponseWithError(w, http.StatusNotFound, "course not found")
		default:
			httpResp.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.ResponseWithJson(w, http.StatusOK, course)
	}

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	courseId, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := model.Course{Cid: courseId}
	if err := c.Delete(); err != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.ResponseWithJson(w, http.StatusOK, map[string]string{"Status": " Deleted"})

}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, getErr := model.GetAllCourses()
	if getErr != nil {
		httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		return
	}
	httpResp.ResponseWithJson(w, http.StatusOK, courses)

}
