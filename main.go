package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	godotenv.Load()
	// dbURL := os.Getenv("DB_URL")
	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	os.Exit(0x1)
	// }
	// dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"
	apiCfg := &apiConfig{}

	handler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))

	mux := http.NewServeMux()

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	hits := cfg.fileserverHits.Load()
	html := fmt.Sprintf(`
		<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
		</html>
	`, hits)
	w.Write([]byte(html))
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Hits reset to 0"))
}

type ChirpRequest struct {
	Body string `json:"body"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ValidResponse struct {
	Valid bool `json:"valid"`
}

type CleanedResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

var profaneWords = []string{"kerfuffle", "sharbert", "fornax"}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := ChirpRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid JSON"})
		return
	}

	if len(req.Body) > 140 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Chirp is too long"})
		return
	}

	words := strings.Split(req.Body, " ")
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		for _, profane := range profaneWords {
			if lowerWord == profane {
				words[i] = "****"
			}
		}
	}
	cleaned := strings.Join(words, " ")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CleanedResponse{CleanedBody: cleaned})
}

// package main

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"
// 	"os"
// 	"sync/atomic"

// 	"github.com/bootdotdev/learn-http-servers/internal/database"
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// )

// type apiConfig struct {
// 	fileserverHits atomic.Int32
// 	db             *database.Queries
// }

// func main() {
// 	const filepathRoot = "."
// 	const port = "8080"

// 	godotenv.Load()
// 	dbURL := os.Getenv("DB_URL")
// 	if dbURL == "" {
// 		log.Fatal("DB_URL must be set")
// 	}

// 	dbConn, err := sql.Open("postgres", dbURL)
// 	if err != nil {
// 		log.Fatalf("Error opening database: %s", err)
// 	}
// 	dbQueries := database.New(dbConn)

// 	apiCfg := apiConfig{
// 		fileserverHits: atomic.Int32{},
// 		db:             dbQueries,
// 	}

// 	mux := http.NewServeMux()
// 	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
// 	mux.Handle("/app/", fsHandler)

// 	mux.HandleFunc("GET /api/healthz", handlerReadiness)
// 	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

// 	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
// 	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)

// 	srv := &http.Server{
// 		Addr:    ":" + port,
// 		Handler: mux,
// 	}

// 	log.Printf("Serving on port: %s\n", port)
// 	log.Fatal(srv.ListenAndServe())
// }
