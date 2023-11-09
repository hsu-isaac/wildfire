package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct for Joke API response
type Joke struct {
	Type  string      `json:"type"`
	Value valueStruct `json:"value"`
}

type valueStruct struct {
	Categories []string `json:"categories"`
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
}

/*
getJoke (1) calls the joke API, (2) reads the response, and (3) unmarshalls it
into the Joke struct.
*/
func GetJoke(c *gin.Context) (string, error) {
	// (1) Call the joke API
	jokeUrl := "http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=firstName&lastName=lastName"
	resp, err := http.Get(jokeUrl)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// (2) Read the API response
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// (3) Unmarshall the API response into the Joke struct
	var jokeResult Joke
	err = json.Unmarshal(responseData, &jokeResult)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return jokeResult.Value.Joke, nil
}
