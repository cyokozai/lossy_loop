###################### BUILDER ######################
FROM golang:latest

SHELL ["/bin/bash", "-c"]

ARG lang="C"
ARG dir="src"

ENV DEBIAN_FRONTEND noninter active
ENV TERM xterm
ENV DISPLAY host.docker.internal:0.0
ENV LANG ${lang}
ENV LANGUAGE ${lang}
ENV LC_ALL ${lang}
ENV TZ Asia/Tokyo

WORKDIR /root/${dir}

COPY ./${dir}/*.go /root/${dir}

RUN apt -y update && apt -y upgrade && go mod download && go mod verify && go build -v -o lossyloop

CMD ["lossyloop 90"]