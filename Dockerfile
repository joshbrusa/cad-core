FROM golang AS build

WORKDIR /code

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/httpServer ./cmd/httpServer/main.go

FROM alpine

RUN apk --no-cache add curl

COPY --from=build /code/bin/httpServer .

CMD [ "/httpServer" ]