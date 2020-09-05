## Intial project setupt or scalable TCP/IP websockets with Redis and round robin load balancing

### To run this project use:-
1 ) `docker build -t app .` then
2 ) `docker-compose up`

### To stop all containers use :-
`docker-compose down`

### To kill and remove all docker containers do the followig 
### stop all containers:
`docker kill $(docker ps -q)`

### remove all containers:
`docker rm $(docker ps -a -q)`
