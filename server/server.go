package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	controller "github.com/ashish036/GO-no-code-api/controller"
)

func Serve(ctx context.Context, port string) (err error) {
	r := mux.NewRouter()
	r = initialiseRouter(r)

	srv := createServerObj(port, r)

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Fatalf("listen failed: %v", err)
		}
	}()

	log.Printf("Server started on port : %s", port)

	// Graceful shutdown
	<-ctx.Done()
	log.Println("Server stopped")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctxShutDown); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Printf("Server exited properly")
	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func initialiseRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/", controller.HandleRequest)

	alpha_neumeric_regex := "[A-Za-z0-9_.-]+"
	path := ""
	param_counts := 9
	for i := 1; i < param_counts; i++ {
		path = fmt.Sprintf("%s/{param%d:%s}", path, i, alpha_neumeric_regex)
		r.HandleFunc(path, controller.HandleRequest)
	}

	path = fmt.Sprintf("%s/{param%d:[A-Za-z0-9_.-/]+}", path, param_counts)
	r.HandleFunc(path, controller.HandleRequest)
	return r
}

func createServerObj(port string, r *mux.Router) http.Server {
	credentialsOk := handlers.AllowCredentials()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})

	return http.Server{
		Addr:        ":" + port,
		Handler:     handlers.CORS(credentialsOk, headersOk, methodsOk)(r),
		IdleTimeout: 0,
		ReadTimeout: 0,
	}
}
