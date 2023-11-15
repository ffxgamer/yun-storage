FROM golang:1.20 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/main ./...


FROM xhofe/alist:latest

WORKDIR /opt/alist/

COPY --from=builder --chmod=0755 /usr/local/bin/main .

ENV DB_TYPE postgres
ENV DB_PORT 5432
ENV DB_TABLE_PREFIX alist_
# prefer require disable
ENV DB_SSL_MODE prefer
ENV PORT 5244

# ENV SITE_URL
# ENV DB_HOST tiny.db.elephantsql.com
# ENV DB_NAME 
# ENV DB_USER
# ENV CDN

COPY --chmod=0755 entrypoint.sh /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]