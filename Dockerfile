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

# 画像ファイルを取得し、ファイル数分のjpgを作成
RUN mkdir -p /root/input && \
    for file in /root/input/*; do \
        ./lossyloop "$file" "${file%.*}.jpg"; \
    done

CMD ["lossyloop 90"]