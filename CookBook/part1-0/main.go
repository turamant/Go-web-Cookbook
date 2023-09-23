package main

    import (
        "crypto/subtle"
        "fmt"
        "log"
        "net/http"
    )

    const (
        CONN_HOST = "localhost"
        CONN_PORT = "8080"
        ADMIN_USER = "admin"
        ADMIN_PASSWORD = "admin"
    )

    func helloWorld(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World!")
    }

    func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request){
            user, pass, ok := r.BasicAuth()
            if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || 
						subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
                w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
                w.WriteHeader(401)
                w.Write([]byte("Вы не авторизованы для доступа к приложению.\n"))
                return
            }
            handler(w, r)
        }
    }

    func main(){
        http.HandleFunc("/", BasicAuth(helloWorld, "Пожалуйста введите свой логин и пароль"))
        err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
        if err != nil {
            log.Fatal("error starting http server : ", err)
            return
        }
    }