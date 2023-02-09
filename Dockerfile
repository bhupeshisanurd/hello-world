FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

ENV NAME "default_name"
ENV JOB "default_job"
ENV FAV_ANIMAL "default_animal"

CMD ["./main"]
