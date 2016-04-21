package main

import (
  "log"
  "net/http"
  "projects/webApp/routes"
)

func main() {
  router := routes.NewRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}
