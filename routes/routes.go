package route

import (
    "fmt"
    "encoding/json"
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
        var entry dictionary.Entry

    		// Décode le corps de la requête JSON
    		decoder := json.NewDecoder(r.Body)
    		if err := decoder.Decode(&entry); err != nil {
    			http.Error(w, "Format JSON invalide", http.StatusBadRequest)
    			return
    		}

    		// Ajoute l'entrée au dictionnaire
    		dictionary.Add(entry.Word, entry.Definition)

    		w.WriteHeader(http.StatusCreated)
    		w.Write([]byte("Entrée ajoutée avec succès"))
    }
}

func getDefinitionHandler(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour récupérer une définition ici
        // Récupère le paramètre du mot depuis l'URL de la requête
    		params := mux.Vars(r)
    		word := params["word"]

    		// Récupère la définition depuis le dictionnaire
    		entry, err := dictionary.Get(word)
    		if err != nil {
    			http.Error(w, err.Error(), http.StatusNotFound)
    			return
    		}

    		// Écrit la définition dans la réponse
    		json.NewEncoder(w).Encode(entry)
    }
}

func removeEntryHandler(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour supprimer une entrée ici
        // Récupère le paramètre du mot depuis l'URL de la requête
    		params := mux.Vars(r)
    		word := params["word"]

    		// Supprime l'entrée du dictionnaire
    		dictionary.Remove(word)

    		w.Write([]byte("Entrée supprimée avec succès"))
    }
}
