package main

import(
	"net/http"
  "fmt"
)


func middleware(originalHandler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Running before handler")

        w.Write([]byte("Hijacking Request "))
        originalHandler.ServeHTTP(w, r)

        fmt.Println("Running after handler")
  })
}


func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Executing the handler")
  w.Write([]byte("OK"))
}


func main() {
    // converting our handler function to handler 
    // type to make use of our middleware 
    myHandler := http.HandlerFunc(handler)
    http.Handle("/", middleware(myHandler)) 
    http.ListenAndServe(":8000", nil)
}




