package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	EMail string  `json:"email"`
	Age   int     `json:"age"`
	gpa   float64 `json:"gpa"`
	Year  int     `json:"year"`
}

func (s *Student) IsHonors() bool {
	return s.gpa >= 3.5
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1 and 4")
	}
	if s.gpa < 0.0 || s.gpa > 4.0 {
		return errors.New("gpa must be between 0.0 and 4.0")
	}
	return nil
}
func main() {
	//var st Student = Student{ID: "1", Name: "Supachok", EMail: "Aintharasopa_s@silpakorn.edu", Year: 3, gpa: 3.21}

	students := []Student{
		{ID: "1", Name: "Supachok", EMail: "Aintharasopa_s@silpakorn.edu", Year: 3, gpa: 3.21},
		{ID: "2", Name: "EZ", EMail: "EZ@silpakorn.edu", Year: 3, gpa: 3.78},
	}
	newStudent := Student{ID: "3", Name: "COrin", EMail: "CO2@silpakorn.edu", Year: 3, gpa: 4.00}
	students = append(students, newStudent)

	for _, student := range students {
		fmt.Printf("Honors = %v \n", student.IsHonors())
		fmt.Printf("Validate = %v \n", student.Validate())
	}

}
