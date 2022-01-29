FROM golang:alpine as build

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

# RUN go build -o main . 
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main . 

FROM alpine
COPY --from=build /app/main /app/main
COPY --from=build /app/frontend/ /app/frontend/
WORKDIR /app

EXPOSE 9100

ENTRYPOINT ["/app/main"]