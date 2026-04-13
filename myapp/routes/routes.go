package routes

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	// creating mux router
	router := mux.NewRouter()
	//student route
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/all", controller.GetAllStuds) //the default method will be GET
	router.HandleFunc("/student/{sid}", controller.GetStudent).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")

	// course routes
	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/all", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")

	log.Println("Application running successfully..")
	// start my server
	log.Fatal(http.ListenAndServe(":8080", router))
}
