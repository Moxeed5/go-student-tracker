package main

import (
	"fmt"
	"time"
	"unicode"
)

type Student struct {
	firstName string
	lastName string
}

type AttendenceRecord struct {
	student Student
	isPresent bool
	Records []time.Weekday

}

type ClassRoom struct {
	className string
	studentList []Student
	classAttendence []AttendenceRecord
}

type School struct {
	classRoomList []ClassRoom
}



func main() {
	fmt.Println("Welcome to the ClassRoom Manager")

	var school School 
	for {
		fmt.Println("Please choose from the following: ")
		fmt.Println("Press 1 to create a new classroom, 2 to add students to a class, 3 to add a student to a class, and 4 to exit the program.")
		var menuChoice int
		fmt.Scan(&menuChoice)

		switch menuChoice {
		case 1:
			newClass := createClassRoom()
			school.classRoomList = append(school.classRoomList, newClass)

		case 2:
			var roomSelection ClassRoom = selectClassRoom(school)
			for{
				newStu := createStudent()
				roomSelection.studentList = append(roomSelection.studentList, newStu)

				var addAnotherStudent string
				fmt.Println("Enter Y to add another student and N to exit")
				fmt.Scan(&addAnotherStudent)
				if !validateBool(addAnotherStudent) {
					break
				}

			}
		case 3: 
			createStudent()
		case 4:
			break
		}
	}
	
}

func selectClassRoom(school School) ClassRoom {
	fmt.Println("Select the class that you would like to add Student to")
	for index, class := range school.classRoomList {
		fmt.Printf("%d. %s\n", index + 1, class.className)
	}
	var selection int
	fmt.Scan(&selection)
	return school.classRoomList[selection-1]
}

func createClassRoom() ClassRoom {
	var class ClassRoom
	var name string
	for {
	fmt.Println("Enter the name for the classroom: ")
	fmt.Scan(&name)
	if isValidName(name) == true {
		class.className = name
		break
	} else {
		fmt.Printf("%v is not a valid name\n", name)
	}
}
	return class
}

func createStudent() Student {
	var student Student
	var tempFirstName string
	var tempLastName string
	
	for {
	fmt.Printf("Enter first name of Student: \n")
	fmt.Scan(&tempFirstName)
	if isValidName(tempFirstName) {
	student.firstName = tempFirstName
	break
	} else {
		fmt.Printf("%v is not a valid name\n", tempFirstName)
	}
	}
	for {
		fmt.Printf("Enter last name of Student: \n")
		fmt.Scan(&tempLastName)
		if isValidName(tempLastName){
		student.lastName = tempLastName
		break
		} else {
			fmt.Printf("%v is not a valid name\n", tempLastName)
		}
	}

	return student

}

func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func validateBool(inputString string) bool { 
	var validatedInput bool

	for {
		
		if inputString == "y" || inputString == "Y" {
			validatedInput = true
			break
		} else if inputString == "n" || inputString == "N" {
			validatedInput = false
			break
		} else {
			fmt.Printf("%v is not a valid option.\n", inputString)
			continue
		}
	}
	return validatedInput
}


