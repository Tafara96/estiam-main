package middleware

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

// LoggerMiddleware enregistre les informations de la requête dans un fichier log.
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Enregistre l'heure actuelle, la méthode HTTP et le chemin de la requête
        logMessage := fmt.Sprintf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)

        // Enregistre le message dans le fichier journal
        logFile, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            // En cas d'erreur, continue la requête sans journalisation
            fmt.Println("Erreur lors de l'ouverture du fichier log :", err)
            next.ServeHTTP(w, r)
            return
        }
        defer logFile.Close()

        _, err = logFile.WriteString(logMessage)
        if err != nil {
            fmt.Println("Erreur lors de l'écriture dans le fichier log :", err)
        }

        // Passe la requête au gestionnaire suivant
        next.ServeHTTP(w, r)
    })
}
