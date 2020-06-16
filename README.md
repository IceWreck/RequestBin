# RequestBin

RequestBin is like httpbin, but it logs requests so that you can see them later. Might be shoddy work, and probably noone else will use this, but I built this to replace my PHP based soln for the same and I wanted to play with Go.

* Single endpoint for accepting all kinds of requests.
* Sends back everything it receives as response. (except multipart files for now)
* Saves request logs so that they can be seen later.
* Basic http authentication to view logs.
* Usecase: Test API clients, common endpoint to send random stuff.
* Lean (14 MB binary, uses < 2 MB RAM>)

## Instructions
* Compile a binary. `go build`
* Upload to your server.
* Create a `.env` file similar to the provided one
* Make sure the binary has write access to the directory where its placed.
* Create a systemd service (or systemd user service) to autostart.
* You don't need httpd/nginx to serve it. RequestBin has Go's highly performant http server built in. It just listens at your specified port. You may use them a reverse proxy.
![requestbin](./requestbin.png)
