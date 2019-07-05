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

	lesson, err := client.GetLesson(context.Background(), int64(1))
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

func TestGtLessonByCourseID(t *testing.T) {
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

	lessons, err := client.GetAllLessonsByCourseID(context.Background(), int64(1))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("LESSONS---> ", lessons)
	}

}

func TestLessonDev(t *testing.T) {
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

func TestGtLessonByLanguageDev(t *testing.T) {
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

	lessons, err := client.GetLessonsByLanguage(context.Background(), "es")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("LESSONS---> ", lessons)
	}

}

func TestCreateLessonsDev(t *testing.T) {

	mockLesson4 := Lesson{
		CourseID: int64(1),
		Title:    "Los controles RadioGroup y RadioButton",
		Text:     "Como utilizar RadioGroups y Radiobuttons en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561478361/radio_button_krw1we.png",
		Position: 3,
	}

	mockLesson5 := Lesson{
		CourseID: int64(1),
		Title:    "Control CheckBox",
		Text:     "Como utilizar CheckBoxes en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561478361/checkbox_ansqpb.png",
		Position: 4,
	}

	mockLesson6 := Lesson{
		CourseID: int64(1),
		Title:    "Control Spinner",
		Text:     "Como utilizar Spinners en Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561478361/spinner1_uuauhv.png",
		Position: 5,
	}

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

	ctx := context.Background()

	t.Run("TestCreateLesson", func(t *testing.T) {
		err := client.CreateLesson(ctx, &mockLesson4)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson5)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson6)
		if err != nil {
			t.Fatal(err)
		}

	})
}

func TestCreateLessonsPack2(t *testing.T) {

	mockLesson7 := Lesson{
		CourseID: int64(1),
		Title:    "Control ListView",
		Text:     "Como utilizar ListViews con una lista de strings",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561524940/ListView64x64_kpk20m.png",
		Position: 6,
	}

	mockLesson8 := Lesson{
		CourseID: int64(1),
		Title:    "Image Button",
		Text:     "Como utilizar una imagen como botón",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561525090/ImageButton64x64_aoixsw.png",
		Position: 7,
	}

	mockLesson9 := Lesson{
		CourseID: int64(1),
		Title:    "Noticicaciones sencillas",
		Text:     "Como crear notificaciones en un Toast",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561525377/Notification64x64_symw3t.png",
		Position: 8,
	}

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

	ctx := context.Background()

	t.Run("TestCreateLesson", func(t *testing.T) {
		err := client.CreateLesson(ctx, &mockLesson7)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson8)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson9)
		if err != nil {
			t.Fatal(err)
		}

	})
}

