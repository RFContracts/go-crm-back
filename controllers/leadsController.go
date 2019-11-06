package controllers

import (
	"../models"
	u "../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateLead = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("qwe")
	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	lead := &models.Lead{}
	fmt.Println(user)
	err := json.NewDecoder(r.Body).Decode(lead)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	lead.UserId = user
	resp := lead.Create()
	fmt.Println(resp)
	u.Respond(w, resp)
}

var GetLeadsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user") . (uint)
	data := models.GetLeads(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetLeadFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
	}
	data := models.GetLead(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetSources = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetSources()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}