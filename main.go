package main

import (
	"fmt"
	"go-101/lesson"
)

type Lesson struct {
	Name     string
	Function func()
	Enabled  bool
}

func main() {

	lessons := []Lesson{
		{
			"Hello World",
			lesson.HelloWorld,
			false,
		},
		{
			"Values",
			lesson.Values,
			false,
		},
		{
			"Variables",
			lesson.Variables,
			true,
		},
	}

	separator := "----------------------------------------"
	lessonIndex := 1
	for _, lesson := range lessons {
		if !lesson.Enabled {
			continue
		}

		fmt.Println(separator)
		fmt.Printf("Lesson %d: %s\n", lessonIndex, lesson.Name)
		fmt.Println(separator)
		fmt.Println("Output:")
		lesson.Function()
		fmt.Println()
		lessonIndex++
	}
}
