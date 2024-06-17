# pdfjoiner

This is an absolutely bare bones web-service for uploading multiple pdf files and joining them into one. No considerations for multiple concurrent users, large files, persistent storage etc. have been made.

### How to build

Anyone can make this app by running `make image` in the project root, given that they have docker installed on their system.

### How to run in production

If you're putting this app behind a reverse proxy, launch it with the `--no-cert` command (as is default in the dockerfile), else launch it with a `cert.pem` and `cert.key` in the working directory in order to run TLS natively

### Contribution

Any contributions are welcome!
