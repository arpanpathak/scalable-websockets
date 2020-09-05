## To kill and remove all docker containers do the followig 
## stop all containers:
docker kill $(docker ps -q)

## remove all containers:
docker rm $(docker ps -a -q)