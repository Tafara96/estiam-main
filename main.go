package main

import (
  "fmt"
	"github.com/Tafara96/estiam-main/dictionary"
)

func main() {
	//Créer un nouveau dictionnaire
	leDictionnaire := dictionary.New()

	//Ajouter des entrées au dictionnaire
	leDictionnaire.Add("tesla", "une voiture")
	leDictionnaire.Add("arafat", "une personne")
	leDictionnaire.Add("paris", "une ville")

	//Récupère et affiche la définition d'un mot spécifique
	motRecherche := "arafat"
	entry, err := leDictionnaire.Get(motRecherche)
	if err != nil {
		fmt.Printf("Erreur: %s\n", err)
	} else {
		fmt.Printf("La definition de '%s' est %s\n", motRecherche, entry)
	}

	//Supprimer un mot du dictionnaire
	motIndesire := "tesla"
	leDictionnaire.Remove(motIndesire)
	fmt.Printf("'%s' à été retiré du dictionnaire \n", motIndesire)

	//Liste tous les mots et leurs entrées dans le dictionnaire
	listeDeMots, entries := leDictionnaire.List()
	fmt.Println("Liste triée des mots du dictionnaire:")
	for _, word := range listeDeMots {
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
