package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAdr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAdr,
		store:      store,
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS method (preflight request)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Static(r *mux.Router, pathPrefix string, dir string) {
	r.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix, http.FileServer(http.Dir(dir))))
}

// router
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/product", makeHandleFunc(s.handleProduct))
	router.HandleFunc("/product/{id}", makeHandleFunc(s.handleProductByID))
	router.HandleFunc("/media", makeHandleFunc(s.handleMedia))
	router.HandleFunc("/media/{id}", makeHandleFunc(s.handleMediaByID))
    router.HandleFunc("/user", makeHandleFunc(s.handleAccount))

	Static(router, "/uploads/", "./uploads")
	Static(router, "/medias/", "./medias")

	http.ListenAndServe(s.listenAddr, corsMiddleware(router))
}

// get ID
func getID(r *http.Request) string {
	id := mux.Vars(r)["id"]

	return id
}

// parse error
func badRequestError(w http.ResponseWriter, errStr string) error {
	return WriteJSON(w, http.StatusBadRequest, ApiError{Error: errStr})
}

// server error
func serverError(w http.ResponseWriter, errStr string) error {
	return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: errStr})
}

// string to float
func stringToFloat(value string) (float64, error) {
	valueParsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return valueParsed, nil
}

// string to int
func stringToInt(value string) (int64, error) {
	valueParsed, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return 0, err
	}
	return valueParsed, nil
}

// string to bool
func stringToBool(value string) (bool, error) {
	valueParsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return valueParsed, nil
}

// json output
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

// api error
type ApiError struct {
	Error string `json:"error"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// handle func wrapper
func makeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
