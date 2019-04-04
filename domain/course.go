package domain

//Lesson represent every box inside the app
type Course struct {
	CourseID    int64  `json:"courseID" gorm:"column:courseID;primary_key"`
	Name        string `json:"titlee" gorm:"column:Title"`
	Description string `json:"description" gorm:"column:Description;type:text"`
}
