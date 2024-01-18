package dictionary

import (
	"fmt"
	"sort"
)

//L'entrée représente une entrée dans le dictionnaire avec un mot et sa définition
type Entry struct {
	Definition string
}

func (e Entry) String() string {

	return e.Definition //Puisse que le retour de doit etre un string et que la definition de entry est un string
}

type Dictionary struct {
	entries map[string]Entry
}

//New crée une nouvelle instance de Dictionary
func New() *Dictionary {

	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

//Ajouter ajoute un mot et sa définition au dictionnaire
func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Definition: definition}
  d.entries[word] = entry
}

//Get récupère l'entrée d'un mot spécifique dans le dictionnaire
func (d *Dictionary) Get(word string) (Entry, error) {

	entry, exists := d.entries[word]
	if !exists {
		return Entry{}, fmt.Errorf("word not found: %s", word)
	}
	return entry, nil
}

//Remove supprime un mot et son entrée du dictionnaire
func (d *Dictionary) Remove(word string) {
  delete(d.entries, word)
}

//List renvoie une liste triée de mots et leurs entrées du dictionnaire
func (d *Dictionary) List() ([]string, map[string]Entry) {

	var wordList []string
	for word := range d.entries {
		wordList = append(wordList, word)
	}

	sort.Strings(wordList)

	var entriesMap = make(map[string]Entry)
	for _, word := range wordList {
		entriesMap[word] = d.entries[word]
	}

	return wordList, entriesMap
}