func TestCreateLessonsPack3(t *testing.T) {

	mockLesson10 := Lesson{
		CourseID: int64(1),
		Title:    "Objeto EditText",
		Text:     "Como utilizar los objetos Editext",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561554171/EditText64x64_gofnni.png",
		Position: 9,
	}

	mockLesson11 := Lesson{
		CourseID: int64(1),
		Title:    "Iniciar Activities",
		Text:     "Como imiciar una segunda Activity",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561554691/android_mcv7xx.png",
		Position: 10,
	}

	mockLesson12 := Lesson{
		CourseID: int64(1),
		Title:    "Iniciar Activities 2",
		Text:     "Como iniciar unas segunda Activity con parámetros",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561525377/Notification64x64_symw3t.png",
		Position: 11,
	}

	mockLesson13 := Lesson{
		CourseID: int64(1),
		Title:    "La clase SharedPreferences",
		Text:     "Como almacenar datos mediante la clase SharedPreferences",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561554974/SharedPrefs64x64_gcqgsp.png",
		Position: 12,
	}

	mockLesson14 := Lesson{
		CourseID: int64(1),
		Title:    "Usar un archivo externo",
		Text:     "Como almacenar datos en un archivo externo",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561555131/ExternalFile64x64_y8zom9.png",
		Position: 13,
	}

	mockLesson15 := Lesson{
		CourseID: int64(1),
		Title:    "Usar un archivo SD",
		Text:     "Como almacenar datos en un archivo de una memoria SD",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561555551/sd_y1zina.png",
		Position: 14,
	}

	mockLesson16 := Lesson{
		CourseID: int64(1),
		Title:    "Base de datos SQLite",
		Text:     "Como almacenar datos en una base de datos SQLite",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561555862/SQLITE_rfegkg.png",
		Position: 15,
	}

	mockLesson17 := Lesson{
		CourseID: int64(1),
		Title:    "Instalar App",
		Text:     "Como instalar una app en un dispositivo real",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561555862/SQLITE_rfegkg.png",
		Position: 16,
	}

	mockLesson18 := Lesson{
		CourseID: int64(1),
		Title:    "Layout: LineraLayout",
		Text:     "Como utilizar un Layout LinearLayout",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556379/layout_io35gf.png",
		Position: 17,
	}

	mockLesson19 := Lesson{
		CourseID: int64(1),
		Title:    "Layout: TableLayout",
		Text:     "Como utilizar un Layout TableLayout",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556379/layout_io35gf.png",
		Position: 18,
	}

	mockLesson20 := Lesson{
		CourseID: int64(1),
		Title:    "Layout: FrameLayout",
		Text:     "Como utilizar un Layout FrameLayout",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556379/layout_io35gf.png",
		Position: 19,
	}

	mockLesson21 := Lesson{
		CourseID: int64(1),
		Title:    "Layout: ScrollView y LinearLayout",
		Text:     "Como utilizar un Layout FrameLayout y ScrollView",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556379/layout_io35gf.png",
		Position: 20,
	}

	mockLesson22 := Lesson{
		CourseID: int64(1),
		Title:    "Layout: ScrollView y LinearLayout",
		Text:     "Como utilizar un Layout FrameLayout y ScrollView",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556379/layout_io35gf.png",
		Position: 21,
	}

	mockLesson23 := Lesson{
		CourseID: int64(1),
		Title:    "Icono de la app",
		Text:     "Como definir el icono de la app",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556794/Apps-Android-icon_jozqkt.png",
		Position: 22,
	}

	mockLesson24 := Lesson{
		CourseID: int64(1),
		Title:    "Reproducción de audio",
		Text:     "Como reproducir audio en una app",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 23,
	}

	mockLesson25 := Lesson{
		CourseID: int64(1),
		Title:    "Reproducción de audio: controles",
		Text:     "Comtroles de audio en una app (play, pause and stop)",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 24,
	}

	mockLesson26 := Lesson{
		CourseID: int64(1),
		Title:    "Reproducción de audio: de un archivo en memoria SD",
		Text:     "Como reproducir archivos de audio localizados en una memoria SD",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 25,
	}

	mockLesson27 := Lesson{
		CourseID: int64(1),
		Title:    "Reproducción de audio: de un archivo en internet",
		Text:     "Como reproducir archivos de audio localizados en internet",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 26,
	}

	mockLesson28 := Lesson{
		CourseID: int64(1),
		Title:    "Reproducción de audio: de un archivo en internet",
		Text:     "Como reproducir archivos de audio localizados en internet",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 27,
	}

	mockLesson29 := Lesson{
		CourseID: int64(1),
		Title:    "Reproducción de audio: utlizando el reproductor - Intent",
		Text:     "Como reproducir archivos de audio utilizando el reproductor de audio de Android mediante un Intent",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 28,
	}

	mockLesson30 := Lesson{
		CourseID: int64(1),
		Title:    "Grabar audio",
		Text:     "Como grabar audio mediante el grabador de audio del dispositivo Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 29,
	}

	mockLesson31 := Lesson{
		CourseID: int64(1),
		Title:    "La clase MediaRecorder",
		Text:     "Como grabar audio mediante la clase MediaRecorder",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561556986/64x64_yy0nj7.png",
		Position: 30,
	}

	mockLesson32 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: pixel",
		Text:     "Como dibujar pixeles en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 31,
	}

	mockLesson33 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: lineas y fondo",
		Text:     "Como dibujar lineas y el fondo en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 32,
	}

	mockLesson34 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: rectángulos",
		Text:     "Como dibujar rectángulos en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 33,
	}

	mockLesson35 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: círculos",
		Text:     "Como dibujar círculos en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 34,
	}

	mockLesson36 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: óvalos",
		Text:     "Como dibujar óvalos en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 35,
	}

	mockLesson37 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: texto",
		Text:     "Como añadir texto en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 36,
	}

	mockLesson38 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: texto con tipos de letra externos",
		Text:     "Como añadir texto con fuentes externas en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 37,
	}

	mockLesson39 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: texto sobre un camino",
		Text:     "Como añadir siguiendo un camino en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 38,
	}

	mockLesson40 := Lesson{
		CourseID: int64(1),
		Title:    "Gráficos: imágenes",
		Text:     "Como añadir imágenes en la app Android",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561557986/paint_kc3y4w.png",
		Position: 39,
	}

	mockLesson41 := Lesson{
		CourseID: int64(1),
		Title:    "Eventos: touch",
		Text:     "Como dibujar un círculo con un evento touch",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561559008/event_qsfzsc.png",
		Position: 40,
	}

	mockLesson42 := Lesson{
		CourseID: int64(1),
		Title:    "Eventos touch: círculo",
		Text:     "Como dibujar un círculo con un evento touch",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561559008/event_qsfzsc.png",
		Position: 41,
	}

	mockLesson43 := Lesson{
		CourseID: int64(1),
		Title:    "Eventos touch: juego buscaminas",
		Text:     "Implentar el juego buscaminas",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561559484/mine_z6se5s.png",
		Position: 42,
	}

	mockLesson44 := Lesson{
		CourseID: int64(1),
		Title:    "Archivo strings.xml",
		Text:     "Características y utilidades del archivo string.xml",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561559891/xml_i9hbz0.png",
		Position: 43,
	}

	mockLesson45 := Lesson{
		CourseID: int64(1),
		Title:    "Archivo strings.xml: Internacionalización",
		Text:     "Como utilizar diferentes idiomas con el archivo strings.xml",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561560347/languages_lxwfoj.png",
		Position: 44,
	}

	mockLesson46 := Lesson{
		CourseID: int64(1),
		Title:    "Archivo strings.xml: Localización",
		Text:     "Como utilizar diferentes idiomas por localización con el archivo strings.xml",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561560347/languages_lxwfoj.png",
		Position: 45,
	}

	mockLesson47 := Lesson{
		CourseID: int64(1),
		Title:    "ActionBar: Introducción",
		Text:     "Introducción al componente ActionBar",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561561114/actionbar_qdtcfl.png",
		Position: 46,
	}

	mockLesson48 := Lesson{
		CourseID: int64(1),
		Title:    "ActionBar: Botones de acción",
		Text:     "Como implementar botones de acción en el componente ActionBar",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561561114/actionbar_qdtcfl.png",
		Position: 47,
	}

	mockLesson49 := Lesson{
		CourseID: int64(1),
		Title:    "ActionBar: Visualización",
		Text:     "Como oculta y mostrar el componente ActionBar",
		Language: "es",
		Image:    "https://res.cloudinary.com/dnvu5jzwt/image/upload/v1561561114/actionbar_qdtcfl.png",
		Position: 48,
	}

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

	ctx := context.Background()

	t.Run("TestCreateLesson", func(t *testing.T) {
		err := client.CreateLesson(ctx, &mockLesson10)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson11)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson12)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson13)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson14)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson15)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson16)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson17)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson18)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson19)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson20)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson21)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson22)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson23)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson24)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson25)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson26)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson27)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson28)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson29)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson30)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson31)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson32)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson33)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson34)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson35)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson36)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson37)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson38)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson39)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson40)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson41)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson42)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson43)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson44)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson45)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson46)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson47)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson48)
		if err != nil {
			t.Fatal(err)
		}

		err = client.CreateLesson(ctx, &mockLesson49)
		if err != nil {
			t.Fatal(err)
		}

	})
}
