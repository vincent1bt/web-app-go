package routes
//lo agregamos al mismo paquete para que puedan comunicarse entre si, y tengamos la estructura
//disponible cuando importemos el paquete

import (
  "net/http"
)

type Router struct {
  Name string
  Method string
  Path string
  HandlerFunc http.HandlerFunc
}
//creamos una estructura nueva que tenga los valores que las rutas necesitan
