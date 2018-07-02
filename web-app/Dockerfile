FROM golang:latest
RUN mkdir -p /go/src/web-app
WORKDIR /go/src/web-app
COPY . /go/src/web-app
RUN go-wrapper download
RUN go-wrapper install
ENV PORT 8080
EXPOSE 8080
CMD ["go-wrapper", "run"]
