package main

import (
	"fmt"
	"log"
	"net/http"
	"sql-go/controllers"
	"sql-go/models"
)

func main() {
	models.InitDB() //Connect

	http.HandleFunc("/students", controllers.GetStudents)
	http.HandleFunc("/students/create", controllers.CreateStudent)
	http.HandleFunc("/students/update", controllers.UpdateStudent)
	http.HandleFunc("/students/delete", controllers.DeleteStudent)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
