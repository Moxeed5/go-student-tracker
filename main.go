package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/*important note. When you pass by value, you are createing a copy. When you pass by reference, you are passing the mem address of the object rather than a copy.
What does this mean?
It means if I create a classRoom like var newClass ClassRoom and then I use a func that accepts Classroom type as an arg, and I pass in newClass,
I am passing by value. I am making a copy of newClass. So any modifications that the function does to the ClassRoom arg won't be reflected in my original
newClass var. And the classRoom arg that I provide only has scope within the function.

So, if I want to modify the original newClass var, I need to pass by reference using a pointer. My func should be modified to accept a ClassRoom pointer rather than
a ClassRoom data type. you do this by using "*"before the data type. Now there isn't a copy of newClass being made. I am passing the mem address of the newClass
var that I declared earlier and any manipulations done in the func will be reflected in my newClass var.

Something to consider is that I could also pass by value and set the original var = to the return val of a func. For example newClass = someFunc(newClass)
Here I am passing by value, and using a copy of the newClass var as the arg for somefunc. I modifying the original newClass var by setting it equal to the
return val of somefunc(newClass). So I was able to pass by value, and also modify my original var without need to pass by reference. An issue with this is that
when you are working with huge amounts of data, creating a copy by passing by value is expensive relaitve to passing a mem address with a pointer. These are my
in the moment thoughts so I'll prob need to update this...*/

//declaring struct types
type Student struct {
	firstName string
	lastName  string
}

type AttendenceRecord struct {
	student   Student
	isPresent bool
	Records   []time.Weekday
}

type ClassRoom struct {
	className       string
	studentList     []Student
	classAttendence []AttendenceRecord
}

type School struct {
	classRoomList []ClassRoom
}

func main() {
	fmt.Println("Welcome to the ClassRoom Manager")

	var school School

	//for loop and switch statements to manage menu items
	for {
		fmt.Println("Please choose from the following: ")
		fmt.Println("Press 1 to create a new classroom, 2 to add students to a class, 3 take attendance by class and 4 to exit the program.")
		menuChoiceStr, err := readLine()
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		menuChoice, err := strconv.Atoi(menuChoiceStr)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		switch menuChoice {
		case 1:
			newClass := createClassRoom()
			school.classRoomList = append(school.classRoomList, newClass)

		case 2:
			selection := school.selectClassRoom()

			for {
				newStu := createStudent()
				school.classRoomList[selection].studentList = append(school.classRoomList[selection].studentList, newStu)

				school.listStudentsInClass(selection)

				fmt.Println("Would you like to add another student? (Y/N)")
				addAnotherStudent, err := readLine()
				if err != nil {
					fmt.Println("Error reading input. Please try again.")
					continue
				}

				shouldContinue, err := yesOrNo(addAnotherStudent)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				if !shouldContinue {
					break
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
//prints the classRooms from the classRoom slice in School struct
func (school *School) selectClassRoom() int {
	//this loop is similar to foreach in c#. It returns index of each element and every element within a slice or array or map
	for index, class := range school.classRoomList {
		fmt.Printf("%d. %s\n", index+1, class.className)
	}
	for {
		fmt.Println("Select the class")

		selectionStr, err := readLine()
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue //if user input is invalid, we skip an interation of this loop and start at the top which would be "Select the Class"
		}
		
		//readLine returns a string to we import strconv and use func Atoi to convert a string to int
		//readLine also returns an error and we use if statements to check if error is not nil, if the selection is within the range of our slice..
		//..and then we return the int selection -1 bc we added 1 to the index value so that it would start at 1 when we printed index to the user in the ..
		//"foreach" loop

		selection, err := strconv.Atoi(selectionStr) 
		if err != nil {
			fmt.Println("Invalid input, please enter the class number to make a selection")
			continue
		} else if selection <= 0 || selection > len(school.classRoomList) {
			fmt.Println("Invalid selection")
			continue
		} else {
			return selection - 1
		}
	}
}

func (school *School) listStudentsInClass(selection int) {
	for index, student := range school.classRoomList[selection].studentList {
		fmt.Printf("%d. %s %s\n", index+1, student.firstName, student.lastName)
	}
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
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
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

	for {
		fmt.Println("Enter first name of Student:")
		tempFirstName, err := readLine()
		if err != nil || !isValidName(tempFirstName) {
			fmt.Println("Invalid first name. Please try again.")
			continue
		}
		student.firstName = tempFirstName
		break
	}

	for {
		fmt.Println("Enter last name of Student:")
		tempLastName, err := readLine()
		if err != nil || !isValidName(tempLastName) {
			fmt.Println("Invalid last name. Please try again.")
			continue
		}
		student.lastName = tempLastName
		break
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


