package routes
//es parte del paquete routes para que sea un archivo que se pueda importar

import ( //importamos las librerias que necesitamos
  "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
  //se crea un nuevo router de la libreria gorilla/mux
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes { //se recorren todas las rutas que tengamos creadas, ver archivo routes.go
    router.
      Methods(route.Method).
      Path(route.Path).
      Name(route.Name).
      Handler(route.HandlerFunc)
      //cada ruta se agrega a router
  }
  return router //se regresan todas las rutas ya guardadas
}
