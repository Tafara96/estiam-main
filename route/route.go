package route

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "github.com/Tafara96/estiam-main/dictionary"
    "github.com/Tafara96/estiam-main/middleware"
)

func RoutesDefinition(dictionary *dictionary.Dictionary) *mux.Router {
    router := mux.NewRouter()

    // Ajoute le middleware LoggerMiddleware à toutes les routes
    router.Use(middleware.LoggerMiddleware)

    // Ajoute le middleware AuthMiddleware à toutes les routes nécessitant une authentification
    authenticatedRouter := router.PathPrefix("/dictionary").Subrouter()
    authenticatedRouter.Use(middleware.AuthMiddleware)

    // Route pour ajouter une entrée au dictionnaire (POST)
    router.HandleFunc("/dictionary/add", gestionnairePointEntreeAdd(dictionary)).Methods("POST")

    // Route pour récupérer une définition par mot (GET)
    router.HandleFunc("/dictionary/{mot}", gestionnairePointEntreeGet(dictionary)).Methods("GET")

    // Route pour supprimer une entrée par mot (DELETE)
    router.HandleFunc("/dictionary/{mot}", gestionnairePointEntreeRemove(dictionary)).Methods("DELETE")

    return router
}

// Handlers
func gestionnairePointEntreeAdd(dictionnaire *dictionary.Dictionary) http.HandlerFunc {
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
    		dictionnaire.Add(entry.Word, entry.Definition)

    		w.WriteHeader(http.StatusCreated)
    		w.Write([]byte("Entrée ajoutée avec succès"))
    }
}

func gestionnairePointEntreeGet(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour récupérer une définition ici
        // Récupère le paramètre du mot depuis l'URL de la requête
    		params := mux.Vars(r)
    		mot := params["mot"]

    		// Récupère la définition depuis le dictionnaire
    		entry, err := dictionary.Get(mot)
    		if err != nil {
    			http.Error(w, err.Error(), http.StatusNotFound)
    			return
    		}

    		// Écrit la définition dans la réponse
    		json.NewEncoder(w).Encode(entry)
    }
}

func gestionnairePointEntreeRemove(dictionary *dictionary.Dictionary) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Votre logique pour supprimer une entrée ici
        // Récupère le paramètre du mot depuis l'URL de la requête
    		params := mux.Vars(r)
    		mot := params["mot"]

    		// Supprime l'entrée du dictionnaire
    		dictionary.Remove(mot)

    		w.Write([]byte("Entrée supprimée avec succès"))
    }
}
