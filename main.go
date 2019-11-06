package main

import (
	"./app"
	"./controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// init router
	r := mux.NewRouter()
	r.Use(app.JwtAuthentication)

	addr := os.Getenv("ADDR") //Get address from .env file, we did not specify any port so this should return an empty string when tested locally
	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if addr == "" {
		addr = "crmback.loldev.ru"
	}
	if port == "" {
		port = "3010"
	}

	headersOk := handlers.AllowedHeaders([]string{"Authorization, Origin, X-Requested-With, Content-Type, Accept"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	credentialsOk := handlers.AllowCredentials()
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/user/signup", controllers.CreateAccount).Methods("POST")
	r.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	r.HandleFunc("/api/user", controllers.Info).Methods("GET")
	r.HandleFunc("/api/lead", controllers.GetLeadsFor).Methods("GET")
	r.HandleFunc("/api/lead", controllers.CreateLead).Methods("POST")
	r.HandleFunc("/api/lead/{id}", controllers.GetLeadFor).Methods("GET")
	r.HandleFunc("/api/sources", controllers.GetSources).Methods("GET")
	//r.HandleFunc("/api/leads/{id}", updateLead).Methods("PUT")
	//r.HandleFunc("/api/leads/{id}", deleteLead).Methods("DELETE")

	log.Fatal(http.ListenAndServe(addr + ":" + port, handlers.CORS(originsOk, headersOk, credentialsOk, methodsOk)(r)))
}

// Get all leads
//func getLeads(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(leads)
//}

// Get single lead
//func getLead(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) // Get params
//	id, err := strconv.Atoi(params["id"])
//	if err == nil {
//		for _, item := range leads {
//			if item.Id == id {
//				json.NewEncoder(w).Encode(item)
//				return
//			}
//		}
//	}
//	json.NewEncoder(w).Encode(&lead.Lead{})
//}

// Create lead
//func createLead(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var lead lead.Lead
//	err := json.NewDecoder(r.Body).Decode(&lead)
//	if err != nil {
//		fmt.Println(err)
//	}
//	lead.Id = rand.Intn(1000)
//	leads = append(leads, lead)
//	json.NewEncoder(w).Encode(lead)
//}

// Update lead
//func updateLead(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	id, err := strconv.Atoi(params["id"])
//	if err == nil {
//		for index, item := range leads {
//			if item.Id == id {
//				leads = append(leads[:index], leads[index+1:]...)
//				var lead lead.Lead
//				err := json.NewDecoder(r.Body).Decode(&lead)
//				if err != nil {
//					fmt.Println(err)
//				}
//				lead.Id = id
//				leads = append(leads, lead)
//				json.NewEncoder(w).Encode(lead)
//			}
//		}
//	}
//	json.NewEncoder(w).Encode(leads)
//}

// Delete lead
//func deleteLead(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	id, err := strconv.Atoi(params["id"])
//	if err == nil {
//		for index, item := range leads {
//			if item.Id == id {
//				leads = append(leads[:index], leads[index+1:]...)
//				break
//			}
//		}
//	}
//	json.NewEncoder(w).Encode(leads)
//}