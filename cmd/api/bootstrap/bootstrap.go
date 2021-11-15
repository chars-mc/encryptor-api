package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chars-mc/encryptor-api/internal/api/router"
	"github.com/chars-mc/encryptor-api/internal/api/server"
	"github.com/chars-mc/encryptor-api/internal/config"
	"github.com/chars-mc/encryptor-api/internal/database"
	"github.com/joho/godotenv"

	dataApplication "github.com/chars-mc/encryptor-api/internal/encryption/application"
	dataInfra "github.com/chars-mc/encryptor-api/internal/encryption/infrastructure"
	userApplication "github.com/chars-mc/encryptor-api/internal/user/application"
	userInfra "github.com/chars-mc/encryptor-api/internal/user/infrastructure"
)

func Start() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dbcfg := config.NewDBConfig(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	db, err := database.NewMySQLClient(dbcfg)
	if err != nil {
		log.Fatalf("cannot connect to mysql server: %+v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("cannot make ping to mysql server: %+v", err)
	}

	userRepository := userInfra.NewUserMySQLRepository(db)
	userService := userApplication.NewUserService(userRepository)
	userHandler := userInfra.NewUserHandler(userService)

	dataRepository := dataInfra.NewDataMySQLRepository(db)
	dataService := dataApplication.NewDataService(dataRepository)
	dataHandler := dataInfra.NewDataHandler(dataService)

	// routes
	routes := &router.Routes{
		router.NewRoute("/hello", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		}),
		router.NewRoute("/user/login", http.MethodPost, userHandler.LoginHandler),
		router.NewRoute("/user/signup", http.MethodPost, userHandler.SignUpHandler),
		router.NewRoute("/encrypt", http.MethodPost, dataHandler.EncryptHandler),
	}
	router := router.NewRouter(routes)

	// server
	apiPort := os.Getenv("API_PORT")
	server := server.NewServer(apiPort, router)
	fmt.Printf("api server on port %s\n", apiPort)

	return server.Run()
}
