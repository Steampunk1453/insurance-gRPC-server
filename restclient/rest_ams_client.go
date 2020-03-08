package restclient

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	amsApiUrl = flag.String("domain", "http://localhost:8080/", "Domain, including protocol")
)

func GetUsers(agentID int64) []byte {
	response, err :=  http.Get(*amsApiUrl + "api/users/" + strconv.Itoa(int(agentID)))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

func GetPolicies(agentID int64) []byte {
	response, err := http.Get(*amsApiUrl + "api/policies/" + strconv.Itoa(int(agentID)))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}
