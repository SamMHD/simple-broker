FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o all -ldflags="-X 'main.TargetService=all'" main.go
RUN mv all bin/

# Run stage (will remove source code in order to reduce image size)
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/bin/all .
COPY app.env .

# won't expose anything because the whole project is a test project
# in a real project, you should expose the port you want to use
# EXPOSE 8080

CMD [ "" ]
ENTRYPOINT [ "/app/all" ]