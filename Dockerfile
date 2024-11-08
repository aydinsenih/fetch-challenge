FROM golang:1.22.5-alpine3.20

WORKDIR /fetch

COPY . .

EXPOSE 1323

CMD ["go", "run", "."]