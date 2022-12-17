package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"heroes/models"
	"heroes/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router    *mux.Router
	SubRouter *mux.Router
	DB        *sql.DB
}

func (a *App) Initialize() {

	user, password, dbhost, dbname := "postgres", "postgres", "host.docker.internal", "postgres"

	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, dbhost, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.SubRouter = a.Router.PathPrefix("/api").Subrouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	originsOK := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOK)(a.Router)))
}

// Init Routes
func (a *App) initializeRoutes() {
	a.SubRouter.HandleFunc("/request", a.createRequestHandler).Methods("POST")
	a.SubRouter.HandleFunc("/request/{email}", a.getRequestsByEmailHandler).Methods("GET")
	a.SubRouter.HandleFunc("/courses", a.getCoursesHandler).Methods("GET")
	a.SubRouter.HandleFunc("/users", a.getUsersHandler)
}

// Handler functions
func (a *App) createRequestHandler(w http.ResponseWriter, r *http.Request) {
	var tutorialRequest models.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tutorialRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid tutorial request payload")
		return
	}
	defer r.Body.Close()
	tutorialRequest, err := services.CreateRequest(tutorialRequest, a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, tutorialRequest)
}

func (a *App) getRequestsByEmailHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	requests, err := services.GetRequestsByEmail(email, a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, requests)
}

func (a *App) getCoursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := services.GetCourses(a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, courses)
}

func (a *App) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	courseId, err := strconv.ParseInt(r.URL.Query().Get("courseId"), 10, 32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	preferredTime, err := strconv.ParseInt(r.URL.Query().Get("preferredTime"), 10, 32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	users, err := services.GetUsers(courseId, preferredTime, a.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, users)
}

// Helper functions
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
