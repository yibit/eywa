package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"eywa/browser"
	"eywa/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:        "serve",
	Usage:       "Serve the docs",
	Description: "Serve the docs",
	Action:      serve,
	Flags: []cli.Flag{
		utils.StringFlag("directory", "d", "docs", "directory of server"),
		utils.StringFlag("port", "p", "1036", "port of server"),
		utils.BoolFlag("fix", "f", false, "`fix` path of URL (whether add `.html` to path)"),
		utils.BoolFlag("open", "o", false, "`open` the browser"),
	},
}

func serve(ctx *cli.Context) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, ctx.String("directory")))
	FileServer(r, "/", filesDir, ctx.Bool("fix"))

	url := "http://localhost:" + ctx.String("port")
	log.Printf("serve: %s", url)

	if ctx.Bool("open") {
		go func() {
			if !browser.Open(url) {
				log.Printf("Failed to open browser window. Please visit %s in your browser.", url)
			}
		}()
	}

	log.Fatal(http.ListenAndServe(":"+ctx.String("port"), r))

	return nil
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem, fix bool) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		if fix && !strings.Contains(r.URL.Path, ".css") &&
			!strings.Contains(r.URL.Path, ".js") &&
			!strings.Contains(r.URL.Path, ".html") {
			texts := strings.Split(r.URL.Path, "/")
			if n := len(texts); n > 2 {
				r.URL.Path = "/" + strings.Join(texts[2:], "/") + ".html"
			} else {
				r.URL.Path += ".html"
			}

			fmt.Printf("XPath:%s\n", r.URL.Path)
		}

		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
