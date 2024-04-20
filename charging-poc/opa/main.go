package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/open-policy-agent/opa/rego"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{{ID: 1, Name: "John", Age: 30}, {ID: 2, Name: "Doe", Age: 25}}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/show-user-details", authorize(showUserDetails)).Methods("POST")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func showUserDetails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the policy from file
		policyBytes, err := ioutil.ReadFile("auth_policy.rego")
		if err != nil {
			http.Error(w, "Failed to read policy file", http.StatusInternalServerError)
			return
		}

		// Compile the policy
		regoModule, err := rego.New(
			rego.Query("data.authz.allow"),
			rego.Module("authz.rego", string(policyBytes)),
		).PrepareForEval(context.Background())

		if err != nil {
			http.Error(w, "Failed to compile policy", http.StatusInternalServerError)
			return
		}

		// Extract the role from the request body
		var userRole struct {
			Role string `json:"role"`
		}

		err = json.NewDecoder(r.Body).Decode(&userRole)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Evaluate the policy
		resultSet, err := regoModule.Eval(context.Background(), rego.EvalInput(map[string]interface{}{
			"sub":  userRole.Role,
			"obj":  "showUserDetails",
			"act":  "POST",
			"user": users,
		}))

		if err != nil {
			http.Error(w, "Failed to evaluate policy", http.StatusInternalServerError)
			return
		}

		if len(resultSet) > 0 && resultSet[0].Expressions[0].Value.(bool) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Authorization error", http.StatusUnauthorized)
		}
	}
}
