package main

import (
	"net/http"
  "github.com/Tafara96/estiam-main/routes"
	"github.com/Tafara96/estiam-main/dictionary"
)

func main() {
	// Create a new dictionary with a filename
  myDictionary := dictionary.New("dictionary.txt")

  // Setup routes
  router := route.SetupRoutes(myDictionary)

  // Start the server
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
