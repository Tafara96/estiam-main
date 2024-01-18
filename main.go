package main

import (
	"bufio"
  "fmt"
	"github.com/Tafara96/estiam-main/dictionary"
)

func main() {
	//Créer un nouveau dictionnaire
	myDictionary := dictionary.New()

	//Ajouter des entrées au dictionnaire
	myDictionary.Add("apple", "a fruit")
	myDictionary.Add("gopher", "the Go mascot")
	myDictionary.Add("map", "a collection of key-value pairs")

	//Récupère et affiche la définition d'un mot spécifique
	wordToGet := "gopher"
	entry, err := myDictionary.Get(wordToGet)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Definition of '%s': %s\n", wordToGet, entry)
	}

	//Supprimer un mot du dictionnaire
	wordToRemove := "apple"
	myDictionary.Remove(wordToRemove)
	fmt.Printf("'%s' removed from the dictionary\n", wordToRemove)

	//Liste tous les mots et leurs entrées dans le dictionnaire
	wordList, entries := myDictionary.List()
	fmt.Println("List of words in the dictionary:")
	for _, word := range wordList {
		fmt.Printf("%s: %s\n", word, entries[word])
	}
}

/*func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionList(d *dictionary.Dictionary) {

}*/
