FROM golang:1.10-alpine as builder

ENV REPO_NAME=github.com/sjuls/soup-ranking

RUN apk update \
  && apk add curl \
  && apk add git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p $GOPATH/src/$REPO_NAME
COPY . $GOPATH/src/$REPO_NAME/
WORKDIR $GOPATH/src/$REPO_NAME

RUN mkdir -p /out
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /out/soup-ranking .

FROM scratch as soup-ranking

ENV PORT=8080
ENV DATABASE_URL=

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /out/soup-ranking /

CMD [ "/soup-ranking" ]
