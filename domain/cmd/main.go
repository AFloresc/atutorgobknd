package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/atutor/domain"
	"github.com/atutor/tapi/rbac"

	"github.com/atutor/ahttp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Application struct
type Application struct {
	db     *sql.DB
	db2    *gorm.DB
	Config Config
	Client *domain.Client
}

// Config struct
type Config struct {
	Database struct {
		User string `json:"user"`
		Pass string `json:"pass"`
		IP   string `json:"ip"`
		Port int    `json:"port"`
		Name string `json:"name"`
	} `json:"database"`
}

type JWT struct {
	Token string `json:"token"`
}

func init() {
	gotenv.Load()
}

func main() {
	app := Application{
		Client: &domain.Client{},
	}

	//app.Client = domain.Client{}

	app.Initialize(os.Getenv("MYSQL_CONNECTION"))
	err := app.Client.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
	//ctx := context.Background()

	router := mux.NewRouter()
	app.InitializeRoutes(router)
	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func (ap *Application) InitializeRoutes(router *mux.Router) {
	//Signup handlers
	router.HandleFunc("/signup", ap.signup).Methods("POST")
	//Signin handers
	router.HandleFunc("/login", ap.login).Methods("POST")
	//token protected routes
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")
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
	var jwt JWT
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
	// hashedPasword := user.Password
	// err = bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
	// if err != nil {
	// 	error.Message = "Invalid Password"
	// 	ahttp.RespondWithError(w, http.StatusBadRequest, error)
	// 	return
	// }

	token, err := GenerateToken(user)
	if err != nil {
		fmt.Println("Error generating token->", err.Error())
		return
	}

	jwt.Token = token

	ahttp.RespondWithJSON(w, http.StatusOK, jwt)
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("protectedEndpoint invoked.")
	ahttp.RespondWithJSON(w, http.StatusOK, "succesfully called protectedEndpoints.")

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

				return []byte("If I kill you, I am bound for hell. It is a price I shall gladly pay."), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				ahttp.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
			// spew.Dump(token)
			if token.Valid {
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

// Initialize : load data from json file
func (ap *Application) Initialize(connection string) {
	fmt.Println("Starting the application...")
	//config, _ := loadConfiguration(filename)
	// fmt.Println(config.Database.IP)
	// println("loadconfig USER: ", config.Database.User)
	// println("loadconfig PASS: ", config.Database.Pass)
	// println("loadconfig IP: ", config.Database.IP)
	// println("loadconfig PORT: ", config.Database.Port)
	// println("loadconfig NAME: ", config.Database.Name)
	//a.Client.Initialize(config.Database.User, config.Database.Pass, config.Database.IP, config.Database.Port, config.Database.Name)
	ap.Client.Initialize(connection)
}

func loadConfiguration(filename string) (Config, error) {
	var config Config
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
	secret := "If I kill you, I am bound for hell. It is a price I shall gladly pay."

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "ATutorCourse",
		"role":  rbac.RoleLogin,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
