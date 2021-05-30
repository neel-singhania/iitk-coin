package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Student struct {
	id     int
	rollno int
	name   string
	//I created a struct with a struct to select the rows in the table and add data.
}

func main() {
	os.Remove("students.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("students.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("students.db created")

	studentsDatabase, _ := sql.Open("sqlite3", "./students.db") // Open the created SQLite File
	defer studentsDatabase.Close()                              // Defer Closing the database
	createTable(studentsDatabase)                               // Create Database Tables

	// INSERT RECORDS
	insertStudent(studentsDatabase, 190538, "Neelabh Singhania")
	insertStudent(studentsDatabase, 190666, "Palak Khandelwal")

	// DISPLAY INSERTED RECORDS
	displayStudents(studentsDatabase)
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"rollno" integer,
		"name" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, rollno int, name string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(rollno, name) VALUES (?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(rollno, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var newStudent Student
		//var program string
		row.Scan(&newStudent.id, &newStudent.rollno, &newStudent.name)
		log.Println("Student: ", newStudent.rollno, " ", newStudent.name)
	}
}
