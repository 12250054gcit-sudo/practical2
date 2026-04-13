package model

import (
	"fmt"
	"myapp/dataStore/postgres"
)

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES ($1, $2, $3, $4);"
const queryGetUser = "SELECT*FROM student WHERE stdid=$1;"
const queryUpdateUser = "UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid;"
const queryDeleteUser = "DELETE FROM student WHERE stdid=$1 RETURNING stdid;"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

func (s *Student) Read() error {
	row := postgres.Db.QueryRow(queryGetUser, s.StdId)
	return row.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

func (s *Student) Update(oldId int64) error {
	return postgres.Db.QueryRow(queryUpdateUser, s.StdId, s.FirstName, s.LastName, s.Email, oldId).Scan(&s.StdId)
}

func (s *Student) Delete(oldId int64) error {
	err := postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllStuds() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM student;")
	if getErr != nil {
		return nil, getErr
	}

	//create a slice of type student
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {//checking the error while extracting value from database
			return students, dbErr
		}

		students = append(students, s)
	}
	rows.Close()
	return students, nil
}

type Course struct {
	Cid        int64  `json:"cid"`
	CourseName string `json:"coursename"`
}

const queryCourse = "INSERT INTO course(cid, coursename) VALUES ($1, $2)"
const queryGetCourse = "SELECT * FROM course WHERE cid=$1"
const queryUpdateCourse = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid"
const queryDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid"

func (c *Course) Create() error {
	// Changed to match the actual DB column (e.g., 'course_name')
	// CORRECT: Using 'coursename'
	_, err := postgres.Db.Exec("INSERT INTO course(cid, coursename) VALUES($1, $2)", c.Cid, c.CourseName)
	return err
}

func (c *Course) Read() error {
	row := postgres.Db.QueryRow("SELECT * FROM course WHERE cid=$1", c.Cid)
	return row.Scan(&c.Cid, &c.CourseName)
}
func (c *Course) Update(oldID int64) error {
	err := postgres.Db.QueryRow("UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid", c.Cid, c.CourseName, oldID).Scan(&c.Cid)
	fmt.Println(err)
	return err
}
func (c *Course) Delete() error {
	if err := postgres.Db.QueryRow("DELETE FROM course WHERE cid=$1 RETURNING cid", c.Cid).Scan(&c.Cid); err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM course")
	if getErr != nil {
		return nil, getErr
	}
	courses := []Course{}
	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.Cid, &c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil

}
