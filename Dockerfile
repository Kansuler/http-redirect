# Build server stage
FROM golang:1.15 as Builder
WORKDIR /server
ENV CGO_ENABLED 0
ENV GOOS linux
COPY . .
RUN go build -a -installsuffix cgo -o binary ./cmd

# Production stage
FROM scratch
COPY --from=Builder /server/binary ./
EXPOSE 8080
ENTRYPOINT ["./binary"]
