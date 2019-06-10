package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/atutor/ahttp"
	"github.com/atutor/domain"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

// Application struct
type Application struct {
	Config domain.Config
	Client *domain.Client
}

// InitializeRoutes : defines de endpoints routes and initializes them
func (ap *Application) InitializeRoutes(router *mux.Router) {
	baseRoute := "/api/1.0/atapi"
	fmt.Println(baseRoute + "/test")
	//Signup handlers
	router.HandleFunc(baseRoute+"/signup", ap.signup).Methods("POST")
	//Signin handers
	router.HandleFunc(baseRoute+"/login", ap.login).Methods("POST")
	//token protected routes
	router.HandleFunc(baseRoute+"/protected", TokenVerifyMiddleWare(ap.protectedEndpoint)).Methods("GET")
	router.HandleFunc(baseRoute+"/test", TokenVerifyMiddleWare(TestEndpoint)).Methods("GET")

	//Concept routes
	router.HandleFunc(baseRoute+"/concepts/{conceptid}", ap.GetConcept).Methods("GET")
	router.HandleFunc(baseRoute+"/concepts/languages/{language}", ap.GetAllConceptsByLanguage).Methods("GET")
	router.HandleFunc(baseRoute+"/concepts/", ap.CreateConcept).Methods("POST")
	router.HandleFunc(baseRoute+"/concepts/", ap.UpdateConcept).Methods("PUT")
	router.HandleFunc(baseRoute+"/concepts/{conceptid}", ap.DeleteConcept).Methods("DELETE")

	//Lesson routes
	router.HandleFunc(baseRoute+"/lessons/{lessonid}", ap.GetLesson).Methods("GET")
	router.HandleFunc(baseRoute+"/lessons/languages/{language}", ap.GetAllLessonsByLanguage).Methods("GET")
	router.HandleFunc(baseRoute+"/lessons/", ap.CreateLesson).Methods("POST")
	router.HandleFunc(baseRoute+"/lessons/", ap.UpdateLesson).Methods("PUT")
	router.HandleFunc(baseRoute+"/lessons/{lessonid}", ap.DeleteLesson).Methods("DELETE")

	//Course routes
	router.HandleFunc(baseRoute+"/courses/{courseid}", ap.GetCourse).Methods("GET")
	router.HandleFunc(baseRoute+"/courses/{courseid}/statistics", ap.GetCourseStatistics).Methods("GET")

	//User routes
	router.HandleFunc(baseRoute+"/users/{courseid}", ap.GetUser).Methods("GET")

	//Questionary routes
	//TODO
}

func (ap Application) signup(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	var error ahttp.Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		ahttp.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		ahttp.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)

	err = ap.Client.CreateUser(r.Context(), &user)

	if err != nil {
		error.Message = "Server error."
		ahttp.RespondWithError(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""

	ahttp.RespondWithJSON(w, http.StatusOK, user)

}

func (ap Application) login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	var jwt domain.JWT
	var error ahttp.Error
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		error.Message = "Invalid payload"
		ahttp.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Email == "" {
		error.Message = "Email is missing"
		ahttp.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		ahttp.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	//spew.Dump(user)
	//password := user.Password

	user, err := ap.Client.GetUserByEmail(r.Context(), user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "The user does not exist"
			ahttp.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
		if strings.Contains(err.Error(), "record not found") {
			error.Message = "The user does not exist"
			ahttp.RespondWithError(w, http.StatusBadRequest, error)
		} else {
			log.Fatal(err)
		}
		return
	}

	//spew.Dump(u)

	token, err := GenerateToken(user)
	if err != nil {
		fmt.Println("Error generating token->", err.Error())
		return
	}

	jwt.Token = token

	ahttp.RespondWithJSON(w, http.StatusOK, jwt)
}

func (ap Application) protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("protectedEndpoint invoked.")
	ahttp.RespondWithJSON(w, http.StatusOK, "succesfully called protectedEndpoints.")

}

// GetConcept : Gets a concept by conceptID
func (ap Application) GetConcept(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var errorObject ahttp.Error
	id, err := strconv.ParseInt(vars["conceptid"], 10, 0)
	if err != nil {
		errorObject.Message = "Invalid concept ID"
		ahttp.RespondWithError(w, http.StatusBadRequest, errorObject)
		return
	}

	concept, err := ap.Client.GetConcept(r.Context(), id)
	if err != nil {
		errorObject.Message = err.Error()
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	ahttp.RespondWithJSON(w, http.StatusOK, concept)
}

// GetAllConcepts :
func (ap Application) GetAllConceptsByLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var errorObject ahttp.Error

	lang := vars["language"]

	concept, err := ap.Client.GetConceptsByLanguage(r.Context(), lang)
	if err != nil {
		errorObject.Message = err.Error()
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	ahttp.RespondWithJSON(w, http.StatusOK, concept)
}

// CreateConcept : Creates a concept into persistence
func (ap Application) CreateConcept(w http.ResponseWriter, r *http.Request) {
	var concept domain.Concept
	var errorObject ahttp.Error
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&concept); err != nil {
		errorObject.Message = err.Error()
		ahttp.RespondWithError(w, http.StatusBadRequest, errorObject)
		return
	}
	ahttp.RespondWithJSON(w, http.StatusOK, concept)
}

