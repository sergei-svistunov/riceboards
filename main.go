//go:generate gostatic2lib -path static/dist/ -package gostatic -out gostatic/static.go
package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"

	qrbac "github.com/go-qbit/rbac"
	mysql "github.com/go-qbit/storage-mysql"
	"github.com/go-qbit/web/auth"
	"gopkg.qsoa.cloud/service"
	"gopkg.qsoa.cloud/service/qhttp"
	_ "gopkg.qsoa.cloud/service/qmysql"

	"riceboards/api"
	"riceboards/db"
	"riceboards/gostatic"
	"riceboards/rbac"
)

func main() {
	flag.Parse()

	qrbac.SetStorage(&rbac.RbacStorage{})

	rbDb := db.New()
	if dsn := os.Getenv("DSN"); dsn != "" { // Local instance
		if err := rbDb.Connect(dsn); err != nil {
			log.Fatal(err)
		}
	} else { // Cloud instance
		mysql.SqlDriver = "qmysql"
		if err := rbDb.Connect("db"); err != nil {
			log.Fatal(err)
		}
	}
	defer rbDb.Disconnect()

	webAuth := auth.New(os.Getenv("QSOA_WEBAUTH_SALT"), os.Getenv("DEV") != "")

	if os.Getenv("DEV") != "" {
		rbDb.PrintCreateSql(os.Stderr)
	}

	if os.Getenv("DEV") == "" {
		staticH := gostatic.NewHTTPHandler()
		qhttp.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if staticH.GetFile(r.URL.Path) == nil {
				_, file := path.Split(r.URL.Path)
				if strings.IndexByte(file, '.') == -1 {
					r.URL.Path = "/"
				}
			}

			staticH.ServeHTTP(w, r)
		}))
	} else {
		rpURL, err := url.Parse("http://localhost:8081")
		if err != nil {
			log.Fatal(err)
		}

		qhttp.Handle("/", httputil.NewSingleHostReverseProxy(rpURL))
	}

	qhttp.Handle("/api/", http.StripPrefix("/api", api.New(rbDb, webAuth)))

	service.Run()
}
