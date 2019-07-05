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

func TestConceptDataDevCreationPack1(t *testing.T) {
	t.Skip()
	dbConfig, err := utils.GetMySQLConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConfig.Passwd = "a1234"

	dbConfig.Addr = "35.205.235.6:3306"
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

	courseID := int64(1)

	mockConcept := Concept{
		Title:       "Activity",
		Description: "Podemos decir que todas las pantallas de una aplicación son una activity. Es decir, si una aplicación tiene cinco pantallas, tiene 5 “Actividades” o activities. Una activity tiene tiene separada la lógica (java) de la parte gráfica (XML). La parte lógica extiende de la clase Activity.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept2 := Concept{
		Title:       "Service",
		Description: "No tienen interfaz visual y se ejecutan en segundo plano, se encargan de realizar tareas que deben continuar ejecutandose cuando nuestra aplicación no está en primer plano. Todos los servicios extienden de la clase Service. Los servicios disponen de un mecanismo para ejecutar tareas pesadas sin bloquear la aplicación ya que se ejecutan en un hilo distinto.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept3 := Concept{
		Title:       "BroadcastReceiver",
		Description: "Simplemente reciben un mensaje y reaccionan ante él, extienden de la clase BroadcastReceiver, no tienen interfaz de usuario, pero pueden lanzar Actividades como respuesta a un evento o usar NotificationManager para informar al usuario.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept4 := Concept{
		Title:       "ContentProvider",
		Description: "Ponen un grupo de datos a disposición de distintas aplicaciones, extienden de la clase ContentProvider para implementar los métodos de la interfaz, pero para acceder a esta interfaz se ha de usar una clase llamada ContentResolver.",
		Language:    "es",
		CourseID:    courseID,
	}

	ctx := context.Background()

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
}

func TestConceptDataDevCreationPack2(t *testing.T) {
	dbConfig, err := utils.GetMySQLConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConfig.Passwd = "a1234"

	dbConfig.Addr = "35.205.235.6:3306"
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

	courseID := int64(1)

	mockConcept := Concept{
		Title:       "Android Manifest",
		Description: "Antes de que Android pueda iniciar un componente de la aplicación, el sistema debe conocer que ese componente existe leyendo el archivo AndroidManifest.xml. La app debe declarar todos sus componentes en este archivo. Este archivo está en la raíz sel proyecto de la app.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept2 := Concept{
		Title:       "RecyclerView",
		Description: "Cuando queremo mostrar en pantalla un gran número de elementos en una lista con scroll no todos los elementos son visibles, el usuario solo ve los elementos que caben en la pantalla. Cada uno de los elementos de la lista tiene la misma estructura.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept3 := Concept{
		Title:       "Drawable",
		Description: "Un drawable es un gráfico que puede ser dibujado en la pantalla. Pueden estar definidos en XML.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept4 := Concept{
		Title:       "Material Design",
		Description: "Es una filosofía visual de diseño creada por Google en 2014. Define guias de estilo, de diseño, movimiento y diversos aspectos relacionados con el diseño de la app.",
		Language:    "es",
		CourseID:    courseID,
	}

	ctx := context.Background()

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
}

func TestConceptDataDevCreationPack3(t *testing.T) {
	dbConfig, err := utils.GetMySQLConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConfig.Passwd = "a1234"

	dbConfig.Addr = "35.205.235.6:3306"
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

	courseID := int64(1)

	mockConcept := Concept{
		Title:       "Toast",
		Description: "La clase Toast nos permite mostrar un mensaje corto en la pantalla durante un tiempo determinado, no se puede interactuar con este elemento.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept2 := Concept{
		Title:       "SnackBar",
		Description: "Nos permite mostrar un mensaje corto y pedir la interacción del usuario. Desaparecen automaticamente despues de un tiempo o después de la interacción con el usuario.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept3 := Concept{
		Title:       "TextView",
		Description: "Es un widget que permite mostrar texto al usuario, no es editable por el usuario.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept4 := Concept{
		Title:       "EditText",
		Description: "Es un widget estandard de entrada de texto. Es editable por el usuario.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept5 := Concept{
		Title:       "Button",
		Description: "Es un widget de tipo botón. Permite iniciar una acción al ser presionado por el usuario.",
		Language:    "es",
		CourseID:    courseID,
	}

	mockConcept6 := Concept{
		Title:       "ImageButton",
		Description: "Muestra una imagen con características de botón. Al ser presionado por el usuario puede iniciar una acción.",
		Language:    "es",
		CourseID:    courseID,
	}

	ctx := context.Background()

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

		err = client.CreateConcept(ctx, &mockConcept5)
		if err != nil {
			t.Fatal(err)
		}
		err = client.CreateConcept(ctx, &mockConcept6)
		if err != nil {
			t.Fatal(err)
		}
	})
}
