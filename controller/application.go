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

func (ap *Application) InitializeRoutes(router *mux.Router) {
	//Signup handlers
	router.HandleFunc("/signup", ap.signup).Methods("POST")
	//Signin handers
	router.HandleFunc("/login", ap.login).Methods("POST")
	//token protected routes
	router.HandleFunc("/protected", TokenVerifyMiddleWare(ap.protectedEndpoint)).Methods("GET")
	router.HandleFunc("/test", TokenVerifyMiddleWare(TestEndpoint)).Methods("GET")
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
