FROM golang AS build

WORKDIR /code

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/server/main.go

FROM alpine

RUN apk --no-cache add curl

COPY --from=build /code/bin/server .

CMD [ "/server" ]