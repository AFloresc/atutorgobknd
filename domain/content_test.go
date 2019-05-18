package domain

import (
	"context"
	"log"
	"testing"

	"github.com/atutor/utils"
	"github.com/stretchr/testify/assert"
)

func TestContent(t *testing.T) {

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
	mockContent := Content{
		Title:    "\nDescarga de Java SDK\n",
		LessonID: int64(3),
		Content:  "El primer paso es instalar el compilador de Java y la máquina virtual (Java SE Development Kit (JDK)). Se puede descargar de:\n \n https://www.oracle.com/technetwork/es/java/javase/downloads/index.html \n\n",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000074/java_sdk_qqpztl.pngword",
	}

	mockContent2 := Content{
		Title:    "\nDescarga de Android Studio\n",
		LessonID: int64(3),
		Content:  "El segundo paso es descargar el entorno que permite el desarrollo de aplicaciones Android,  Android Studio (Android SDK) en el siguiente enlace:\n \n https://developer.android.com/studio \n\n",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000074/java_sdk_qqpztl.pngword",
	}

	mockContent3 := Content{
		Title:    "|nDescarga de Android Studio\n",
		LessonID: int64(3),
		Content:  "El segundo paso es descargar el entorno que permite el desarrollo de aplicaciones Android,  Android Studio (Android SDK) en el siguiente enlace:\n \n https://developer.android.com/studio \n \n Procedemos a la instalación de Android Studio. \n",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558080802/c1l1c3_smbnzs.png",
	}

	mockContent4 := Content{
		Title:    "Componentes de instalación",
		LessonID: int64(3),
		Content:  "Dejamos por defecto para que instale Android Studio, Android SDK, Android Virtual Device y Performance. \n",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558081406/c1l1c4_ubp50e.png",
	}

	mockContent5 := Content{
		Title:    "Instalamos",
		LessonID: int64(3),
		Content:  "Dejamos los valores por defecto y una vez instalado ejecutamos Android Studio. \n\nAl ejecutarlo es posible que nos aparezca una ventana de diálogo y empiece a descargarse actualizaciones.\n\nUna vez finalizado nos aparecerá la siguiente pantalla:\n",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558081406/c1l1c4_ubp50e.png",
	}

	ctx := context.Background()

	t.Run("TestGetContentNotExist", func(t *testing.T) {
		_, err := client.GetContent(ctx, mockContent.ContentID)

		assert := assert.New(t)

		assert.Nil(err)

	})

	t.Run("TestCreateContents", func(t *testing.T) {
		// Lesson 1 contents
		err := client.CreateContent(ctx, &mockContent)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent2)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent3)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent4)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent5)
		if err != nil {
			t.Fatal(err)
		}
	})

}
