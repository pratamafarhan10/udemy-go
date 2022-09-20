package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Code, Name string
	Sks        int
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	college := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"WD", "Web Development", 3},
				{"ML", "Machine Learning", 3},
				{"SE", "Software Engineering", 3},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"UI", "User Interface Design", 3},
				{"UX", "User Experience Design", 3},
				{"BP", "Basic Programming", 3},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{"AL", "Basic Algorithm", 3},
				{"DS", "Data Structures", 3},
				{"PM", "Project Management", 3},
			},
		},
	}

	err := tpl.Execute(os.Stdout, college)
	if err != nil {
		log.Println(err)
	}
}
