package dictionary

import (
	"bufio"
  "fmt"
  "os"
  "sort"
  "strings"
	"errors"
)

//L'entrée représente une entrée dans le dictionnaire avec un mot et sa définition
type Entry struct {
	Word string `json:"mot"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	filename string
}

//New crée une nouvelle instance de Dictionary
func New(filename string) *Dictionary {

	return &Dictionary{
        filename: filename,
    }
}

//Ajouter ajoute un mot et sa définition au dictionnaire
func (d *Dictionary) Add(mot string, definition string) error {

	  //Verification de la validité des données avant de les ajouter
    if err := verifierDonnees(mot, definition); err != nil {
        return err
    }

	fichier, err := os.OpenFile(d.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer fichier.Close()

    entry := fmt.Sprintf("%s: %s\n", mot, definition)
    _, err = fichier.WriteString(entry)
    return err
}

//Get récupère l'entrée d'un mot spécifique dans le dictionnaire
func (d *Dictionary) Get(mot string) (Entry, error) {

	fichier, err := os.Open(d.filename)
    if err != nil {
        return Entry{}, err
    }
    defer fichier.Close()

    scanner := bufio.NewScanner(fichier)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 && strings.TrimSpace(parts[0]) == mot {
            return Entry{Word: strings.TrimSpace(parts[0]), Definition: strings.TrimSpace(parts[1])}, nil
        }
    }

    return Entry{}, fmt.Errorf("mot introuvable: %s", mot)
}

//Remove supprime un mot et son entrée du dictionnaire
func (d *Dictionary) Remove(mot string) error {
	lignes, err := readLines(d.filename)
    if err != nil {
        return err
    }

    var updatedLines []string
    for _, line := range lignes {
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 && strings.TrimSpace(parts[0]) == mot {
            continue //Passer la ligne pour supprimer l'entrée
        }
        updatedLines = append(updatedLines, line)
    }

    return writeLines(d.filename, updatedLines)
}

//List renvoie une liste triée de mots et leurs entrées du dictionnaire
func (d *Dictionary) List() ([]string, map[string]Entry) {

	lignes, err := readLines(d.filename)
    if err != nil {
        return nil, nil
    }

    var listeDeMots []string
    entriesMap := make(map[string]Entry)

    for _, line := range lignes {
        parts := strings.SplitN(line, ":", 2)
        if len(parts) == 2 {
            mot := strings.TrimSpace(parts[0])
            definition := strings.TrimSpace(parts[1])

            listeDeMots = append(listeDeMots, mot)
            entriesMap[mot] = Entry{Definition: definition}
        }
    }

    sort.Strings(listeDeMots)
    return listeDeMots, entriesMap
}

// verifierDonnees verifier si les donnees que l'on veut ajouter correspond au regles definis dans la function si non on retourne une erreur
func verifierDonnees(mot string, definition string) error {
    // Si la longueur du mot est inférieur à 3 on refuse l'entrée
    if len(mot) < 3 {
        return errors.New("La longueur du mot est inférieur à 3")
    }

    // Si la longueur de la definition est inférieur à 5 on refuse l'entrée
    if len(definition) < 5 {
        return errors.New("La longueur de la definition ne doit pas etre inférieur à 3")
    }

    return nil
}

//Fonction pour lire les lignes d'un fichier
func readLines(filename string) ([]string, error) {
    fichier, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer fichier.Close()

    var lignes []string
    scanner := bufio.NewScanner(fichier)
    for scanner.Scan() {
        lignes = append(lignes, scanner.Text())
    }

    return lignes, scanner.Err()
}

//Fonction pour écrire des lignes dans un fichier
func writeLines(filename string, lignes []string) error {
    fichier, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer fichier.Close()

    writer := bufio.NewWriter(fichier)
    for _, line := range lignes {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }

    return writer.Flush()
}
