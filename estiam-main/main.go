package main

import (
	"bufio"

	"github.com/Tafara96/estiam-main/tree/main/estiam-main/dictionary"
)

func main() {
	// Création d'un nouveau dictionnaire
	dictionnaire := dictionary.New()

	// Ajout des mots et des définitions
	dictionnaire.Add("element1", "Ce-ci est l'element 1")
	dictionnaire.Add("element2", "Ce-ci est l'element 2")
	dictionnaire.Add("element3", "Ce-ci est l'element 3")

	// Afficher la définition d'un mot spécifique
	element := "element1"
	entree, err := dictionnaire.Get(element)
	if err == nil {
		fmt.Printf("Definition of '%s': %s\n", element, entree)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	// Supprime un mot du dictionnaire
	wordToRemove := "slice"
	dictionnaire.Remove(wordToRemove)
	fmt.Printf("'%s' removed from the dictionary.\n", wordToRemove)

	// Lister les mots et leurs définitions
	wordList, entries := dictionnaire.List()
	sort.Strings(wordList) // Trie les mots à afficher

	fmt.Println("\nDictionary Entries:")
	for _, word := range wordList {
		fmt.Printf("%s: %s\n", word, entries[word])
	}
}
