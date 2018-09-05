package main

import (
	"schoolsystem/auth-microservice/auth"
	"schoolsystem/auth-microservice/core"
	"schoolsystem/auth-microservice/db"

	authRepository "schoolsystem/auth-microservice/auth/repository"
	authUsecase "schoolsystem/auth-microservice/auth/usecase"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	// setup logger
	logger, err := core.NewLogger()
	if err != nil {
		logger.Panic(err.Error())
	}

	// setup config
	conf := core.NewConfig()

	// database connection
	dbConn, err := db.NewDatabaseConnection(conf.Database.URL)

	if err != nil {
		logger.Panic(err.Error())
	}

	defer dbConn.Close()

	// setup routes and middleware
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	/****************************************
	*	REPOSITORIES && USECASES	   		*
	*****************************************/
	authRepo := authRepository.NewMongoAuthRepository(dbConn)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	auth.NewAuthHttpHandler(r, authUsecase)

	handler := cors.Default().Handler(r)
	http.ListenAndServe(conf.BaseURL, handler)
}

/*
	custom errors (db related) are in the repository level
	custom errors (usecase related) are in the usecase layer
	validations are in the transport layer
	logs are in the transport level since errors are passed from REPO -> USECASE -> TRANSPORT
*/
