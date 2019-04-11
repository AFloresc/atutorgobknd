package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

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

//var books []Book

func main() {
	//app := Application{}

	//client := MyClient{}
	//app.Initialize("config.json")

	// ctx := context.Background()

	router := mux.NewRouter()

	//Signup handlers
	router.HandleFunc("/signup", signup).Methods("POST")
	//Signin handers
	router.HandleFunc("/login", login).Methods("POST")
	//token protected routes
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {

	ahttp.RespondWithJSON(w, http.StatusOK, "succesfully called signup.")

}

func login(w http.ResponseWriter, r *http.Request) {
	user := domain.User{}

	json.NewDecoder(r.Body).Decode(&user)

	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Token: ", token)

	//ahttp.RespondWithJSON(w, http.StatusOK, "succesfully called login.")

}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {

	ahttp.RespondWithJSON(w, http.StatusOK, "succesfully called protectedEndpoints.")

}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	log.Println("TokenVerifyMiddleWare invoked.")

	return nil
}

// Initialize : load data from json file
func (a *Application) Initialize(filename string) {
	fmt.Println("Starting the application...")
	config, _ := loadConfiguration(filename)
	fmt.Println(config.Database.IP)
	println("loadconfig USER: ", config.Database.User)
	println("loadconfig PASS: ", config.Database.Pass)
	println("loadconfig IP: ", config.Database.IP)
	println("loadconfig PORT: ", config.Database.Port)
	println("loadconfig NAME: ", config.Database.Name)
	a.initialize(config.Database.User, config.Database.Pass, config.Database.IP, config.Database.Port, config.Database.Name)
}

// Initialize : Initialize the database
func (a *Application) initialize(user string, password string, ip string, port int, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, ip, port, dbname)
	println("Connection: ", connectionString)
	var err error
	a.db2, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established.")
	}
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
