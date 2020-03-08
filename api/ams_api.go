package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SendUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
}

func SendPolicies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write(policiesJson)
}

var usersJson = []byte(`[
	{	
        "ID": 1, 
      	"name": "user1", 
       	"mobile_number": "1234567890"
     },	
	{	
        "ID": 2, 
      	"name": "user2", 
      	"mobile_number": "123 456 7891"
	},
 	{	
        "ID": 3, 
     	"name": "user3", 
     	"mobile_number": "(123) 456 7892"
	},
	{	
        "ID": 4, 
     	"name": "user4", 
     	"mobile_number": "(123) 456-7893"
	},
	{	
        "ID": 5, 
    	"name": "user5", 
		"mobile_number": "123-456-7894"
	}
   ]`)

var policiesJson = []byte(`[
     {	
        "ID": 1, 
        "mobile_number": "1234567890", 
	 	"premium": 2000, 
		"type": "homeowner"
     },	
	 {
        "ID": 2, 
      	"mobile_number": "123 456 7891", 
		"premium": 200, 
		"type": "renter"
	},
 	{
        "ID": 3, 
    	"mobile_number": "123-456-7892", 
		"premium": 1500, 
		"type": "homeowner"
	},
	{
        "ID": 4, 
    	"mobile_number": "(123) 456-7893",  
		"premium": 155, 
		"type": "personal_auto"
	},
	{
        "ID": 5, 	
   		"mobile_number": "123-456-7894", 
		"premium": 1000, 
		"type": "homeowner"
	},
	{
       "ID": 6, 
   		"mobile_number": "123-456-7890", 
		"premium": 500, 
		"type": "personal_auto"
	},
	{
        "ID": 7, 
   		"mobile_number": "1234567892",  
		"premium": 100, 
		"type": "life"
	},
	 {
		"ID": 8, 
		"mobile_number": "(123)456-7892", 
		"premium": 200, 
		"type": "homeowner"
	}
]`)

func main() {
	router := httprouter.New()
	router.GET("/api/users/:agentId", SendUsers)
	router.GET("/api/policies/:agentId", SendPolicies)
	http.ListenAndServe(":8080", router)
}
