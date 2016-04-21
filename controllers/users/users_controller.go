package usersController
//para tener un orden le llamaremos usersController

import (
  "net/http"
  "html/template" //agregamos la libreria de templates
  "projects/webApp/models/user"
)

//creamos las funciones que corresponden a las rutas
func Index(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("views/users/index.html") //indicamos la ruta del template/html
  t.Execute(w, nil) //se hace render del template
}

func Show(w http.ResponseWriter, r *http.Request) {
  user := &userModel.User{ Name: "Alberto", Email: "albert@email.com" } //creamos un nuevo usuario
  t, _ := template.ParseFiles("views/users/show.html") //indicamos a ruta del template
  t.Execute(w, user) //le pasamos al usuario al template para que este disponible
}
