FROM alpine:3.11.3

WORKDIR /work

ARG github_token

RUN set -x && \
    apk update && \
    apk --no-cache add git wget musl-dev ca-certificates && \
    mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz && \
    tar -C /usr/local/ -xzf go1.14.linux-amd64.tar.gz && \
    rm -f go1.14.linux-amd64.tar.gz && \
    git clone https://$github_token:x-oauth-basic@github.com/[[repo_url]] . && \
    /usr/local/go/bin/go build -o ./main && \
    mkdir mount

CMD ["/work/main"]
