package main

import (
    "fmt"
    //"context"
    "github.com/redis/go-redis/v9"
)



type Student struct {
  Name     string
  ID       int
  Courses  []*Course
  redisCli *redis.Client
}

type Course struct {
  Code  string
  Grade float64
}

func (s *Student) AddCourse(c *Course) error {
  key := fmt.Sprintf("student:%d:course:%s", s.ID, c.Code)
  err := s.redisCli.Set(key, c.Grade, 0).Err()
  if err != nil {
    return err
  }
  s.Courses = append(s.Courses, c)
  return nil
}


func (s *Student) ViewCourses() ([]*Course, error) {
  courses := make([]*Course, 0, len(s.Courses))
  for _, c := range s.Courses {
    key := fmt.Sprintf("student:%d:course:%s", s.ID, c.Code)
    grade, err := s.redisCli.Get(key).Float64()
    if err != nil {
      return nil, err
    }
    c.Grade = grade
    courses = append(courses, c)
  }
  return courses, nil
}

func (s *Student) UpdateCourse(c *Course) error {
  key := fmt.Sprintf("student:%d:course:%s", s.ID, c.Code)
  err := s.redisCli.Set(key, c.Grade, 0).Err()
  if err != nil {
    return err
  }
  for i, oldCourse := range s.Courses {
    if oldCourse.Code == c.Code {
      s.Courses[i] = c
      break
    }
  }
  return nil
}

func (s *Student) DeleteCourse(code string) error {
  key := fmt.Sprintf("student:%d:course:%s", s.ID, code)
  err := s.redisCli.Del(key).Err()
  if err != nil {
    return err
  }
  for i, c := range s.Courses {
    if c.Code == code {
      s.Courses = append(s.Courses[:i], s.Courses[i+1:]...)
      break
    }
  }
  return nil
}

func (c *Course) GetCode() string {
  return c.Code
}
func (c *Course) GetGrade() float64 {
  return c.Grade
}


func main() {
  redisCli := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
  })
defer redisCli.Close()

  

  
  pong, err := redisCli.Ping().Result()
  if err != nil {
    log.Fatalf("Failed to ping Redis: %v", err)
  }
  log.Printf("Connected to Redis: %s", pong)

  
  student := &Student{
    Name:     "Yerassyl",
    ID:       123,
    Courses:  make([]*Course, 0),
    redisCli: redisCli,
  }

  
  err = student.AddCourse(&Course{Code: "CS101", Grade: 90.5})
  if err != nil {
    log.Fatalf("Failed to add course: %v", err)
  }
  err = student.AddCourse(&Course{Code: "MATH101", Grade: 85.0})
  if err != nil {
    log.Fatalf("Failed to add course: %v", err)
  }

  courses, err := student.ViewCourses()
  if err != nil {
    log.Fatalf("Failed to view courses: %v",
  )
}
}
