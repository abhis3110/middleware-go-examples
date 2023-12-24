package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*func middleware(originalHandler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Running before handler")// server log

        w.Write([]byte("Hijacking Request ")) // client response
        originalHandler.ServeHTTP(w, r)

        fmt.Println("Running after handler")
  })
} */

//Write a middleware that makes sure request has Header "Content-Type" application/json
func filterContentType(handler http.Handler)http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Inside first Middleware")
    if r.Header.Get("Content-Type")!="application/json"{
      w.WriteHeader(http.StatusUnsupportedMediaType)
      w.Write([]byte("405-header content type incorrect"))
      return
    }
    handler.ServeHTTP(w,r)
  })
}



//Write a middleware that adds current server time to the reponse cookie
func setCookieTime(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Inside second Middleware")
    // Here Cookie is a struct that represents an HTTP cookie
    // as sent in the Set-Cookie header of HTTP request
    cookie:=http.Cookie{
      Name:"server time (UTC)",
      Value : strconv.Itoa(int(time.Now().Unix())), // convert time to string
    }
    // now set the cookie to response
    http.SetCookie(w, &cookie)
    handler.ServeHTTP(w,r)

  })

}



// handler for input to middleware
func postHandler(w http.ResponseWriter, r *http.Request){
  log.Println("Inside Handler")
  if r.Method!="POST"{
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte("405 - Method Not Allowed"))
    return
  }

  var p Person
  err:= json.NewDecoder(r.Body).Decode(&p)
  if err!=nil{
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Internal Server Error"))
    return
  }
  defer r.Body.Close()

  //
  fmt.Printf("Got firstName and lastName as %s, %s", p.Firstname, p.Lastname)
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("201 - Created"))
}
 
// basic handler 
/*func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Executing the handler")
  w.Write([]byte("OK"))
}*/

type Person struct{
  Firstname string
  Lastname string
}

func main() {
    // converting our handler function to handler 
    // type to make use of our middleware 

   /* myHandler := http.HandlerFunc(handler)
    http.Handle("/", middleware(myHandler)) */

    myHandler := http.HandlerFunc(postHandler)
    chain := filterContentType(setCookieTime(myHandler))
    http.Handle("/",chain)

    http.ListenAndServe(":8000", nil)
}




