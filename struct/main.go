package main

import "fmt"

type Student struct {
	ID      int
	Name    string
	Grade   float32
	Batch   int
	Section rune
}

func printStudent(stu Student) {
	fmt.Println("ID:", stu.ID)
	fmt.Println("Name:", stu.Name)
	fmt.Println("Grade:", stu.Grade)
	fmt.Println("Batch:", stu.Batch)
	fmt.Println("Section:", string(stu.Section))
}

func (stu Student) printStudent() {
	fmt.Println("ID:", stu.ID)
	fmt.Println("Name:", stu.Name)
	fmt.Println("Grade:", stu.Grade)
	fmt.Println("Batch:", stu.Batch)
	fmt.Println("Section:", string(stu.Section))
}

func (stu *Student) setGrade(x float32) float32 {
	stu.Grade = x
	return stu.Grade
}
func (stu *Student) setState(key string, value any) {
	switch key {
	case "ID":
		stu.ID = value.(int)
	case "Name":
		stu.Name = value.(string)
	case "Grade":
		switch v := value.(type) {
			case int:
				stu.Grade = float32(v)
			case float32:
				stu.Grade = v
			case float64:
				stu.Grade = float32(v)
		}
	case "Batch":
		stu.Batch = value.(int)
	case "Section":
		stu.Section = value.(rune)
	default:
		fmt.Println("Invalid key")
	}
}

func main() {
	student1 := Student{
		ID:      1,
		Name:    "John Doe",
		Grade:   90.5,
		Batch:   2020,
		Section: 'A',
	}

	// var student2 Student = Student{
	// 	ID:      2,
	// 	Name:    "Jane Doe",
	// 	Grade:   95.5,
	// 	Batch:   2021,
	// 	Section: 'B',
	// }

	student1.printStudent()
	fmt.Println(student1.setGrade(95.5))
	student1.printStudent()
	student1.setState("ID", 2)
	student1.setState("Name", "Amiyo Kumar")
	student1.setState("Grade", 99)
	student1.printStudent()

}
