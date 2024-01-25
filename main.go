package main

import (
	"fmt"
	"net/http"
  "github.com/Tafara96/estiam-main/route"
	"github.com/Tafara96/estiam-main/dictionary"
)

func main() {
	//Création d'un nouveau dictionnaire avec un nom de fichier
  leDictionnaire := dictionary.New("dictionary.txt")

	fmt.Sprintf("Serveur Démarré");

  //Mise en place des routes
  router := route.RoutesDefinition(leDictionnaire)

  //Demmarer le serveur
  http.Handle("/", router)
  http.ListenAndServe(":8080", nil)
}

/*func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionList(d *dictionary.Dictionary) {

}*/
