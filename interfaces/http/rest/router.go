/*
|--------------------------------------------------------------------------
| Router
|--------------------------------------------------------------------------
|
| This file contains the routes mapping and groupings of your REST API calls.
| See README.md for the routes UI server.
|
*/
package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"gomora-dapp/interfaces"
	"gomora-dapp/interfaces/http/rest/middlewares/cors"
	"gomora-dapp/interfaces/http/rest/viewmodels"
)

// ChiRouterInterface declares methods for the chi router
type ChiRouterInterface interface {
	InitRouter() *chi.Mux
	Serve(port int)
}

type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

// InitRouter initializes main routes
func (router *router) InitRouter() *chi.Mux {
	// DI assignment
	nftQueryController := interfaces.ServiceContainer().RegisterNFTRESTQueryController()

	// create router
	r := chi.NewRouter()

	// global and recommended middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(cors.Init().Handler)
	r.Use(middleware.Recoverer)

	// default route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := viewmodels.HTTPResponseVM{
			Status:  http.StatusOK,
			Success: true,
			Message: "alive",
			Data:    map[string]string{"version": os.Getenv("API_VERSION")},
		}

		response.JSON(w)
	})

	// docs routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth(os.Getenv("API_NAME"), map[string]string{
			"sudo": os.Getenv("OPENAPI_DOCS_PASSWORD"),
		}))

		workDir, _ := os.Getwd()
		docsDir := http.Dir(filepath.Join(workDir, "docs"))
		FileServer(r, "/docs", docsDir)
	})

	// API routes
	r.Group(func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// routes for nft
			r.Route("/nft", func(r chi.Router) {
				r.Get("/greeting/latest", nftQueryController.GetGreeting)
				r.Get("/greeting/logs", nftQueryController.GetGreeterContractEventLogs)
				// nft-related routes (example only)
				// r.Get("/metadata/{tokenID}", nftQueryController.GetNFTByID)
				// r.Get("/images/{fileName}", nftQueryController.GetNFTImage)

				// FIXME: should not be public
				r.Get("/replay", func(w http.ResponseWriter, r *http.Request) {
					fromBlockString := r.URL.Query().Get("fromBlock")
					toBlockString := r.URL.Query().Get("toBlock")

					fromBlock, _ := strconv.Atoi(fromBlockString)
					toBlock, _ := strconv.Atoi(toBlockString)

					err := interfaces.GreeterEventListenerReplayer(int64(fromBlock), int64(toBlock))
					if err != nil {
						response := viewmodels.HTTPResponseVM{
							Status:  http.StatusInternalServerError,
							Success: false,
							Message: "Cannot run Greeter contract replayer.",
							Data:    err.Error(),
						}

						response.JSON(w)
						return
					}

					response := viewmodels.HTTPResponseVM{
						Status:  http.StatusOK,
						Success: true,
						Message: "Successfully replayed Greeter events.",
					}

					response.JSON(w)
				})
			})
		})
	})

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func (router *router) Serve(port int) {
	log.Printf("[SERVER] REST server running on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router.InitRouter())
	if err != nil {
		log.Fatalf("[SERVER] REST server failed %v", err)
	}
}

func registerHandlers() {}

// ChiRouter export instantiated chi router once
func ChiRouter() ChiRouterInterface {
	if m == nil {
		routerOnce.Do(func() {
			// register http handlers
			registerHandlers()

			m = &router{}
		})
	}

	return m
}
