package main

import (
	"encoding/json"
	"net/http"
	"os"

	lab "finis/msg"
)

func main() {
	Webservices()

}
func SendMsg(w http.ResponseWriter, r *http.Request) {
	m := lab.MSG{}
	json.NewDecoder(r.Body).Decode(&m)
	m.SetKey(os.Getenv("KEY1"))
	m.Decrypted()
	json.NewEncoder(w).Encode(m)
}

func GetMsg(w http.ResponseWriter, r *http.Request) {
	m := lab.MSG{}
	m.SetMSG(os.Getenv("MSG"))
	m.SetKey(os.Getenv("KEY2"))
	m.Encrypted()
	json.NewEncoder(w).Encode(m)
}

func Webservices() {
	http.HandleFunc("/SendMsg", SendMsg)
	http.HandleFunc("/GetMsg", GetMsg)
	port := os.Getenv("PORT")
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
