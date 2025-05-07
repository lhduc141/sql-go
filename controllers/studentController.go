package controllers

import (
	"encoding/json"
	_ "fmt"
	_ "log"
	"net/http"
	"sql-go/models"
	"strconv"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := models.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := models.CreateStudent(student.Name, student.Class, student.BusID, student.Avatar, student.FeatureVector)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"message": "Student created successfully",
		"student": student.Name,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Check exist Student
	existingStudent, err := models.GetStudentByID(student.StudentID)
	if err != nil {
		http.Error(w, "Error checking student existence: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if existingStudent == nil {
		http.Error(w, "Student with this ID not found", http.StatusNotFound)
		return
	}

	if student.Name == "" || student.Class == "" || student.BusID == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err = models.UpdateStudent(student.StudentID, student.Name, student.Class, student.BusID, student.Avatar, student.FeatureVector)
	if err != nil {
		http.Error(w, "Error updating student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Student updated successfully",
		"student": student.Name,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("student_id")
	if studentID == "" {
		http.Error(w, "Student ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(studentID)
	if err != nil {
		http.Error(w, "Invalid student ID format", http.StatusBadRequest)
		return
	}

	existingStudent, err := models.GetStudentByID(id)
	if err != nil {
		http.Error(w, "Error checking student existence: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if existingStudent == nil {
		http.Error(w, "Student with this ID not found", http.StatusNotFound)
		return
	}

	// Xóa học sinh
	err = models.DeleteStudent(studentID)
	if err != nil {
		http.Error(w, "Error deleting student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Phản hồi thành công
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message":    "Student deleted successfully",
		"student_id": studentID,
	}
	json.NewEncoder(w).Encode(response)
}
