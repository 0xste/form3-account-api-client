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

### Suggested Improvements
- Remove main.go! Replace with [godog](https://github.com/cucumber/godog) for integration tests
- Move package structure to a more [domain-driven design](https://youtu.be/MzTcsI6tn-0) (all models and interfaces defined in account.go) and separate files implementing behaviours, apprehensive as there's some custom types implemented at the top level
- Run tests, fmt, vet staticcheck in pipeline
- Custom type & validator for bankIdCode
- Move custom types into separate modules with their own `go.mod` files respectively for reuse across different packages
- Refactor countries type to use a struct to handle more attributes (the current setup only supports storing 2 digit country codes against the country)
- Run container as non-root & convert to using base alpine
- Upgrade to go v1.15 (works local, but had some issues when running on alpine)