// UpdateConcept :
func (ap Application) UpdateConcept(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// DeleteConcept :
func (ap Application) DeleteConcept(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// GetLesson : returns a lesson by lessonID
func (ap Application) GetLesson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var errorObject ahttp.Error
	id, err := strconv.ParseInt(vars["lessonid"], 10, 0)
	if err != nil {
		errorObject.Message = "Invalid lesson ID"
		ahttp.RespondWithError(w, http.StatusBadRequest, errorObject)
		return
	}

	lesson, err := ap.Client.GetLesson(r.Context(), id)
	if err != nil {
		errorObject.Message = err.Error()
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	if lesson == nil {
		errorObject.Message = "Lesson not found"
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	ahttp.RespondWithJSON(w, http.StatusOK, lesson)
}

// GetAllLessonsByLanguage : retrieves all lesson by language
func (ap Application) GetAllLessonsByLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var errorObject ahttp.Error

	lang := vars["language"]

	lessons, err := ap.Client.GetLessonsByLanguage(r.Context(), lang)
	if err != nil {
		errorObject.Message = err.Error()
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	if len(lessons) < 1 {
		errorObject.Message = "Language " + lang + " has no lessons."
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	ahttp.RespondWithJSON(w, http.StatusOK, lessons)
}

// CreateLesson :
func (ap Application) CreateLesson(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// UpdateLesson :
func (ap Application) UpdateLesson(w http.ResponseWriter, r *http.Request) {
	//TODO
}

// DeleteLesson :
func (ap Application) DeleteLesson(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	token := context.Get(req, "user")
	var user domain.User
	var errorObject ahttp.Error
	mapstructure.Decode(token.(jwt.MapClaims), &user)
	if user.Role != string(domain.RoleLogin) {
		errorObject.Message = "Unauthorized acces."
		ahttp.RespondWithError(w, http.StatusForbidden, errorObject)
	} else {
		ahttp.RespondWithJSON(w, http.StatusOK, user)
		//json.NewEncoder(w).Encode(user)
	}
}

// GetCourse : returns a whole course
func (ap Application) GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var errorObject ahttp.Error
	id, err := strconv.ParseInt(vars["courseid"], 10, 0)
	if err != nil {
		errorObject.Message = "Invalid course ID"
		ahttp.RespondWithError(w, http.StatusBadRequest, errorObject)
		return
	}

	course, err := ap.Client.GetCourse(r.Context(), id)
	if err != nil {
		errorObject.Message = err.Error()
		ahttp.RespondWithError(w, http.StatusInternalServerError, errorObject)
		return
	}
	ahttp.RespondWithJSON(w, http.StatusOK, course)
}

// GetCourseStatistics :
func (ap Application) GetCourseStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// GetCourseStatistics :
func (ap Application) GetUser(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject ahttp.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		// fmt.Println(bearerToken)

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte(os.Getenv("SECRET")), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				ahttp.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
			// spew.Dump(token)
			if token.Valid {
				log.Println("Token claims: ", token.Claims)
				context.Set(r, "user", token.Claims)
				next.ServeHTTP(w, r)
			} else {
				errorObject.Message = error.Error()
				ahttp.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invaid token."
			ahttp.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := os.Getenv("SECRET") // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

// Initialize : load data from json file
func (ap *Application) Initialize() {
	fmt.Println("Starting the application...")

	fmt.Println(`    _   _____        _`)
	fmt.Println(`   / \ |_   _|_   _ | |_  ___   _ __ `)
	fmt.Println(`  / _ \  | | | | | || __|/ _ \ | '__|`)
	fmt.Println(` / ___ \ | | | |_| || |_| (_) || |   `)
	fmt.Println(`/_/   \_\|_|  \__,_| \__|\___/ |_|by Alex Flores`)

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		fmt.Println("Error parsing por. ", err.Error())
	}
	ap.Client.Initialize(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_IP"), port, os.Getenv("DATABASE_NAME"))
}

func loadConfiguration(filename string) (domain.Config, error) {
	var config domain.Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

// GenerateToken : Generates a valid token
func GenerateToken(user domain.User) (string, error) {
	var err error
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "ATutorCourse",
		"role":  domain.RoleLogin,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
