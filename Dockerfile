FROM golang:1.15-rc-alpine
COPY f3-client-validation f3-client-validation
RUN ls -ltra f3-client-validation
ENTRYPOINT ./f3-client-validation