package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func main() {
	student := Student{ID: 1, Name: "Nichola Saputra", GPA: 10.0}
	fmt.Println(student.Name)
	student.graduate()
	fmt.Println(student.Name)
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.KOM"
}
