package marvel

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const BaseURL = "https://gateway.marvel.com:443/v1/public/"

type Client struct {
	baseURL    string
	publicKey  string
	privateKey string
	httpClient *http.Client
}

type CharHTTPResponse struct {
	Data struct {
		Offset  int         `json:"offset"`
		Limit   int         `json:"limit"`
		Total   int         `json:"total"`
		Count   int         `json:"count"`
		Results []Character `json:"results"`
	} `json:"data"`
}

type Character struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func marvelKeys() (string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	pub := os.Getenv("MARVEL_PUBLIC_KEY")
	priv := os.Getenv("MARVEL_PRIVATE_KEY")
	return pub, priv
}
