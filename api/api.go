package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-qbit/rbac"
	"github.com/go-qbit/rpc"
	"github.com/go-qbit/rpc/openapi"
	webAuth "github.com/go-qbit/web/auth"

	mAuthGoogle "riceboards/api/method/auth/google"
	mAuthMe "riceboards/api/method/auth/me"
	mIdeasAdd "riceboards/api/method/ideas/add"
	mIdeasDelete "riceboards/api/method/ideas/delete"
	mIdeasEdit "riceboards/api/method/ideas/edit"
	mIdeasList "riceboards/api/method/ideas/list"
	mIdeasOptions "riceboards/api/method/ideas/options"
	mProjectsAdd "riceboards/api/method/projects/add"
	mProjectsEdit "riceboards/api/method/projects/edit"
	mProjectsGet "riceboards/api/method/projects/get"
	mProjectsList "riceboards/api/method/projects/list"
	"riceboards/authctx"
	"riceboards/db"
)

func New(db *db.Db, auth *webAuth.Auth) http.Handler {
	gamesRpc := rpc.New("riceboards/api/method")

	if err := gamesRpc.RegisterMethods(
		mAuthGoogle.New(db, auth),
		mAuthMe.New(db),
		mProjectsAdd.New(db),
		mProjectsGet.New(db),
		mProjectsEdit.New(db),
		mProjectsList.New(db),
		mIdeasAdd.New(db),
		mIdeasDelete.New(db),
		mIdeasEdit.New(db),
		mIdeasList.New(db),
		mIdeasOptions.New(db),
	); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if token := r.Header.Get("X-API-Key"); token != "" {
			uid, err := auth.GetUserId(token)
			if err == nil && uid > 0 {
				user, err := db.Users.GetUserById(r.Context(), strconv.FormatUint(uint64(uid), 10))
				if err != nil {
					log.Printf("Cannot get user: %v", err)
				} else {
					u := user.(authctx.User)
					ctx = authctx.ToContext(ctx, &u)
				}
			}
		}

		ctx, _ = rbac.ContextWithPermissions(ctx, nil)
		gamesRpc.ServeHTTP(w, r.WithContext(ctx))
	})

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swagger := gamesRpc.GetSwagger(r.Context())
		swagger.Info.Title = "Football AI bot API"
		swagger.Servers = []openapi.Server{
			{"/api", ""},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(swagger); err != nil {
			log.Printf("Cannot marshal swagger.json: %v", err)
			return
		}
	})

	mux.HandleFunc("/index.ts", typeScriptLibHandler(gamesRpc))

	return mux
}
