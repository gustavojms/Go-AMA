package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/gustavojms/go-ama/internal/store/pgstore"
	"net/http"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()

	a.r = r
	return a
}
