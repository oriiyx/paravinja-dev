package router

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oriiyx/paravinja-dev/internal/handlers"
)

type Controller struct {
	Router *chi.Mux
}

func (c *Controller) RegisterUses() {
	c.Router.Use(middleware.Logger)
}

func (c *Controller) RegisterRoutes() {
	c.Router.Handle("/*", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))
	c.Router.Get("/", handlers.HomepageIndex())
	c.Router.Get("/contact", handlers.ContactIndex())
	c.Router.Get("/about", handlers.AboutIndex())
	c.Router.Get("/posts/{slug}", handlers.PostsIndex())
}
