package route

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "github.com/Tafara96/estiam-main/dictionary"
)

func SetupRoutes(dictionary *dictionary.Dictionary) *mux.Router {
    router := mux.NewRouter()

    // Route pour ajouter une entrée au dictionnaire (POST)
    router.HandleFunc("/dictionary/add", addEntryHandler(dictionary)).Methods("POST")

    // Route pour récupérer une définition par mot (GET)
    router.HandleFunc("/dictionary/{word}", getDefinitionHandler(dictionary)).Methods("GET")

    // Route pour supprimer une entrée par mot (DELETE)
    router.HandleFunc("/dictionary/{word}", removeEntryHandler(dictionary)).Methods("DELETE")

    return router
}

// Handlers
func addEntryHandler(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour ajouter une entrée ici
    }
}

func getDefinitionHandler(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour récupérer une définition ici
        fmt.Sprintf("Yes")
    }
}

func removeEntryHandler(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour supprimer une entrée ici
    }
}
