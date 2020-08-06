FROM alpine:3.7
COPY form3-account-client-test /form3-account-client-test
RUN ls -ltra
CMD ./form3-account-client-test