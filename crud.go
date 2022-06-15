package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	database = make(map[string]User)
)

func responseJson(res http.ResponseWriter, message []byte, httpCode int) {
	res.Header().Set("Content-type", "application-json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func main() {

	database["1"] = User{ID: "1", Name: "Eka"}

	http.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			message := []byte(`{"message": "method harus GET"}`)
			responseJson(res, message, http.StatusMethodNotAllowed)
			return
		}

		var users []User

		for _, user := range database {
			users = append(users, user)
		}

		response, err := json.Marshal(users)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		responseJson(res, response, http.StatusOK)
	})

	http.HandleFunc("/user-add", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			message := []byte(`{"message": "method harus POST"}`)
			responseJson(res, message, http.StatusMethodNotAllowed)
			return
		}

		var newUser User

		payload := req.Body    //menampung request
		defer req.Body.Close() //close setelah menampung request

		err := json.NewDecoder(payload).Decode(&newUser)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		database[newUser.ID] = newUser

		message := []byte(`{"message": "success create data"}`)
		responseJson(res, message, http.StatusCreated)
	})

	http.HandleFunc("/user-get", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			message := []byte(`{"message": "method harus GET"}`)
			responseJson(res, message, http.StatusMethodNotAllowed)
			return
		}

		id := req.URL.Query().Get("id")
		if id == "" {
			message := []byte(`{"message": "id parameter harus diisi"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		data, isAvailable := database[id]
		if !isAvailable {
			message := []byte(`{"message": "data tidak ditemukan"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(&data)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(res, message, http.StatusInternalServerError)
			return
		}

		responseJson(res, response, http.StatusOK)
	})

	http.HandleFunc("/user-delete", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "DELETE" {
			message := []byte(`{"message": "method harus DELETE"}`)
			responseJson(res, message, http.StatusMethodNotAllowed)
			return
		}

		id := req.URL.Query().Get("id")
		if id == "" {
			message := []byte(`{"message": "id parameter harus diisi"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		_, isAvailable := database[id]
		if !isAvailable {
			message := []byte(`{"message": "data tidak ditemukan"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		delete(database, id)

		message := []byte(`{"message": "berhasil hapus data"}`)
		responseJson(res, message, http.StatusOK)
	})

	http.HandleFunc("/user-update", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "PUT" {
			message := []byte(`{"message": "method harus PUT"}`)
			responseJson(res, message, http.StatusMethodNotAllowed)
			return
		}

		id := req.URL.Query().Get("id")
		if id == "" {
			message := []byte(`{"message": "id parameter harus diisi"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		data, isAvailable := database[id]
		if !isAvailable {
			message := []byte(`{"message": "data tidak ditemukan"}`)
			responseJson(res, message, http.StatusBadRequest)
			return
		}

		var newUser User

		payload := req.Body

		defer req.Body.Close()

		err := json.NewDecoder(payload).Decode(&newUser)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(res, message, http.StatusInternalServerError)
			return
		}

		data.Name = newUser.Name

		database[data.ID] = data

		response, err := json.Marshal(&data)
		if err != nil {
			message := []byte(`{"message": "error parsing data"}`)
			responseJson(res, message, http.StatusInternalServerError)
			return
		}

		responseJson(res, response, http.StatusOK)
	})

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		message := []byte(`{"message": "server is run"}`)
		responseJson(res, message, http.StatusOK)
	})

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
