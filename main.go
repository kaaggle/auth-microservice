package main

import (
	"church-adoration/appointment"
	"church-adoration/auth"
	"church-adoration/middlewares"

	"church-adoration/core"
	"church-adoration/db"
	"log"

	authRepository "church-adoration/auth/repository"
	authUsecase "church-adoration/auth/usecase"

	appointmentRepository "church-adoration/appointment/repository"
	appointmentUsecase "church-adoration/appointment/usecase"

	"net/http"

	"github.com/casbin/casbin"
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
	conf, err := core.NewConfig()
	if err != nil {
		logger.Panic(err.Error())
	}
	log.Println(conf.String())

	// database connection
	dbConn, err := db.NewDatabaseConnection(conf.Database.URL)

	if err != nil {
		logger.Panic(err.Error())
	}
	log.Println("Successfully connected to mlab.")

	defer dbConn.Close()

	// setup casbin
	e, err := casbin.NewEnforcerSafe(conf.CasbinConfPath+"model.conf", conf.CasbinConfPath+"policy.csv")

	if err != nil {
		logger.Panic(err.Error())
	}

	// setup routes and middleware
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(middlewares.Authorizer(e))
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

	appointmentRepo := appointmentRepository.NewMongoAppointmentRepository(dbConn)
	appointmentUsecase := appointmentUsecase.NewAppointmentUsecase(appointmentRepo)
	appointment.NewAppointmentHttpHandler(r, appointmentUsecase)

	/* cours and server setup */
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "localhost:3000"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := c.Handler(r)
	log.Printf("Server running on: %s", conf.BaseURL)
	err = http.ListenAndServe(conf.BaseURL, handler)
	if err != nil {
		log.Printf("Error running server: %s", err.Error())
	}

}

/*
	custom errors (db related) are in the repository level
	custom errors (usecase related) are in the usecase layer
	validations are in the transport layer
	logs are in the transport level since errors are passed from REPO -> USECASE -> TRANSPORT
*/
