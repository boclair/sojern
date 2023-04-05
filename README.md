# sojern
Take home assignment for sojern problems 1 (semantic version) and 3 (/ping and /img webserver)

`$ go test` to run the unit tests for the semantic version code and the web server code

`$ go build` to build

`$ ./web_server_and_version` to run
* Starts a webserver on `localhost:8080`

Once the web server is running, then use these commands to test it:
* `$ curl -v http://localhost:8080/ping`
* `$ curl -v http://localhost:8080/img`
* `$ touch /tmp/ok` to create a the file that impacts `/ping`
* `$ rm /tmp/ok` to remove...
