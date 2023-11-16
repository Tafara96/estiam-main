package dictionary

import (
	"errors"
	"sort"
	"strings"
)

// L'entrée représente une entrée de dictionnaire avec une définition.
type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

// Dictionary représente un simple dictionnaire.
type Dictionary struct {
	entries map[string]Entry
}

// New crée un nouveau dictionnaire.
func New() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

// Ajouter ajoute un mot et sa définition au dictionnaire.
func (d *Dictionary) Add(word, definition string) {
	word = strings.ToLower(word)
	d.entries[word] = Entry{Definition: definition}
}

// Get récupère l'entrée d'un mot donné dans le dictionnaire.
func (d *Dictionary) Get(word string) (Entry, error) {
	word = strings.ToLower(word)
	entry, exists := d.entries[word]
	if !exists {
		return Entry{}, errors.New("La clé est introuvable dans la dictionnaire")
	}
	return entry, nil
}

// Remove supprime un mot et sa définition du dictionnaire.
func (d *Dictionary) Remove(word string) {
	word = strings.ToLower(word)
	delete(d.entries, word)
}

// List renvoie une liste de mots et leurs entrées correspondantes dans le dictionnaire.
func (d *Dictionary) List() ([]string, map[string]Entry) {
	wordList := make([]string, 0, len(d.entries))
	for word := range d.entries {
		wordList = append(wordList, word)
	}
	return wordList, d.entries
}
