package services

import (
	"encoding/json"
	"fmt"
	"io"
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

const FirstName = "John"
const LastName = "Doe"

/*
getJoke (1) calls the joke API, (2) reads the response, and (3) unmarshalls it
into the Joke struct.
*/
func GetJoke(c *gin.Context) (string, error) {
	// (1) Call the joke API
	const jokeUrl = "http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s"
	url := fmt.Sprintf(jokeUrl, FirstName, LastName)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// (2) Read the API response
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// (3) Unmarshall the API response into the Joke struct
	var jokeResult Joke
	err = json.Unmarshal(responseData, &jokeResult)
	if err != nil {
		return "", err
	}

	return jokeResult.Value.Joke, nil
}
