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

/*
dbConfig.Passwd = "a1234"

	dbConfig.Addr = "35.205.235.6:3306"
	if err != nil {
		return
	}

*/
func TestContentDev(t *testing.T) {

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
	mockContent := Content{
		Title:    "Descarga de Java SDK",
		LessonID: int64(3),
		Content:  "El primer paso es instalar el compilador de Java y la máquina virtual (Java SE Development Kit (JDK)). Se puede descargar de:\n https://www.oracle.com/technetwork/es/java/javase/downloads/index.html",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558000074/java_sdk_qqpztl.png",
	}

	mockContent2 := Content{
		Title:    "Descarga de Android Studio",
		LessonID: int64(3),
		Content:  "El segundo paso es descargar el entorno que permite el desarrollo de aplicaciones Android,  Android Studio (Android SDK) en el siguiente enlace:\n https://developer.android.com/studio",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561746840/android-studio-logo_rhxpfw.png",
	}

	mockContent3 := Content{
		Title:    "Instalación de Android Studio",
		LessonID: int64(3),
		Content:  "Procedemos a la instalación de Android Studio.",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558080802/c1l1c3_smbnzs.png",
	}

	mockContent4 := Content{
		Title:    "Componentes de instalación",
		LessonID: int64(3),
		Content:  "Dejamos los valores por defecto para que instale Android Studio, Android SDK, Android Virtual Device y Performance.",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1558081406/c1l1c4_ubp50e.png",
	}

	mockContent5 := Content{
		Title:    "Instalamos",
		LessonID: int64(3),
		Content:  "Dejamos los valores por defecto y una vez instalado ejecutamos Android Studio. Al ejecutarlo es posible que nos aparezca una ventana de diálogo y empiece a descargarse actualizaciones.\n\nUna vez finalizado nos aparecerá la siguiente pantalla:",
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

func TestContentDevPackLesson2(t *testing.T) {

	lsnID := int64(2)

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
	mockContent := Content{
		Title:    "Pasos para crear el primer proyecto Android",
		LessonID: lsnID,
		Content:  "Una vez iniciado Android Studio nos aparece la ventana de diálogo principal.",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561563628/L2001PNG_qdz9wr.png",
		Position: 0,
	}

	mockContent2 := Content{
		Title:    `Elegimos la opción "Start a New Android Studio project"`,
		LessonID: lsnID,
		Content:  "Aparecerán una serie de ventanas para la configuración del proyecto, el primer diálogo es el Nombre de la aplicación, la url de nuestra empresa (su dominio) que será el nombre del paquete que qsigna java para los archivos fuente y la ubicación en el disco de nuestro proyecto:",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561564039/L2002_fxuurv.png",
		Position: 1,
	}

	mockContent3 := Content{
		Title:    "En la segunda ventana de diálogo seleccionaremos la versión de Android mínima donde se ejecutará la apliación que desarrollamos.",
		LessonID: lsnID,
		Content:  "Seleccionaremos la version API 15: Android 4.0.3 (IceCreamSandwitch)",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561564230/L2003_cwf0rp.png",
		Position: 2,
	}

	mockContent4 := Content{
		Title:    "En la siguiente ventana nos pregunta que tipo de actividad nos va a añadir, podemos añadir actividades en blanco o con componentes ya creados.",
		LessonID: lsnID,
		Content:  `Seleccionaremos "Blank Activity" para que nos cree una actividad sin ningún componente. Los podemos añadir posteriormente si los necesitamos.`,
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561564610/L2004_xcmod7.png",
		Position: 3,
	}

	mockContent5 := Content{
		Title:    "Finalmente en la última ventana de diálogo indicamos el nombre de la ventana principa de la aplicación (Activity Name) y otros datos que veremos a lo largo del curso (dejamos los nombres que vienen por defecto).",
		LessonID: lsnID,
		Content:  "Le damos a Finish y nuestro ptoyecto inicial quedará creado.",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561564809/L2005_hmedsx.png",
		Position: 4,
	}

	mockContent6 := Content{
		Title:    "En poco rato se abre la ventana principal de Android Studio con el proyecto creado.",
		LessonID: lsnID,
		Content:  "La ventana principal tiene diferentes secciones que iremos conociendo a lo largo del curso.",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561564972/L2006_kdysdm.png",
		Position: 5,
	}

	mockContent7 := Content{
		Title:    "Android Studio nos genera todos los directorios y archivos básicos para iniciar nuestro proyecto.",
		LessonID: lsnID,
		Content:  "Podemos ver las carpetas creadas en lado izquiero del entorno de desarrollo.",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561565192/L2007_bzfnph.png",
		Position: 6,
	}

	mockContent8 := Content{
		Title:    "Por ahora no analizaremos el significado y objetivo de cada una de estas secciones y arrhivos generados, a medida que avance el curso se irán detallando.",
		LessonID: lsnID,
		Content:  "La interfaz visual de nuestro programa para Android se almacena en un archivo XML en la carpeta res, subcarpeta layout y el archivo se llama activity_main.xml. Este nombre sale del nombre que dimo a la Actividad principal pero con orden invertido (MainActivity). En esta carpeta está creada nuestra primera pantalla. Al seleccionar este archivo Android Studio nos permite visualizar el contenido en modo Diseño o en modo Texto (vista de Diseño o de código).",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561565192/L2007_bzfnph.png",
		Position: 7,
	}

	mockContent9 := Content{
		Title:    "Al hacer click en la pestaña Design vamos a la vista de diseño.",
		LessonID: lsnID,
		Content:  "La vista de diseño tiene el siguiente aspecto:",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561565610/L2008_y71j3w.png",
		Position: 8,
	}

	mockContent10 := Content{
		Title:    "Al hacer click en la pestaña Text vamos a la vista de texto.",
		LessonID: lsnID,
		Content:  "La vista de código tiene el siguiente aspecto:",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561565779/L2009_gomwn5.png",
		Position: 9,
	}

	mockContent11 := Content{
		Title:    "Si nos fijamos en el código Android Studio por defecto nos inserta un control de tipo RelativeLayout que permite añadir controles visuales alineados a los bordes y a otros controles que haya en la ventana. Este archivo se puede modificar completamente para que se adapte a la aplicación que queremos desarrollar.",
		LessonID: lsnID,
		Content:  `Antes de probar la aplicación en el emulador de un dispositivo Android procederemos a hacer algún pequeño cambio en la interfaz que aparece en el dispositivo Android: borraremos el label "Hello World" (simplemente seleccionandolo con el puntero del ratón y presionado la tecla delete) y de la "Palette" arrastraremos un objeto de tipo "Button" al centro del dispositivo móvil y en la ventana "Properties" estando seleccionado el "Button" cambiamos la propiedad "text" por "Hola Mundo"`,
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561566259/L2010_qcctvj.png",
		Position: 10,
	}

	mockContent12 := Content{
		Title:    "Para ejecutar la aplicación sele presionamos el triángulo verde o seleccionamos del menu de opciones",
		LessonID: lsnID,
		Content:  `Para ejecutar la aplicación"`,
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561566259/L2010_qcctvj.png",
		Position: 11,
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

		err = client.CreateContent(ctx, &mockContent6)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent7)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent8)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent9)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent10)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent11)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateContent(ctx, &mockContent12)
		if err != nil {
			t.Fatal(err)
		}

	})
}
