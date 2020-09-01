package main
import (
  "net/http"
  "log"
  "fmt"
)

func main() {
  log.Print("starting server...")
	http.HandleFunc("/", run)
  log.Printf("listening on port %s", "8000")
  http.ListenAndServe(":8000", nil)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8000"), nil))
}

func run(w http.ResponseWriter, r *http.Request) {
  sqrt(0.0001, 100000);  
}