FROM golang:1.16

WORKDIR /sdu_student_grade_tracker

COPY sdu_student_grade_tracker .

RUN go get github.com/redis/go-redis/v9

RUN go run main.go
