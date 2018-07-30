package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(gundar)
	http.Handle("/", http.FileServer(http.Dir("front-end")))
	http.HandleFunc("/api/mahasiswa", user)
	http.ListenAndServe(":8085", nil)

}

var gundar = []mahasiswa{
	{
		Nama: "aria",
		Npm:  "10115962",
	},
	{
		Nama: "lala",
		Npm:  "1021212",
	},
	{
		Nama: "lala",
		Npm:  "1021212",
	},
}

type mahasiswa struct {
	Nama string `json:"nama"`
	Npm  string `json:"npm"`
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {

		result, err := json.Marshal(gundar)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}


