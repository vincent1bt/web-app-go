package usersController
//para tener un orden le llamaremos usersController

import (
  "net/http"
  "fmt"
)

//creamos las funciones que corresponden a las rutas
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Index users")
}

func Show(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Show users")
}
