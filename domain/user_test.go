package domain

import (
	"context"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

//TODO
func TestUser(t *testing.T) {
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

	mockUser := User{
		Name:     "Rodolfo Millán",
		Password: "1234",
		Email:    "rodolfo.millan@gmail.com",
	}

	mockMark := Mark{
		QuestionaryID: int64(1),
		Val:           7,
	}

	mockMark2 := Mark{
		QuestionaryID: int64(2),
		Val:           4,
	}

	ctx := context.Background()

	t.Run("TestGetUsernNotExist", func(t *testing.T) {
		_, err := client.GetLesson(ctx, mockUser.UserID)

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateUser", func(t *testing.T) {
		err := client.CreateUser(ctx, &mockUser)
		if err != nil {
			t.Fatal(err)
		}

		mockMark.UserID = mockUser.UserID
		err = client.CreateMark(ctx, &mockMark)
		if err != nil {
			t.Fatal(err)
		}

		mockMark2.UserID = mockUser.UserID
		err = client.CreateMark(ctx, &mockMark2)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("TestGetUser", func(t *testing.T) {
		user, err := client.GetUserByID(ctx, mockUser.UserID)
		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockUser.UserID, user.UserID)
	})

	// t.Run("TestUpdateUser", func(t *testing.T) {
	// 	mockUser.Email = "rodolfo.millan@hotmail.com"
	// 	err := client.UpdateUser(ctx, &mockUser)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	user, err := client.GetUserByPassword(ctx, mockUser.UserID)

	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	assert := assert.New(t)

	// 	assert.Equal(mockUser.Email, user.Email)
	// })

	// t.Run("TestDeleteUser", func(t *testing.T) {
	// 	err := client.DeleteUser(ctx, mockUser.UserID)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// })

	// t.Run("TestHardDeleteCourse", func(t *testing.T) {
	// 	err := client.hardDeleteLesson(ctx, mockLesson.LessonID)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	err = client.hardDeleteLesson(ctx, mockLesson.LessonID)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })
}

func TestGetUser(t *testing.T) {
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

	ctx := context.Background()

	t.Run("TestGetUser", func(t *testing.T) {
		user, err := client.GetUserByID(ctx, int64(1))
		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal("Alex Flores", user.Name)
	})
}
