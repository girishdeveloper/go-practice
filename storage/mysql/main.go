package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Teacher struct {
	id        int64
	firstname string
	lastname  string
}

func AddTeacher(firstname string, lastname string, db *sql.DB) (int64, error) {
	insertStatement := "INSERT INTO teacher(create_time, firstname, lastname) VALUES(NOW(), ?, ?)"
	res, err := db.Exec(insertStatement, firstname, lastname)
	if err != nil {
		return 0, err
	}
	id, err1 := res.LastInsertId()
	if err1 != nil {
		return 0, err1
	}
	return id, nil
}

func GetTeacherById(id int64, db *sql.DB) (*Teacher, error) {
	teacher := Teacher{id: id}
	err := db.QueryRow("SELECT firstname, lastname FROM teacher WHERE id=?", id).Scan(&teacher.firstname, &teacher.lastname)
	if err != nil {
		return &Teacher{}, err
	}
	return &teacher, nil
}

func UpdateTeacher(id int64, newFirstname string, db *sql.DB) {
	res, err := db.Exec("update teacher set firstname=? where id=?", newFirstname, id)
	if err != nil {
		panic(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if affected == 1 {
		fmt.Println("Single row updated")
	} else {
		fmt.Println("Unexpectedly %d rows updated, expected 1", affected)
	}
}

func GetTeachers(db *sql.DB) ([]Teacher, error) {
	teachers := make([]Teacher, 0)
	rows, err := db.Query("SELECT id, firstname, lastname FROM teacher")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		teacher := Teacher{}
		err := rows.Scan(&teacher.id, &teacher.firstname, &teacher.lastname)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	if err != nil {
		return teachers, err
	}
	return teachers, nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/testcon")
	if err != nil {
		panic(err)
	}
	err1 := db.Ping()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("load sql and execute it")
	f, err2 := os.Open("./create_table.sql")
	if err2 != nil {
		panic(err2)
	}
	b, err3 := ioutil.ReadAll(f)
	if err3 != nil {
		panic(err3)
	}
	res, err4 := db.Exec(string(b))
	if err4 != nil {
		panic(err4)
	}
	fmt.Println("resource result", res)
	entryDone, err5 := AddTeacher("Girish", "Madhavan", db)
	if err5 != nil {
		panic(err5)
	}
	fmt.Println("Id is", entryDone)
	teacher, err6 := GetTeacherById(entryDone, db)
	if err6 != nil {
		panic(err6)
	}
	fmt.Println(*teacher)
	UpdateTeacher(entryDone, "Kilmish", db)
	AllTeachers, err7 := GetTeachers(db)
	if err7 != nil {
		panic(err7)
	}
	fmt.Println(AllTeachers)
	db.Close()
}
