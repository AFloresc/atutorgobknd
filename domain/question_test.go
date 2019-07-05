package domain

import (
	"context"
	"log"
	"testing"

	"github.com/atutor/utils"
)

func TestQuestion(t *testing.T) {
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

	mockQuestion1 := Question{
		Description: "¿Con que lenguaje de programación vamos a programar nuestras aplicaciones Android?",
		Answer1:     "Python",
		Answer2:     "Java",
		Answer3:     "XML",
		Answer4:     "Golang",
		GoodAnswer:  2,
	}

	mockQuestion2 := Question{
		Description: "¿Cómo se llama el IDE que vamos a utilizar para crear nuestros poryectos de Android?",
		Answer1:     "Android Studio",
		Answer2:     "IntelliJ",
		Answer3:     "VisualStudioCode",
		Answer4:     "Android IDE",
		GoodAnswer:  1,
	}

	mockQuestion3 := Question{
		Description: "¿Qué debemos descargar para empezar a programar en Android?",
		Answer1:     "Nada, lo tenemos todo ya en windows instalado por defecto",
		Answer2:     "Android SDK y Java JDK",
		Answer3:     "Android Station",
		Answer4:     "IntelliJ IDEA",
		GoodAnswer:  2,
	}

	mockQuestion4 := Question{
		Description: "¿Qué significa JDK?",
		Answer1:     "Java Development Kit",
		Answer2:     "Java Decompiler Kit",
		Answer3:     "JSON Device Kontrol",
		Answer4:     "Java Device Knowledge",
		GoodAnswer:  1,
	}

	mockQuestion5 := Question{
		Description: "¿Qué significa Android SDK?",
		Answer1:     "Java Development Kit",
		Answer2:     "Java Decompiler Kit",
		Answer3:     "JSON Device Kontrol",
		Answer4:     "Java Device Knowledge",
		GoodAnswer:  1,
	}

	mockQuestion6 := Question{
		Description: "¿Que es Android Studio?",
		Answer1:     "Lirerías Android",
		Answer2:     "Entorno de desarollo integrado de Android",
		Answer3:     "Librerias Java",
		Answer4:     "Entorno de desarrollo de aplicaciones Java",
		GoodAnswer:  2,
	}

	mockQuestion7 := Question{
		Description: "¿Que es una Blank Activity?",
		Answer1:     "Es una actividad de color blanco",
		Answer2:     "Es una actividad sin componentes añadidos",
		Answer3:     "Es una actividad de la clase Blank",
		Answer4:     "No es nada",
		GoodAnswer:  2,
	}

	mockQuestion8 := Question{
		Description: "¿Que componentes de Android Studio podemos instalar opcionalmente?",
		Answer1:     "Android SKD, Android Virtual Device y Performance",
		Answer2:     "Android libs",
		Answer3:     "No se puede instalar nada opcionalmente",
		Answer4:     "XML viewer tool",
		GoodAnswer:  1,
	}

	mockQuestion9 := Question{
		Description: "¿Que necesitamos instalar de Java?",
		Answer1:     "La máquina virtual de Java y el compilador de Java",
		Answer2:     "Solo la máquina virtual de Java",
		Answer3:     "Solo el compilador de Java",
		Answer4:     "No necesitamos nada de java",
		GoodAnswer:  1,
	}

	mockQuestion10 := Question{
		Description: "¿Que necesitamos instalar de Java?",
		Answer1:     "La máquina virtual de Java y el compilador de Java",
		Answer2:     "Solo la máquina virtual de Java",
		Answer3:     "Solo el compilador de Java",
		Answer4:     "No necesitamos nada de java",
		GoodAnswer:  1,
	}

	ctx := context.Background()

	t.Run("TestCreateQuestion", func(t *testing.T) {
		err := client.CreateQuestionary(ctx, &mockQuestionary)
		if err != nil {
			t.Fatal(err)
		}

		mockQuestion1.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion2.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion3.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion4.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion5.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion6.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion7.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion8.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion9.QuestionaryID = mockQuestionary.QuestionaryID
		mockQuestion10.QuestionaryID = mockQuestionary.QuestionaryID

		err = client.CreateQuestion(ctx, &mockQuestion1)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateQuestion(ctx, &mockQuestion2)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateQuestion(ctx, &mockQuestion3)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateQuestion(ctx, &mockQuestion4)
		if err != nil {
			t.Fatal(err)
		}
		err = client.CreateQuestion(ctx, &mockQuestion5)
		if err != nil {
			t.Fatal(err)
		}
		err = client.CreateQuestion(ctx, &mockQuestion6)
		if err != nil {
			t.Fatal(err)
		}
		err = client.CreateQuestion(ctx, &mockQuestion7)
		if err != nil {
			t.Fatal(err)
		}
		err = client.CreateQuestion(ctx, &mockQuestion8)
		if err != nil {
			t.Fatal(err)
		}
		err = client.CreateQuestion(ctx, &mockQuestion9)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateQuestion(ctx, &mockQuestion10)
		if err != nil {
			t.Fatal(err)
		}

	})

	// t.Run("TestGetQuestions", func(t *testing.T) {
	// 	questions, err := client.GetQuestionsByQuestionaryID(ctx, mockLesson.LessonID)

	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	assert := assert.New(t)

	// 	//assert.Equal(mockLesson.LessonID, lesson.LessonID)
	// })

	// t.Run("TestUpdateLesson", func(t *testing.T) {
	// 	mockLesson.Title = "Preparación del entorno de desarollo"
	// 	err := client.UpdateLesson(ctx, &mockLesson)

	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	lesson, err := client.GetLesson(ctx, mockLesson.LessonID)

	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	assert := assert.New(t)

	// 	assert.Equal(mockLesson.Title, lesson.Title)
	// })

	// t.Run("TestDeleteLesson", func(t *testing.T) {
	// 	err := client.DeleteLesson(ctx, mockLesson.LessonID)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	err = client.DeleteLesson(ctx, mockLesson2.LessonID)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	err = client.DeleteLesson(ctx, mockLesson3.LessonID)
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
