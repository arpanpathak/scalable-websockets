FROM golang:1.14

WORKDIR /home/app/dist
COPY app /home/app/dist
RUN go mod download
RUN go build
CMD ./app