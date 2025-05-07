package models

import (
	"database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Kết nối cơ sở dữ liệu
var db *sql.DB

type Student struct {
	StudentID     int    `json:"student_id"`
	Name          string `json:"name"`
	Class         string `json:"class"`
	BusID         int    `json:"bus_id"`
	Avatar        string `json:"avatar"`
	FeatureVector string `json:"feature_vector"`
}

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "user:user@tcp(localhost:3306)/studentdb")
	if err != nil {
		log.Fatal("Không thể kết nối đến cơ sở dữ liệu: ", err)
	}
}

func GetAllStudents() ([]Student, error) {
	var students []Student
	rows, err := db.Query("SELECT * FROM Student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.StudentID, &student.Name, &student.Class, &student.BusID, &student.Avatar, &student.FeatureVector); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func CreateStudent(name, class string, busID int, avatar, featureVector string) error {
	_, err := db.Exec("INSERT INTO Student (name, class, bus_id, avatar, feature_vector) VALUES (?, ?, ?, ?, ?)", name, class, busID, avatar, featureVector)
	return err
}

func UpdateStudent(id int, name, class string, busID int, avatar, featureVector string) error {
	_, err := db.Exec("UPDATE Student SET name = ?, class = ?, bus_id = ?, avatar = ?, feature_vector = ? WHERE student_id = ?",
		name, class, busID, avatar, featureVector, id)
	return err
}

func DeleteStudent(id string) error {
	_, err := db.Exec("DELETE FROM Student WHERE student_id = ?", id)
	return err
}
