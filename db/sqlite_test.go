package db

import (
	"testing"
)

var schema = []string{`CREATE TABLE student ( rowid INTEGER primary key AUTOINCREMENT, name TEXT, constraint unique_student_name unique (name));`,
	`CREATE TABLE class ( rowid INTEGER primary key AUTOINCREMENT, name TEXT, constraint unique_class_name unique (name));`,
	`CREATE TABLE studentClass ( rowid INTEGER primary key AUTOINCREMENT, studentId INTEGER, classId INTEGER, constraint unique_student_class_name unique (studentId, classId));`,
	`INSERT INTO student values (null, 'Neha');`,
	`INSERT INTO student values (null, 'Aditya');`,
	`INSERT INTO class values (null, 'Science');`,
	`INSERT INTO class values (null, 'Lit');`,
	`INSERT INTO studentClass values (null, 'Aditya', 'Lit');`,
	`INSERT INTO studentClass values (null, 'Aditya', 'Science');`,
	`INSERT INTO studentClass values (null, 'Neha', 'Science');`,
}

func initDB() *SqliteHandler {
	sqlHandler := InMemorySqliteHandler
	for _, stmt := range schema {

		_, err := sqlHandler.Execute(stmt)
		if err != nil {
			panic(err.Error())
		}
	}
	return sqlHandler
}

func cleanUp() {
}

type Student struct {
	Rowid int
	Name  string
}

type Class struct {
	Rowid int
	Name  string
}

type StudentClass struct {
	Rowid   int
	Student Student `db:"studentId"`
	Class   Class   `db:"classId"`
}

func TestScan(t *testing.T) {
}
