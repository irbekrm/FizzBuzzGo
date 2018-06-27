package main

import (
  "math"
  "strconv"
)

type Describer interface {
    Describe() string
    NotZero() bool
    IsFactor(float64) bool
    ToString() string
}

type ChosenNumber float64

func (c ChosenNumber) Describe() (string) {
  if c.NotZero() && c.IsFactor(15) {
    return "FizzBuzz"
  }
  if c.NotZero() && c.IsFactor(5) {
    return "Buzz"
  }
  if c.NotZero() && c.IsFactor(3) {
    return "Fizz"
  }
  return c.ToString()
}

func (c ChosenNumber) IsFactor(value float64) bool {
  number := float64(c)
  remainder := math.Remainder(number, value)
  return remainder == 0
}

func (c ChosenNumber) NotZero() bool {
  return c != 0
}

func (c ChosenNumber) ToString() string {
  f := float64(c)
  return strconv.FormatFloat(f, 'f', -1, 64)
}
