package handler

import (
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
func CombineNameAndJoke(c *gin.Context) {
	const API_RETRY_COUNT = 3

	var name services.Name
	var joke string
	var wg sync.WaitGroup

	wg.Add(2)

	// (1a) Get name from name API
	// Name API throws "unlucky error" occasionally, automatic retry to account for that
	go func() {
		defer wg.Done()

		for i := 0; i < API_RETRY_COUNT; i++ {
			fullName, err := services.GetName(c)
			if err != nil && i == API_RETRY_COUNT-1 {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error calling Joke API",
				})
				return
			} else {
				name = *fullName
				break
			}
		}
	}()

	// (1b) Get joke from joke API
	go func() {
		defer wg.Done()

		for i := 0; i < API_RETRY_COUNT; i++ {
			fullJoke, err := services.GetJoke(c)
			if err != nil && i == API_RETRY_COUNT-1 {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error calling Name API",
				})
				return
			} else {
				joke = fullJoke
				break
			}
		}
	}()

	wg.Wait()

	// (2) Replace default name in the joke with the new name
	finalJoke := strings.Replace(joke, "firstName", name.FirstName, 1)
	finalJoke = strings.Replace(finalJoke, "lastName", name.LastName, 1)

	// Returning string instead of JSON per task example
	// (3) Write the full joke
	c.String(http.StatusOK, finalJoke)
}
