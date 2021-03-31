# TODO: multiarch
FROM ubuntu
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -y -qq update && \
    apt-get install -y -qq wget jq curl git gnupg software-properties-common && \
    apt-get clean

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-key C99B11DEB97541F0 && \
    apt-add-repository https://cli.github.com/packages && \
    apt update && \
    apt install gh && apt-get clean

RUN wget "https://github.com/benjvi/PRify/releases/download/0.0.2/prify-linux-amd64" -O prify && install prify /usr/local/bin/ && rm prify
