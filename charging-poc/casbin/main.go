package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{{ID: 1, Name: "John", Age: 30}, {ID: 2, Name: "Doe", Age: 25}}
var enforcer *casbin.Enforcer

func main() {
	enforcer = initEnforcer()
	router := mux.NewRouter()
	router.HandleFunc("/show-user-details", authorize(showUserDetails, enforcer)).Methods("POST")
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcerSafe("./ex/auth_model.conf", "./ex/auth_policy.csv")
	if err != nil {
		log.Fatal("it is here", err)
	}
	return e
}

func showUserDetails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func authorize(next http.HandlerFunc, e *casbin.Enforcer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the role from the request body
		var userRole struct {
			Role string `json:"role"`
		}

		err := json.NewDecoder(r.Body).Decode(&userRole)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Check the authorization policy using Casbin
		if ok, _ := e.EnforceSafe(userRole.Role, "showUserDetails", "POST"); ok {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Authorization error", http.StatusUnauthorized)
		}
	}
}
