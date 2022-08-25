# Golang-Postgres-Docker

This is a simple Go program with a database connection and a server with 2 endpoints.

But this program we can completely run by using docker containers.

In this I have created dockerfile and docker-compose file.


For running this program in docker, run the following the comment:
   
      command : docker-compose up -d --build 
   
Which means:

    up : it is for creating and running containers.
    -d : it is for running container in detach mode.
    --build: it is for build or rebuild services.
    
For seeng logs of a particular container:
    
    command : docker logs -f <container_name>
    
For deleting containers:
    
    command : docker-compose down
