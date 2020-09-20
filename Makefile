
image:
	docker build -t pdfjoiner:latest .

run:
	docker run -p 8080:8080 pdfjoiner:latest

run-interactive:
	docker run -p 8080:8080  --rm -it --entrypoint /bin/bash pdfjoiner:latest

run-production:
	docker run -p 443:8080 --restart unless-stopped pdfjoiner:latest
