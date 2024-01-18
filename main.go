package main

import (
	"fmt"
	"net/http"
  "github.com/Tafara96/estiam-main/route"
	"github.com/Tafara96/estiam-main/dictionary"
)

func main() {
	//Création d'un nouveau dictionnaire avec un nom de fichier
  myDictionary := dictionary.New("dictionary.txt")

  //Mise en place des routes
  router := route.SetupRoutes(myDictionary)

  //Demmarer le serveur
	fmt.Sprintf("Serveur Démarré")
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
