FROM golang AS build

WORKDIR /code

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./src -o ../bin/cad-core

FROM alpine

COPY --from=build /code/bin/cad-core .

CMD [ "/cad-core" ]