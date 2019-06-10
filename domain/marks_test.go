package domain

import (
	"context"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

func TestMark(t *testing.T) {
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

	mockMark := Mark{
		UserID:        int64(1),
		QuestionaryID: int64(2),
		Value:         8,
	}

	ctx := context.Background()

	t.Run("TestGetMarkNotExist", func(t *testing.T) {
		_, err := client.GetMark(ctx, int64(0), int64(1))

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateMark", func(t *testing.T) {
		err := client.CreateMark(ctx, &mockMark)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("TestGetMark", func(t *testing.T) {
		// Get questionaryID from lesson

		mark, err := client.GetMark(ctx, mockMark.UserID, mockMark.QuestionaryID)
		if err != nil {
			t.Fatal(err)
		}
		assert := assert.New(t)

		assert.Equal(mockMark.UserID, mark.UserID)
		assert.Equal(mockMark.QuestionaryID, mark.QuestionaryID)
		assert.Equal(mockMark.Value, mark.Value)
	})

	t.Run("TestGetMarkByLesson", func(t *testing.T) {
		// Get questionaryID from lesson

		lesson, err := client.GetQuestionary(ctx, mockMark.QuestionaryID)
		if err != nil {
			t.Fatal(err)
		}

		mark, err := client.GetMarkByLesson(ctx, mockMark.UserID, lesson.LessonID)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)

		assert.Equal(mockMark.UserID, mark.UserID)
		assert.Equal(mockMark.QuestionaryID, mark.QuestionaryID)
		assert.Equal(mockMark.Value, mark.Value)
	})

	t.Run("TestUpdateMark", func(t *testing.T) {
		mockMark.Value = 9
		err := client.UpdateMark(ctx, &mockMark)
		if err != nil {
			t.Fatal(err)
		}

		mark, err := client.GetMark(ctx, mockMark.UserID, mockMark.QuestionaryID)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)

		assert.Equal(mockMark.Value, mark.Value)
	})

	t.Run("TestDeleteMark", func(t *testing.T) {
		err := client.DeleteMark(ctx, mockMark.UserID, mockMark.QuestionaryID)
		if err != nil {
			t.Fatal(err)
		}
	})
}
