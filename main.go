package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

//thoughts on design of my app

/* I could make a map of student key and isPresent value. I would make a slice of maps of type Student & isPresent and the 3rd value would be day of the week
It seems like I should put that all into the attendence record struct.

*/
type Student struct {
	firstName string
	lastName  string
}

type AttendenceRecord struct {
	student   Student
	isPresent bool
	
}

type ClassRoom struct {
	className       string
	studentList     []Student
	classAttendence []AttendenceRecord
}

type School struct {
	classRoomList []ClassRoom
}

func (class *ClassRoom) takeAttendence() {
	var attendance AttendenceRecord
	for _, student := range class.studentList {
		attendance.student = student
		fmt.Printf("Is %v %v present? Y/N\n", student.firstName, student.lastName)
		input, err := readLine()
		if err != nil{
			fmt.Println("Invalid input")
			continue
		} 
		attendance.isPresent, err = yesOrNo(input)

		if err != nil {
			fmt.Println("Invalid input")
			continue
		}		
		class.classAttendence = append(class.classAttendence, attendance)

	}
	
	for _, record := range class.classAttendence {
		if record.isPresent == true {
			fmt.Printf("%v %v: Present\n", record.student.firstName, record.student.lastName)
		} else {
			fmt.Printf("%v %v: is Absent\n", record.student.firstName, record.student.lastName)
		}
		
	}
}

func (class ClassRoom) viewAttendenceRecords() {
	for _, record := range class.classAttendence {
		if record.isPresent == true {
			fmt.Printf("%v %v: Present\n", record.student.firstName, record.student.lastName)
		} else {
			fmt.Printf("%v %v: is Absent\n", record.student.firstName, record.student.lastName)
		}
		
	}
}

func (class ClassRoom) viewStudentsInClass() {
	for _, student := range class.studentList {
		fmt.Printf("%v %v, ", student.firstName, student.lastName)
	}
}

func main() {
	fmt.Println("Welcome to the ClassRoom Manager")

	var school School

	//for loop and switch statements to manage menu items
	for {
		fmt.Println("Please choose from the following: ")
		fmt.Println("Press 1 to create a new classroom, 2 to add students to a class, 3 take attendance by class, 4 to view attendence records")
		fmt.Println("5 to view student list for a class, 6 to update student info, 7 to delete a student, and 8 to exit the program.")
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

		if menuChoice != 1 && school.classRoomList == nil {
			fmt.Println("You must create at least one classroom before selecting other options.")
			continue
		}
		
		switch menuChoice {
		case 1:
			// newClass := createClassRoom()
			var newClass ClassRoom
			newClass.NameClassRoom()
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
			
			//why does our code in swith case 3 & 4 look like this?
			//Explanation: selectClassRoom func returns an int that we use as the index value for our classRoom slice that we have in school
			//so school.classRoomList[choice] gives us a classRoom instance from the slice of classRooms in school. 
			//takeAttendence and viewAttendenceRecords are method of of type ClassRoom with receiver types of *ClassRoom
			//this means we can use dot notation to call these methods and they will (but tbh they don't need to in this case) modify the classRoom isntance that
			//..we called the methods on
		case 3:
			choice := school.selectClassRoom()
			if school.classRoomList[choice].studentList == nil {
				class := school.classRoomList[choice].className
				fmt.Printf("There are no students currently assigned to class: %v\n", class)
				break
			}
			school.classRoomList[choice].takeAttendence()
			
		case 4:
			choice := school.selectClassRoom()
			if school.classRoomList[choice].studentList == nil {
				class := school.classRoomList[choice].className
				fmt.Printf("There are no students currently assigned to class: %v\n", class)
				break
			}
			school.classRoomList[choice].viewAttendenceRecords()
		case 5:
			choice:= school.selectClassRoom()
			if school.classRoomList[choice].studentList == nil {
				class := school.classRoomList[choice].className
				fmt.Printf("There are no students currently assigned to class: %v\n", class)
				break
			}
			school.classRoomList[choice].viewStudentsInClass()
			
		case 8:
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

//%d returns index value, %s is a placeholder for a string val. To explain the list in the foreach loop...
//school contains a field that is a slice of classRooms. Selection is our index to pick a specific class from the classroom slice. ClassRoom struct has field..
//that is a slice of students. So ultimate we are saying, iterate throught the student list of our chosen classroom.
func (school School) listStudentsInClass(selection int) {
	for index, student := range school.classRoomList[selection].studentList {
		fmt.Printf("%d. %s %s\n", index+1, student.firstName, student.lastName)
	}
}


//made a func for bufio for eease of use. 
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

//methods vs funcs. When you make a method you list the receiver type. it lets you do structInstance.someMethod() and only lets you use the method on that struct type
//unless you use a pointer(I think), the method will operate on a copy of the struct instance and will not modify the original struct instance
//we declared err here because the className is already declared in our ClassRoom struct that has package scope. We can't redeclare classname with := so we 
//... have to use "=" we still needed to declare err so we just did that separately. 
func (class *ClassRoom) NameClassRoom() {

	for {
		fmt.Println("Enter the name for the classroom: ")
		var err error
		class.className, err = readLine()
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		valid, err := isValidName(class.className)
		if valid {
    		break
		} else {
    		fmt.Printf("%v\n", err)
		}
	}
	return 
}

func createStudent() Student {
	var student Student

	for {
		fmt.Println("Enter first name of Student:")
		tempFirstName, err := readLine()
		valid, err := isValidName(tempFirstName)
		if err != nil || !valid {
    		fmt.Println("Invalid first name. Please try again.")
   			continue
		}

		student.firstName = tempFirstName
		break
	}

	for {
		fmt.Println("Enter last name of Student:")
		tempLastName, err := readLine()
		valid, err := isValidName(tempLastName)
		if err != nil || !valid {
    		fmt.Println("Invalid last name. Please try again.")
    		continue
		}

		student.lastName = tempLastName
		break
	}

	return student
}

/*Previously this func just returned a bool and within my program I would use if bool == false {Print invalid input}. 
This is not very user friendly because the user doesn't know what they did wrong. I've updated the func to return bool and error
so that they know how to correct their mistake.*/

func isValidName(name string) (bool, error) {
    if len(name) == 0 {
        return false, fmt.Errorf("Name cannot be empty")
    }
    // First character must be a letter
    if !unicode.IsLetter(rune(name[0])) {
        return false, fmt.Errorf("The class name cannot start with a number")
    }
    // Iterate through the rest of the name
    for _, r := range name {
        if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) && r != '\'' && r != '.' {
            return false, fmt.Errorf("Invalid name, avoid special characters.")
        }
    }
    return true, nil
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


