FROM golang:1.22.0-alpine3.19
WORKDIR /root
COPY .. /root
RUN sh build.sh
CMD ["sh", "output/bootstrap.sh"]