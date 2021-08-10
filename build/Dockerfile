ARG RepositoryUri
FROM ${RepositoryUri}golang:1.14.2-alpine3.11 as builder
WORKDIR /app
COPY . .
RUN go build -o main

FROM ${RepositoryUri}alpine:3.12.0
ARG version
ENV version=$version
LABEL maintainer="ting911111@gmai.com"
RUN apk add --no-cache tzdata
COPY --from=builder /app/build/entrypoint.sh /opt/entrypoint.sh
COPY --from=builder /app/main /opt
COPY --from=builder /app/public /opt/public

WORKDIR /opt
ENTRYPOINT ["/opt/entrypoint.sh"]
CMD ./main
