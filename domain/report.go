package domain

import (
	"context"
	"fmt"
)

//Report represent every the statistics report
type Report struct {
	CourseID          int64    `json:"courseID"`
	CourseTitle       string   `json:"courseTitle"`
	NumberUsers       int      `json:"numberUsers"`
	LessonsRead       int      `json:"lessonsRead"`
	MaxMark           int      `json:"maxMark"`
	MinMark           int      `json:"minMark"`
	AvgMark           float64  `json:"avgMark"`
	MostSearched      []string `json:"mostSearched"`
	MostViewedLessons []Lesson `json:"lessons"`
}

type ReportClient interface {
	GenerateReport(ctx context.Context, courseID int64) (Report *Report, err error)
}

// asserts Client implements the ReportClient interface
var _ ReportClient = (*Client)(nil)

// GenerateReport : Generates a course report
func (c Client) GenerateReport(ctx context.Context, courseID int64) (report *Report, err error) {

	r := Report{}
	r.CourseID = courseID

	course, err := c.GetCourse(ctx, courseID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	r.CourseTitle = course.Name
	r.MaxMark = 10
	r.MinMark = 2
	r.AvgMark = 7.5
	r.NumberUsers = 125
	r.LessonsRead = 500

	// lessons, err := c.GetAllLessonsByCourseID(ctx, courseID)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, err
	// }
	//Get five random lessons
	//rand := rand.Intn(len(lessons)-1) + 1
	// for index, lesson := range lessons {
	// 	if index < 5 {
	// 		r.MostViewedLessons = append(r.MostViewedLessons, lesson)
	// 	}
	// }
	r.MostSearched = []string{"Permisos", "FullScreen", "Activity", "Gradle", "AsyncTask"}

	mockLesson := Lesson{
		LessonID: int64(1),
		CourseID: int64(1),
		Title:    "Instalación de las herramientas necesarias",
		Text:     "Descarga de las herramientas necesarias para seguir el curso.",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000026/c1l1_e9ojcw.png",
		Position: 0,
	}

	mockLesson2 := Lesson{
		LessonID: int64(2),
		CourseID: int64(1),
		Title:    "El primer proyecto Android",
		Text:     "Como crear nuestro primer proyecto Android.",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000089/project_snnkng.png",
		Position: 1,
	}

	mockLesson3 := Lesson{
		LessonID: int64(3),
		CourseID: int64(1),
		Title:    "Capturar el click de un botón",
		Text:     "Como utilizar buttons en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000028/c1l2_dk87zd.png",
		Position: 2,
	}

	mockLesson4 := Lesson{
		LessonID: int64(4),
		CourseID: int64(1),
		Title:    "Los controles RadioGroup y RadioButton",
		Text:     "Como utilizar RadioGroups y Radiobuttons en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561478361/radio_button_krw1we.png",
		Position: 3,
	}

	mockLesson5 := Lesson{
		LessonID: int64(5),
		CourseID: int64(1),
		Title:    "Control CheckBox",
		Text:     "Como utilizar CheckBoxes en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561478361/checkbox_ansqpb.png",
		Position: 4,
	}

	r.MostViewedLessons = append(r.MostViewedLessons, mockLesson)
	r.MostViewedLessons = append(r.MostViewedLessons, mockLesson2)
	r.MostViewedLessons = append(r.MostViewedLessons, mockLesson3)
	r.MostViewedLessons = append(r.MostViewedLessons, mockLesson4)
	r.MostViewedLessons = append(r.MostViewedLessons, mockLesson5)

	return &r, nil
}
