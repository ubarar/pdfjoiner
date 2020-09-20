# pdfjoiner

This is an absolutely bare bones web-service for uploading multiple pdf files and joining them into one. No considerations for multiple concurrent users, large files, persistent storage etc. have been made.

### How to build

Anyone can make this app by running `make image` in the project root, given that they have docker installed on their system.

### How to run in production

In order to simplify the deployment of this app, I chose not to use reverse proxying and designed it to have TLS certs baked right into the image. With minor modifications, namely changing some of the final lines of server.js, this app will run on plain HTTP, and will be ready to be placed by a proxy or load balancer of your choice.

If you _do_ have a `key.pem` and a `cert.pem`, you can simply place them in the main project directory prior to build the image. The resulting image then, will be ready to run in production from a public port.

### Contribution

Any contributions are welcome!