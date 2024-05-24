FROM --platform=linux/amd64 ubuntu:latest
LABEL authors="keirwhitlock"

RUN apt-get update -y
RUN apt-get install make wget -y

WORKDIR /usr/local/src
COPY . .

RUN make setup

ENTRYPOINT ["/usr/local/go/bin/go", "version"]