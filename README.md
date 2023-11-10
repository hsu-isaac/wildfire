# The Task
Create a production ready web service which combines two existing web services.
Fetch a random name from https://names.mcquay.me/api/v0
Fetch a random Chuck Norris joke from http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe
Combine the results and return them to the user.

# Time Guidelines
We recommend you spend around 2 to 4 hours on this task.
You should aim to at least have running code which meets the basic requirements of the task.
Production ready is a broad goal: if you’re unable to fully meet it then please include clear 
TODO comments in your code that are sufficiently detailed that another engineer could complete 
the task without having to do additional design thinking and the result would match your vision of production readiness.
Please let us know how much time you spent on the task when you submit your answer.
## Further requirements:
The web service should be written in Go.
Write a README.md file which (at a minimum) provides instructions for running the web service.
The web service should remain responsive under load and be able to support multiple concurrent requests.
The code, README.md and any other supporting files should be compressed into a single archive and submitted for review or a git repo link.


# Example
## Fetching a name
```
$ curl "https://names.mcquay.me/api/v0/"
{“first_name”:“Hasina”,“last_name”:“Tanweer”}
Fetching a joke
$ curl "http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe"
{ “type”: “success”, “value”: { “id”: 181, “joke”: “John Doe’s OSI network model has only one layer - Physical.“, “categories”: [“nerdy”] } }
```

## Using the new web service
```
$ curl "http://localhost:5000"
Hasina Tanweer’s OSI network model has only one layer - Physical..
```

# The Solution
## How to Run
- Install Go if not already installed locally
- Clone into this repository `git clone https://github.com/hsu-isaac/wildfire.git`
- Navigate into the project `cd wildfire`
- Run service with `go run .`
- Hit service with curl command `curl localhost:8080/` OR visit `localhost:8080/` on your browser

## TODO
I believe that this solution manages to accomplish the basic requirements of the task,
however it is not yet production ready. With more time I would like to accomplish the
following:
- Write more extensive tests for code
  - More extensive test cases in handler_test.go
  - Mock API responses and ensure correct handling of errors
- Check up on error handling for the APIs and what limitations they might have
  - Rate limiting, random errors, etc
- Production level code would likely also involve setting up an environment so that
  URLs and other constants don't have to be explicitly stated





