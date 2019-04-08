package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Book :
type Book struct {
	ID     int    `json:"id" gorm:"column:id;primary_key"`
	Title  string `json:"title" gorm:"column:title;type:text"`
	Author string `json:"author" gorm:"column:author;type:"`
	Year   string `json:"yeah"`
}

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

var books []Book

func main() {
	app := Application{}

	//client := MyClient{}
	app.Initialize("config.json")

	// ctx := context.Background()

	router := mux.NewRouter()

	//Signup handlers
	router.HandleFunc("/signup", signup).Methods("POST")
	//Signin handers
	router.HandleFunc("/login", login).Methods("POST")
	//token protected routes
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	// router.HandleFunc("/books", getBooks).Methods("GET")
	// router.HandleFunc("/books/{id}", getBook).Methods("GET")
	// router.HandleFunc("/books", addBook).Methods("POST")
	// router.HandleFunc("/books", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	log.Println("signup invoked.")

}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("login invoked.")

}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("protectedEndpoints invoked.")

}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	log.Println("TokenVerifyMiddleWare invoked.")
	return nil
}

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Get all books is called.")
// }
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Get book is called.")
// }

// func addBook(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Add book is called.")
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Update book is called.")
// }

// func removeBook(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Remove book is called.")
// }

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

// func timer(d time.Duration) <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		time.Sleep(d)
// 		c <- 1
// 		println("Timer: ", c)
// 	}()
// 	return c
// }

// func main() {
// 	for i := 0; i < 24; i++ {
// 		c := timer(1 * time.Second)
// 		<-c
// 	}
// }

// func main() {
// 	// create new channel of type int
// 	ch := make(chan int)

// 	// start new anonymous goroutine
// 	go func() {
// 		// send 42 to channel
// 		ch <- 42
// 	}()
// 	// read from channel
// 	<-ch
// }

// var yeees = true
// for {
// fmt.Println("This will happen first")

// go func() {
// 	fmt.Println("This will happen at some unknown time")
// }()

// fmt.Println("This will either happen second or third")

// fmt.Scanln()
// fmt.Println("done")
// if !yeees {
// 	break
// }

// }

//
// 	var c chan string = make(chan string)

// 	go pinger(c)
// 	go ponger(c)
// 	go printer(c)

// 	var input string
// 	fmt.Scanln(&input)
// }
// func pinger(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "ping"
// 	}
// }
// func ponger(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- " pong"
// 	}
// }
// func printer(c chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

/*
func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}
func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}
func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
}

*/
