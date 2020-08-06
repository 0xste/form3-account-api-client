# form3-account-client

@Author Stefano Mantini

## This project
- Implements the Create, Fetch, List (inc paging) and Delete operations on the accounts resource. 
- Has 60%+ unit test coverage
- Runs on docker-swarm
- Has 0 non-test dependencies
- Is written in go
- Uses go modules

# Further detail
- Full usage documented in main.go
- main.go runs example api requests every 10 seconds against the api

## Running

### In a swarm ensemble 
- `./build.sh` -> builds the container with the appropriate tags
- `docker-compose up -d` -> runs the ensemble

### Outside of swarm
- `./build.sh` -> builds the container with the appropriate tags
- `docker-compose up -d` -> runs the supporting services
- `./docker-run.sh` -> runs the test service, (having removed the test service from compose file) 

### Local
- `./coverage` -> runs coverage stats (excluding main.go)
- `./run.sh`