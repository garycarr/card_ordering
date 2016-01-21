// The following allows running `go generate` to bundle the static and template
// assets into binary form so that the web console can be a stand-alone binary.
// If you update anything in static/ or templates/ be sure to run `go generate`
// and commit the binary differences as well.
//go:generate go get github.com/jteeuwen/go-bindata/...
//go:generate go-bindata -o resources/resources.go -pkg resources -prefix resources -ignore resources/resources.go resources/...
//go:generate gofmt -w resources/resources.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/unrolled/render"

	"github.com/garycarr/card_ordering/resources"
)

// Config ..
type Config struct {
	maxDeckSize   int
	maxShuffles   int
	printEvery    int
	startDeckSize int
	verbose       bool

	// DocumentRoot is used if IsDev==true and references to root directory in which to read the resources from.
	DocumentRoot string

	isDev bool

	port int
}

type app struct {
	server http.Server

	renderer *render.Render

	conf Config

	staticFileHander http.Handler

	// attempts are the number of shuffles made at the current size
	attempts int

	// upToCard is where we are up to
	upToCard int

	stopListening chan bool

	results map[int]int
}

func newApp(c Config) *app {
	if c.DocumentRoot == "" {
		c.DocumentRoot = "./resources"
	}
	ro := render.Options{
		Asset:                     nil,
		AssetNames:                nil,
		Layout:                    "layout",
		Extensions:                []string{".tmpl"},
		Funcs:                     []template.FuncMap{},
		Delims:                    render.Delims{"{{", "}}"},
		Charset:                   "UTF-8",
		IndentJSON:                false,
		IndentXML:                 false,
		PrefixJSON:                []byte(""),
		PrefixXML:                 []byte(""),
		HTMLContentType:           "text/html",
		IsDevelopment:             c.isDev,
		UnEscapeHTML:              false,
		StreamingJSON:             false,
		RequirePartials:           false,
		DisableHTTPErrorRendering: false,
	}
	var fileServer http.Handler
	if c.isDev {
		ro.Directory = filepath.Join(c.DocumentRoot, "templates")
		fileServer = http.FileServer(http.Dir(c.DocumentRoot))
	} else {
		ro.Directory = "templates"
		ro.Asset = resources.Asset
		ro.AssetNames = resources.AssetNames
		fileServer = http.FileServer(&assetfs.AssetFS{Asset: resources.Asset, AssetDir: resources.AssetDir, AssetInfo: resources.AssetInfo, Prefix: ""})
	}

	return &app{
		conf:             c,
		staticFileHander: fileServer,
		renderer:         render.New(ro),
		results:          make(map[int]int),
		upToCard:         c.startDeckSize,
		stopListening:    make(chan bool, 1),
	}
}

func (a *app) start() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := a.conf.startDeckSize; i < a.conf.maxDeckSize; i++ {
			a.attempts = 0
			a.shuffleAndCheck(i)
		}
	}()
	a.registerEndpoints()

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			// Server closed is an intentional close
			if !strings.Contains(err.Error(), "Server closed") {
				panic(err)
			}

		}
	}()
	wg.Wait()
}

func (a *app) registerEndpoints() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(a.IndexHandlerGET))
	mux.Handle("/about", http.HandlerFunc(a.AboutHandlerGET))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources/static"))))

	// NOTE: figure this out
	addr := "0.0.0.0"
	a.server = http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf("%s:%d", addr, a.conf.port),
	}
}
