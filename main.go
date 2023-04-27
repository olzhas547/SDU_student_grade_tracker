package main

import (
    "fmt"
    "context"
    "github.com/redis/go-redis/v9")


var ctx = context.Background()


type Student {
    ID int
    name string
    courses []Course
}

type Course {
    ID int
    grade float64
}

func (s Student) addCourse {
    return 0
    if err != nil {
        panic(err)
    }
}

func (s Student) viewCourse {
    return 0
    if err != nil {
        panic(err)
    }
}

func (s Student) updateCourse {
    return 0
    if err != nil {
        panic(err)
    }
}
func (s Student) deleteCourse {
    return 0
    if err != nil {
        panic(err)
    }
}

func (c Course) getCourceCode {
    return 0
    if err != nil {
        panic(err)
    }
}

func (c Course) getStudentGrade {
    return 0
    if err != nil {
        panic(err)
    }
}
