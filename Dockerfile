FROM golang:1.14

# Set necessary environmet variables needed for our image
ENV APPPORT=8080\
    REDIS_HOST=redis\
    REDIS_PORT=6379


WORKDIR /home/app/dist
COPY app /home/app/dist
RUN go mod download
RUN go build
CMD ./app