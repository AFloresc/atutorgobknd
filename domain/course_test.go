package domain

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

func TestCourse(t *testing.T) {
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

	mockCourse := Course{
		Name:        "Introducción a Android 2019",
		Description: "Cusrso de introducción a la programación en dispositivos Android.",
		Target:      "Usuarios no expertos con conocimientos de POO",
	}

	mockCourse2 := Course{
		Name:        "Introducción a Goloang 2019",
		Description: "Curso de introduccion a la prgramación en Go",
		Target:      "Usuarios no expertos/expertos en otros lenguajes",
	}

	ctx := context.Background()

	t.Run("TestGetCourseNotExist", func(t *testing.T) {
		_, err := client.GetCourse(ctx, mockCourse.CourseID)

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateCourse", func(t *testing.T) {
		err := client.CreateCourse(ctx, &mockCourse)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateCourse(ctx, &mockCourse2)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("TestGetCourse", func(t *testing.T) {
		course, err := client.GetCourse(ctx, mockCourse.CourseID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockCourse.CourseID, course.CourseID)

	})

	t.Run("TestUpdateCourse", func(t *testing.T) {
		mockCourse.Name = "CursoAndroid"
		err := client.UpdateCourse(ctx, &mockCourse)

		if err != nil {
			t.Error(err)
		}
		course, err := client.GetCourse(ctx, mockCourse.CourseID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockCourse.Name, course.Name)
	})

	t.Run("TestGetGetConceptBySearch", func(t *testing.T) {
		concepts, err := client.GetConceptBySearch(ctx, "Activity")
		if err != nil {
			t.Error(err)
		}

		assert := assert.New(t)
		for _, concept := range concepts {
			assert.Equal(concept.Title, "Activity")
		}

	})

	t.Run("TestDeleteCourse", func(t *testing.T) {
		err := client.DeleteCourse(ctx, mockCourse.CourseID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteCourse(ctx, mockCourse2.CourseID)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("TestHardDeleteCourse", func(t *testing.T) {
		err := client.hardDeleteCourse(ctx, mockCourse.CourseID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.hardDeleteCourse(ctx, mockCourse2.CourseID)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestGetAndroidCourse(t *testing.T) {
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

	course, err := client.GetCourse(context.Background(), int64(1))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Course---> ", course)

}
