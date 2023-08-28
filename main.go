package main

import (
	"fmt"
)

type Student struct {
	firstName string
	lastName string
	isPresent bool

}

type ClassRoom struct {
	studentList []Student
}



func main() {
	fmt.Println("Welcome to the student tracker program")

	var class ClassRoom

	for {
	var student = takeAttendence()

	class.studentList = append(class.studentList, student)

	fmt.Println("Enter y to continue taking attendence or n to exit")
	var exit string
	fmt.Scan(&exit)

	if exit == "n" || exit == "n" {
		break
	}

	}

	fmt.Println("Students in class:")
	for i, student := range class.studentList {
		fmt.Printf("Student %d:\n", i+1)
		fmt.Printf("First Name: %s\n", student.firstName)
		fmt.Printf("Last Name: %s\n", student.lastName)
		fmt.Printf("Is Present: %v\n", student.isPresent)
	}
	
}





func takeAttendence() Student {
	
	var student Student
	var tempPresentVar string
	fmt.Println("Enter student's first name: \n")
	fmt.Scan(&student.firstName)
	fmt.Println("Enter student's last name: \n")
	fmt.Scan((&student.lastName))
	fmt.Println("Is the student present? y/n: \n")
	fmt.Scan(&tempPresentVar)

	if tempPresentVar == "y" || tempPresentVar == "Y"{
		student.isPresent = true
	} else {
		student.isPresent = false
	}
	

	return student

}

