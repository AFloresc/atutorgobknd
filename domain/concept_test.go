package domain

import (
	"context"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

func TestConcept(t *testing.T) {
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

	mockConcept := Concept{
		Title:       "Activity",
		Description: "Podemos decir que todas las pantallas de una aplicación son una activity. Es decir, si una aplicación tiene cinco pantallas, tiene 5 “Actividades” o activities. Una activity tiene tiene separada la lógica (java) de la parte gráfica (XML). La parte lógica extiende de la clase Activity.",
		Language:    "es",
	}

	mockConcept2 := Concept{
		Title:       "Service",
		Description: "No tienen interfaz visual y se ejecutan en segundo plano, se encargan de realizar tareas que deben continuar ejecutandose cuando nuestra aplicación no está en primer plano. Todos los servicios extienden de la clase Service. Los servicios disponen de un mecanismo para ejecutar tareas pesadas sin bloquear la aplicación ya que se ejecutan en un hilo distinto.",
		Language:    "es",
	}

	mockConcept3 := Concept{
		Title:       "BroadcastReceiver",
		Description: "Simplemente reciben un mensaje y reaccionan ante él, extienden de la clase BroadcastReceiver, no tienen interfaz de usuario, pero pueden lanzar Actividades como respuesta a un evento o usar NotificationManager para informar al usuario.",
		Language:    "es",
	}

	mockConcept4 := Concept{
		Title:       "ContentProvider",
		Description: "Ponen un grupo de datos a disposición de distintas aplicaciones, extienden de la clase ContentProvider para implementar los métodos de la interfaz, pero para acceder a esta interfaz se ha de usar una clase llamada ContentResolver.",
		Language:    "es",
	}

	ctx := context.Background()

	t.Run("TestGetConceptNotExist", func(t *testing.T) {
		_, err := client.GetConcept(ctx, mockConcept.ConceptID)

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateConcept", func(t *testing.T) {
		err := client.CreateConcept(ctx, &mockConcept)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateConcept(ctx, &mockConcept2)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateConcept(ctx, &mockConcept3)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateConcept(ctx, &mockConcept4)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestGetConcept", func(t *testing.T) {
		concept, err := client.GetConcept(ctx, mockConcept.ConceptID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockConcept.ConceptID, concept.ConceptID)
	})

	t.Run("TestUpdateConcept", func(t *testing.T) {
		mockConcept.Language = "en"
		err := client.UpdateConcept(ctx, &mockConcept)

		if err != nil {
			t.Error(err)
		}
		concept, err := client.GetConcept(ctx, mockConcept.ConceptID)

		if err != nil {
			t.Error(err)
		}
		assert := assert.New(t)

		assert.Equal(mockConcept.Language, concept.Language)
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

	t.Run("TestDeleteConcept", func(t *testing.T) {
		err := client.DeleteConcept(ctx, mockConcept.ConceptID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteConcept(ctx, mockConcept2.ConceptID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteConcept(ctx, mockConcept3.ConceptID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.DeleteConcept(ctx, mockConcept3.ConceptID)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestHardDeleteConcept", func(t *testing.T) {
		err := client.hardDeleteConcept(ctx, mockConcept.ConceptID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.hardDeleteConcept(ctx, mockConcept2.ConceptID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.hardDeleteConcept(ctx, mockConcept3.ConceptID)
		if err != nil {
			t.Fatal(err)
		}

		err = client.hardDeleteConcept(ctx, mockConcept4.ConceptID)
		if err != nil {
			t.Fatal(err)
		}
	})
}
