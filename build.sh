GOOS=linux CGO_ENABLED=0 go build -tags musl -o form3-account-client-test
docker build --no-cache -t stefanomantini/form3-account-client-test .