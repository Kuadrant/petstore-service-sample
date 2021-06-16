FROM golang:1.16-alpine AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN echo "Configuring go tools" \
  && apk add --no-cache git bash openssl

WORKDIR /opt/petstore/
COPY go.mod ./
RUN echo "Pulling code dependencies" \
  && go mod download
COPY . .
RUN go build -o petstore -v main.go

# Final image
FROM scratch
COPY --from=build /opt/petstore/petstore /petstore
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/petstore"]
