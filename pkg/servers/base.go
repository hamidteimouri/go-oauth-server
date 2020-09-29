package servers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hamidteimouri/go-oauth-server/pkg/controllers"
	"github.com/hamidteimouri/go-oauth-server/pkg/middlewares"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api", middlewares.SetMiddlewareJSON(controllers.Home))
}

func (server *Server) RunningReport(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
