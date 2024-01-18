package middleware

import (
    "net/http"
    "strings"
)

// AuthMiddleware vérifie la présence et la validité du jeton d'authentification.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Récupère le jeton d'authentification de l'en-tête de la requête
        token := r.Header.Get("Authorization")

        // Vérifie si le jeton est présent et valide
        if isValidToken(token) {
            // Si le jeton est valide, passe la requête au gestionnaire suivant
            next.ServeHTTP(w, r)
        } else {
            // Si le jeton n'est pas valide, renvoie une réponse d'erreur non autorisée
            http.Error(w, "Non autorisé", http.StatusUnauthorized)
        }
    })
}

// isValidToken simule la vérification de la validité du jeton (à personnaliser selon vos besoins)
func isValidToken(token string) bool {
    // Simule une vérification de jeton valide (vous devez implémenter votre propre logique de validation)
    return strings.HasPrefix(token, "Bearer ")
}
