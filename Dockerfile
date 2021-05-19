FROM golang

RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build

# TODO: multiarch
FROM ubuntu
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -y -qq update && \
    apt-get install -y -qq wget jq curl git gnupg software-properties-common rsync && \
    apt-get clean

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-key C99B11DEB97541F0 && \
    apt-add-repository https://cli.github.com/packages && \
    apt update && \
    apt install gh && apt-get clean

RUN wget "https://storage.googleapis.com/kubernetes-release/release/$(wget -O- https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" && install kubectl /usr/local/bin/
RUN wget "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv3.8.8/kustomize_v3.8.8_linux_amd64.tar.gz" -O k.tar.gz && tar -xvzf k.tar.gz && install kustomize /usr/local/bin && rm k.tar.gz

RUN wget "https://github.com/benjvi/yshard/releases/download/0.0.1/yshard-linux-amd64" -O yshard && install yshard /usr/local/bin

COPY --from=0 --chmod=777 /build/prify /usr/local/bin/prify
