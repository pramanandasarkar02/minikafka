package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



type RequestType struct{
	ProducerId int64 `json:"producerId"`
	ConsumerId int64 `json:"consumerId"`
	Payload []byte `json:"payload"`
}


func ProducerAPIHandler(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()

	var requestType RequestType
	err := json.NewDecoder(r.Body).Decode(&requestType)

	if err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	fmt.Println(requestType)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payload reached in miniKafka successfully ")
}

func main() {

	fmt.Println([]byte("Pramananda Sarkar"))
	
	http.HandleFunc("/producer", ProducerAPIHandler)
	
	log.Println("Kafka server started at http://localhost:9092")
	http.ListenAndServe(":9092", nil)
	

}
