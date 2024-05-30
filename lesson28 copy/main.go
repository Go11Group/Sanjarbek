package main

import (
	"fmt"
	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {
	new := model.Course{Name: "Fullstack", Field: "Javascript"}
	
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	st := postgres.NewStudentRepo(db)
	
	users, err := st.GetAllStudents()
	if err != nil {
		panic(err)
	}
	
	user, _ := st.GetByID(users[0].ID)
	
	fmt.Println(users)
	
	fmt.Println(user)
	
	cr := postgres.NewCourseRepo(db)
	courses, er := cr.GetAllCourses()
	if er != nil {
		panic(er)
	}
	
	fmt.Println(courses[0].Id)
	
	
	student := model.User{Name: "Ali", Age: 30, Gender: "m", Course: courses[0].Id}

	err = cr.CreateCourse(new)
    if err != nil {
        panic(err)
    }

	err = st.Create(student)
	
	
	
}
