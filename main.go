package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		fmt.Println("Press 1 to create a new classroom, 2 to add students to a class, 3 take attendence by class and 4 to exit the program.")
		var menuChoice int
		fmt.Scanln(&menuChoice)

		switch menuChoice {
		case 1:
			newClass := createClassRoom()
			school.classRoomList = append(school.classRoomList, newClass)

		case 2:
			roomSelection := school.selectClassRoom()
		
			for {
				newStu := createStudent()
				roomSelection.studentList = append(roomSelection.studentList, newStu)
		
				var shouldContinue bool
				for {
					fmt.Println("Would you like to add another student? (Y/N)")
					var addAnotherStudent string
					fmt.Scan(&addAnotherStudent)
		
					var err error
					shouldContinue, err = yesOrNo(addAnotherStudent)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						break // break out of the inner loop
					}
				}
		
				if !shouldContinue {
					break // break out of the outer loop
				}
			}
					
		case 3: 
			//createStudent()
		case 4:
			return

		default: 
			fmt.Println("Invalid choice! Please select from the available options")
		}
	}
	
}

// when you write a func specifically for a struct, it's a method. Methods syntax is a bit different from a regular func. see below: 
//the receiver type is School. This func can only be used on the data type School. Classroom is return val. Params would be in the usual place

func (school *School)selectClassRoom() *ClassRoom {
	
	for index, class := range school.classRoomList {
		fmt.Printf("%d. %s\n", index + 1, class.className)
	}
	var selection int
	for {
	fmt.Println("Select the class that you would like to add Student to")
	
	_, err := fmt.Scan(&selection)
	flushStdin()
	 // scan returns the number of items successfully scanned and and error val. I can use _ if I don't care about # of items scanned
	 if err != nil {
		fmt.Println("Invalid input, please enter the class number to make a selection")
		continue //invalid choice was made, continue takes us to the top of the loop which reprompts the user to select a class
	 } else if selection <= 0 || selection > len(school.classRoomList) {
		 fmt.Println("Invalid selection")
		 continue
	 } else {
		for index, student := range school.classRoomList[selection-1].studentList{
			 fmt.Printf("%d. %s %s\n", index+1, student.firstName, student.lastName)
		}
		break
	 }
	 
	}
	
	return &school.classRoomList[selection-1]
	
}

func flushStdin() {
    var dummy string
    fmt.Scanln(&dummy)
}


func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createClassRoom() ClassRoom {
	var class ClassRoom

	for {
		fmt.Println("Enter the name for the classroom: ")

		name, err := readLine()
		if err != nil || len(name) == 0 {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		if isValidName(name) {
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
	flushStdin()
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
		flushStdin()
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
        if !unicode.IsLetter(r) && !unicode.IsSpace(r) && r != '\'' {
            return false
        }
    }
    return true
}


func yesOrNo(inputString string) (bool, error) {
	switch strings.ToLower(inputString) {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("%v is not a valid option.", inputString)
	}
}


