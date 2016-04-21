package routes

import  (
  "projects/webApp/controllers/users"
)
//importamos los controladores que se encargaran de entregar una vista cuando apunten a una url especifica
// "projects/webApp/controllers/users" lo tienen que sustituir con su GOPATH y apuntar a donde tengan los controllers

//creamos un nuevo dato que es un array de routers
type Routes []Router

var routes = Routes{
  Router{
    "Index",
    "GET",
    "/",
    usersController.Index,
  },
  Router{
    "Show",
    "GET",
    "/show",
    usersController.Show,
  },
}
