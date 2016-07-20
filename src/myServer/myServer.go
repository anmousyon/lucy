package myServer

import(
    "io"
    "net/http"
    "fmt"
)

func hello(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello world!")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

//StartServer starts a server on localhost:8000
func StartServer(){
    fmt.Println("starting server")
    server := http.Server{
        Addr: ":8000",
        Handler: &myHandler{},
    }
    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = hello
    server.ListenAndServe()
}

type myHandler struct{}

func(*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if h, ok := mux[r.URL.String()]; ok{
        h(w, r)
        return
    }
    io.WriteString(w, "My server: "+r.URL.String())
}
