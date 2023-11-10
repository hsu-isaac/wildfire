package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct for Name API response
type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

/*
getName (1) calls the name API, (2) reads the response, and (3) unmarshalls it
into the Name struct.
*/
func GetName(c *gin.Context, errChan chan error) (*Name, error) {
	// (1) Call the name API
	const nameUrl = "http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe"
	resp, err := http.Get(nameUrl)
	if err != nil {
		return nil, err
	}

	// (2) Read the API response
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// (3) Unmarshall the API response into the Name struct
	var nameResult Name
	err = json.Unmarshal(responseData, &nameResult)
	if err != nil {
		return nil, err
	}

	return &nameResult, nil
}
