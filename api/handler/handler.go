package handler

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"wildfire/api/services"

	"github.com/gin-gonic/gin"
)

/*
handler (1a) get name from name API, (1b) get joke from joke API, (2) replace
default name in the joke with the new name, and (3) write the full joke.
*/
func FetchNameAndJokeHandler(c *gin.Context) {
	const API_RETRY_COUNT = 3

	var name services.Name
	var joke string
	var wg sync.WaitGroup

	errChan := make(chan error, 2)
	defer close(errChan)

	// (1a) Get name from name API
	// Name API throws "unlucky error" occasionally, automatic retry to account for that
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < API_RETRY_COUNT; i++ {
			fullName, err := services.GetName(c, errChan)
			if err != nil && i == API_RETRY_COUNT-1 {
				errChan <- err
				return
			} else if err == nil {
				name = *fullName
				break
			}
		}
	}()

	// (1b) Get joke from joke API
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < API_RETRY_COUNT; i++ {
			fullJoke, err := services.GetJoke(c, errChan)
			if err != nil && i == API_RETRY_COUNT-1 {
				errChan <- err
				return
			} else if err == nil {
				joke = fullJoke
				break
			}
		}
	}()

	wg.Wait()

	select {
	case err := <-errChan:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error occured: %v", err),
		})
		return
	default:
		// No errors, continue
	}

	// (2) Replace default name in the joke with the new name
	finalJoke, err := combineNameAndJoke(joke, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error occured: %v", err),
		})
	}

	// Returning string instead of JSON per task example
	// (3) Write the full joke
	c.String(http.StatusOK, finalJoke)
}

func combineNameAndJoke(joke string, name services.Name) (string, error) {
	if !strings.Contains(joke, services.FirstName) || !strings.Contains(joke, services.LastName) {
		return "", jokeError
	}

	finalJoke := strings.Replace(joke, services.FirstName, name.FirstName, 1)
	finalJoke = strings.Replace(finalJoke, services.LastName, name.LastName, 1)

	return finalJoke, nil
}

var jokeError error = fmt.Errorf("Joke does not contain firstName: %s, or lastName: %s", services.FirstName, services.LastName)
