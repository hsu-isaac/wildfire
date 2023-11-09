package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Constant for API retry count
const API_RETRY_COUNT = 3

// Struct for Name API response
type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

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
getName (1) calls the name API, (2) reads the response, and (3) unmarshalls it
into the Name struct.
*/
func getName(w http.ResponseWriter, r *http.Request) (*Name, error) {
	// (1) Call the name API
	resp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// (2) Read the API response
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// (3) Unmarshall the API response into the Name struct
	var nameResult Name
	err = json.Unmarshal(responseData, &nameResult)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &nameResult, nil
}

/*
getJoke (1) calls the joke API, (2) reads the response, and (3) unmarshalls it
into the Joke struct.
*/
func getJoke(w http.ResponseWriter, r *http.Request) (string, error) {
	// (1) Call the joke API
	resp, err := http.Get("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=firstName&lastName=lastName")
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

/*
handler (1a) get name from name API, (1b) get joke from joke API, (2) replace
default name in the joke with the new name, and (3) write the full joke.
*/
func handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(2)

	var name Name
	var joke string
	// (1a) Get name from name API
	// Name API throws "unlucky error" occasionally, automatic retry to account for that
	go func() {
		defer wg.Done()
		for i := 0; i < API_RETRY_COUNT; i++ {
			fullName, err := getName(w, r)
			if err != nil && i == API_RETRY_COUNT-1 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err == nil {
				name = *fullName
				break
			}
		}
	}()

	// (1b) Get joke from joke API
	go func() {
		defer wg.Done()
		for i := 0; i < API_RETRY_COUNT; i++ {
			fullJoke, err := getJoke(w, r)
			if err != nil && i == API_RETRY_COUNT-1 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err == nil {
				joke = fullJoke
				break
			}
		}
	}()

	wg.Wait()

	// (2) Replace default name in the joke with the new name
	finalJoke := strings.Replace(joke, "firstName", name.FirstName, 1)
	finalJoke = strings.Replace(finalJoke, "lastName", name.LastName, 1)

	// (3) Write full joke
	w.Write([]byte(finalJoke))
}

func handleRequests() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
