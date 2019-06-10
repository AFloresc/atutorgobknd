package domain

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

func TestLesson(t *testing.T) {
	dbConfig, err := utils.GetMySQLConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConfig.Passwd = "Bautista21"

	dbConfig.Addr = "localhost:3306"
	if err != nil {
		return
	}

	dbConfig.DBName = "atutor_dev"

	client, err := NewClient(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	err = client.AutoMigrate()
	if err != nil {
		t.Fatal(err)
	}

	mockLesson := Lesson{
		CourseID: int64(1),
		Title:    "Instalación de las herramientas necesarias",
		Text:     "Descarga de las herramientas necesarias para seguir el curso.",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000026/c1l1_e9ojcw.png",
		Position: 0,
	}

	mockLesson2 := Lesson{
		CourseID: int64(1),
		Title:    "El primer proyecto Android",
		Text:     "Como crear nuestro primer proyecto Android.",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000089/project_snnkng.png",
		Position: 1,
	}

	mockLesson3 := Lesson{
		CourseID: int64(1),
		Title:    "Capturar el click de un botón",
		Text:     "Como utilizar buttons en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000028/c1l2_dk87zd.png",
		Position: 2,
	}

	ctx := context.Background()

	t.Run("TestGetLessonNotExist", func(t *testing.T) {
		_, err := client.GetLesson(ctx, mockLesson.LessonID)

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateLesson", func(t *testing.T) {
		err := client.CreateLesson(ctx, &mockLesson)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson2)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson3)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("TestGetLesson", func(t *testing.T) {
		lesson, err := client.GetLesson(ctx, mockLesson.LessonID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockLesson.LessonID, lesson.LessonID)
	})

	t.Run("TestUpdateLesson", func(t *testing.T) {
		mockLesson.Title = "Preparación del entorno de desarollo"
		err := client.UpdateLesson(ctx, &mockLesson)

		if err != nil {
			t.Error(err)
		}
		lesson, err := client.GetLesson(ctx, mockLesson.LessonID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockLesson.Title, lesson.Title)
	})

	t.Run("TestDeleteLesson", func(t *testing.T) {
		err := client.DeleteLesson(ctx, mockLesson.LessonID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteLesson(ctx, mockLesson2.LessonID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteLesson(ctx, mockLesson3.LessonID)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestHardDeleteCourse", func(t *testing.T) {
		err := client.hardDeleteLesson(ctx, mockLesson.LessonID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.hardDeleteLesson(ctx, mockLesson.LessonID)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestGtLesson1(t *testing.T) {
	dbConfig, err := utils.GetMySQLConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConfig.Passwd = "Bautista21"

	dbConfig.Addr = "localhost:3306"
	if err != nil {
		return
	}

	dbConfig.DBName = "atutor_dev"

	client, err := NewClient(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	err = client.AutoMigrate()
	if err != nil {
		t.Fatal(err)
	}

	lesson, err := client.GetLesson(context.Background(), int64(3))
	if err != nil {
		fmt.Println("Error retrieving lesson 3")
	} else {
		fmt.Println("LESSON---> ", lesson)
	}

}

func TestGtLessonByLanguage(t *testing.T) {
	dbConfig, err := utils.GetMySQLConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConfig.Passwd = "Bautista21"

	dbConfig.Addr = "localhost:3306"
	if err != nil {
		return
	}

	dbConfig.DBName = "atutor_dev"

	client, err := NewClient(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	err = client.AutoMigrate()
	if err != nil {
		t.Fatal(err)
	}

	lessons, err := client.GetLessonsByLanguage(context.Background(), "jp")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("LESSONS---> ", lessons)
	}

}
