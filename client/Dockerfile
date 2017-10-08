# FROM golang
FROM golang:1.9.0

# Set timezone to Asia/Tokyo
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apt-get update \
  && apt-get install -y build-essential git curl wget \
                        zlib1g-dev libxml2-dev libxslt1-dev \
                        openssl less libssl-dev libreadline-dev vim

# Install glide
RUN curl https://glide.sh/get | sh

# Add the local package files to the container's workspace.
COPY ./ /go/src/
COPY ./ /go/src/client/

WORKDIR /go/src/client

# Install fresh for live reloading
RUN go get github.com/pilu/fresh

# Install packages via glide
RUN glide install
