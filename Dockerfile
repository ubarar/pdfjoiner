FROM node

RUN apt update && apt install -y poppler-utils ghostscript

# create the directory and cd's to it
RUN mkdir /app /app/storage /app/storage/input /app/storage/output
WORKDIR /app
COPY ./package.json /app/package.json

RUN npm install

COPY . /app

EXPOSE 8080

CMD node server.js