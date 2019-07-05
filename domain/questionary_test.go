package domain

import (
	"context"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

func TestQuestionary(t *testing.T) {
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

	mockQuestionary := Questionary{
		Description: "Questionary test1",
		LessonID:    int64(1),
	}

	ctx := context.Background()

	t.Run("TestGetQuestionaryNotExist", func(t *testing.T) {
		_, err := client.GetLesson(ctx, mockQuestionary.QuestionaryID)

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateQuestionary", func(t *testing.T) {
		err := client.CreateQuestionary(ctx, &mockQuestionary)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("TestGetQuestionary", func(t *testing.T) {
		questionary, err := client.GetQuestionary(ctx, mockQuestionary.QuestionaryID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockQuestionary.QuestionaryID, questionary.QuestionaryID)
	})

	t.Run("TestUpdateQuestionary", func(t *testing.T) {
		mockQuestionary.Description = "Questionary test1-MODIFIED"
		err := client.UpdateQuestionary(ctx, &mockQuestionary)

		if err != nil {
			t.Error(err)
		}
		questionary, err := client.GetQuestionary(ctx, mockQuestionary.QuestionaryID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockQuestionary.Description, questionary.Description)
	})

	t.Run("TestDeleteQuestionary", func(t *testing.T) {
		err := client.DeleteQuestionary(ctx, mockQuestionary.QuestionaryID)
		if err != nil {
			t.Fatal(err)
		}

	})

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
