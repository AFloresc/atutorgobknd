package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/atutor/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestApplicationHandlers(t *testing.T) {
	app := Application{
		Client: &domain.Client{},
	}

	app.InitializeForTest()
	err := app.Client.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	app.InitializeRoutes(router)

	httpclient := &http.Client{}

	//ctx := context.Background()

	ts := httptest.NewServer(router)

	t.Run("TestGetQuestionaryHandler", func(t *testing.T) {

		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		lessonid := strconv.FormatInt(n, 10)
		path := baseRoute + "/lessons/" + lessonid + "/questionary"
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)

		// body, err := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// if err != nil {
		// 	t.Fatal(err)
		// }

		//println("BODY --> ", string(body))

		// profileTopics := []profiles.ProfileTopic{}
		// err = app.Client.GetProfileTopics(ctx, mockProfile.ProfileID, &profileTopics)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// println("SIZE: ", len(profileTopics))

		// assert.Equal(1, len(profileTopics))
		// assert.Equal(mockProfile.ProfileID, profileTopics[0].ProfileID)
		// assert.Equal(mockTopic2.ID, profileTopics[0].TopicID)
	})

	t.Run("TestGetCourseHandler", func(t *testing.T) {

		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		courseid := strconv.FormatInt(n, 10)

		path := baseRoute + "/courses/" + courseid
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetConceptHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		conceptid := strconv.FormatInt(n, 10)

		path := baseRoute + "/concepts/" + conceptid
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetGetAllConceptsByLanguageHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		language := "es"

		path := baseRoute + "/concepts/languages/" + language
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetGetAllLessonsByLanguageHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		language := "es"

		path := baseRoute + "/lessons/languages/" + language
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		// body, err := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// if err != nil {
		// 	t.Fatal(err)
		// }

		// println("BODY --> ", string(body))

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	// "/lessons/{lessonid}/questionary"

	t.Run("TestGetGetAllLessonsByLanguageHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		lessonid := strconv.FormatInt(n, 10)

		path := baseRoute + "/lessons/" + lessonid + "/questionary"
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetQuestionaryByID", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		questionaryid := strconv.FormatInt(n, 10)

		path := baseRoute + "/questionaries/" + questionaryid
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})
}

func TestApplicationHandlersDev(t *testing.T) {
	app := Application{
		Client: &domain.Client{},
	}

	app.InitializeForTestDev()
	err := app.Client.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	app.InitializeRoutes(router)

	httpclient := &http.Client{}

	//ctx := context.Background()

	ts := httptest.NewServer(router)

	t.Run("TestGetQuestionaryHandler", func(t *testing.T) {

		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		lessonid := strconv.FormatInt(n, 10)
		path := baseRoute + "/lessons/" + lessonid + "/questionary"
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)

		// body, err := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// if err != nil {
		// 	t.Fatal(err)
		// }

		//println("BODY --> ", string(body))

		// profileTopics := []profiles.ProfileTopic{}
		// err = app.Client.GetProfileTopics(ctx, mockProfile.ProfileID, &profileTopics)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// println("SIZE: ", len(profileTopics))

		// assert.Equal(1, len(profileTopics))
		// assert.Equal(mockProfile.ProfileID, profileTopics[0].ProfileID)
		// assert.Equal(mockTopic2.ID, profileTopics[0].TopicID)
	})

	t.Run("TestGetCourseHandler", func(t *testing.T) {

		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		courseid := strconv.FormatInt(n, 10)

		path := baseRoute + "/courses/" + courseid
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetConceptHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		conceptid := strconv.FormatInt(n, 10)

		path := baseRoute + "/concepts/" + conceptid
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetGetAllConceptsByLanguageHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		language := "es"

		path := baseRoute + "/concepts/languages/" + language
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetGetAllLessonsByLanguageHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		language := "es"

		path := baseRoute + "/lessons/languages/" + language
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		// body, err := ioutil.ReadAll(res.Body)
		// res.Body.Close()
		// if err != nil {
		// 	t.Fatal(err)
		// }

		// println("BODY --> ", string(body))

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	// "/lessons/{lessonid}/questionary"

	t.Run("TestGetGetAllLessonsByLanguageHandler", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		lessonid := strconv.FormatInt(n, 10)

		path := baseRoute + "/lessons/" + lessonid + "/questionary"
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})

	t.Run("TestGetQuestionaryByID", func(t *testing.T) {
		baseRoute := "/api/1.0/atapi"
		n := int64(1)
		questionaryid := strconv.FormatInt(n, 10)

		path := baseRoute + "/questionaries/" + questionaryid
		req, err := http.NewRequest("GET", ts.URL+path, nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := httpclient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		assert := assert.New(t)
		assert.Equal(http.StatusOK, res.StatusCode)
	})
}
