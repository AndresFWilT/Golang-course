package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

func main() {
	lc := loginClient(url+"/v1/login", "contacto@ed.team", "123456")
	// fmt.Println(lc)

	person := Person{
		Name:        "Andres",
		Age:         26,
		Communities: []Community{Community{Name: "EDteam"}, Community{Name: "Golang en español"}},
	}
	gr := createPerson(url+"/v1/persons", lc.Data.Token, &person)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}

	return response
}

// Login .
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GeneralResponse .
type GeneralResponse struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

// LoginResponse .
type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// Community .
type Community struct {
	Name string `json:"name"`
}

// Communities slice de comunidades
type Communities []Community

// Person .
type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error en marshal de login: %v", err)
	}

	resp := httpClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error en lectura del body en login: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Se esperaba código 200, se obtuvo: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body en login: %v", err)
	}

	return dataResponse
}

func createPerson(url, token string, person *Person) GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error en marshal de persona: %v", err)
	}

	resp := httpClient(http.MethodPost, url, token, data)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error en lectura del body en login: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Se esperaba código 201, se obtuvo: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body en login: %v", err)
	}

	return dataResponse
}
