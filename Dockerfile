FROM alpine:3.8
LABEL maintainer="John McKenzie<jmckind@gmail.com>"

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.2/dumb-init_1.2.2_amd64
RUN chmod +x /usr/local/bin/dumb-init

COPY build /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]
CMD ["/usr/local/bin/rechat"]

EXPOSE 8000/tcp
