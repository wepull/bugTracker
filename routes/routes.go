package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	"github.com/wepull/bugTracker/handlers"
	"github.com/wepull/bugTracker/models"
)

const (
	LISTEN_ADDR        = "127.0.0.1"
	HTTP_PORT          = "60001"
	HTTP_READ_TIMEOUT  = 180
	HTTP_WRITE_TIMEOUT = 180
	REQ_TIMEOUT_PERIOD = 180
	SWAGGERDOCS        = "/docs"
	SWAGGER_YML        = "/swagger.yml"
)

const (
	API_HEALTH = "/api/health"

	// User Endpoints
	USER_REGISTER = "/user/register"
	USER_LOGIN    = "/user/login"

	// Bug Endpoints
	CREATE_BUG    = "/bugs/create"
	GET_BUG_BY_ID = "/bugs/get/{id}"
	GET_BUGS      = "/bugs/get"
	UPDATE_BUG    = "/bugs/update/{id}"
	DELETE_BUG    = "/bugs/delete/{id}"

	// Project Endpoints
	CREATE_PROJECT    = "/projects/create"
	GET_PROJECT_BY_ID = "/projects/get/{id}"
	GET_PROJECTS      = "/projects/get"
	UPDATE_PROJECT    = "/projects/update/{id}"
	DELETE_PROJECT    = "/projects/delete/{id}"
)

// swagger:route POST /api/health Health HealthStatus
// Get health status of roostapi
// This doesn't requires api authentication
//
//	Consumes:
//	- application/json
//	Produces:
//	- application/json
//	Responses:
//	  200: ApiResp
func HandleServerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request %+v", r)
	//dummy API
	resp := models.ApiResp{
		ResponseCode:        200,
		ResponseDescription: "Bug Tracker Backend Server API is up",
	}

	defer func() {
		out, _ := json.Marshal(resp)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write(out)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}()
}

func GetRouter() (*mux.Router, error) {
	router := mux.NewRouter()

	router.HandleFunc(API_HEALTH, HandleServerHealth).Methods(http.MethodGet)
	router.HandleFunc(USER_REGISTER, handlers.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc(USER_LOGIN, handlers.LoginUser).Methods(http.MethodPost)

	router.HandleFunc(CREATE_BUG, handlers.CreateBug).Methods(http.MethodPost)
	router.HandleFunc(GET_BUG_BY_ID, handlers.GetBugByID).Methods(http.MethodGet)
	router.HandleFunc(GET_BUGS, handlers.GetBugs).Methods(http.MethodGet)
	router.HandleFunc(UPDATE_BUG, handlers.UpdateBug).Methods(http.MethodPut)
	router.HandleFunc(DELETE_BUG, handlers.DeleteBug).Methods(http.MethodDelete)

	router.HandleFunc(CREATE_PROJECT, handlers.CreateProject).Methods(http.MethodPost)
	router.HandleFunc(GET_PROJECT_BY_ID, handlers.GetProjectByID).Methods(http.MethodGet)
	router.HandleFunc(GET_PROJECTS, handlers.GetProjects).Methods(http.MethodGet)
	router.HandleFunc(UPDATE_PROJECT, handlers.UpdateProject).Methods(http.MethodPut)
	router.HandleFunc(DELETE_PROJECT, handlers.DeleteProject).Methods(http.MethodDelete)

	opts := middleware.SwaggerUIOpts{SpecURL: SWAGGER_YML}
	swaggerDoc := middleware.SwaggerUI(opts, nil)

	router.Handle(SWAGGERDOCS, swaggerDoc).Methods(http.MethodGet)
	router.Handle(SWAGGER_YML, http.FileServer(http.Dir("./openapi"))).Methods(http.MethodGet)

	return router, nil
}

func StartHTTPSvr() error {

	router, err := GetRouter()
	if err != nil {
		fmt.Printf("Error in setting up router: %v", err)
		return err
	}

	fmt.Println("Get Router successful")
	server := &http.Server{
		Addr:         LISTEN_ADDR + ":" + HTTP_PORT,
		Handler:      router,
		ReadTimeout:  time.Duration(HTTP_READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(HTTP_WRITE_TIMEOUT) * time.Second,
	}
	fmt.Printf("### HTTP Server listening on %v\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
		return err
	}
	return nil
}
