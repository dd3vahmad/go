package main

import (
	"fmt"
	"math"
)

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

type authInfo struct {
	username string 
	password string
}

type shape interface {
	area() float64
	parameter() float64
}

func (i authInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", i.username, i.password)
}

func (r rect) area() float64 {
	return (r.width * r.height)
}

func (r rect) perimeter() float64 {
	return (2*r.height + 2*r.width)
}

func (c circle) area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return (math.Pi * c.radius * 2)
}

// Type Assertions - Interfaces
// "c" is a new circle cast from "s"
// which is an instance of a shape
// "ok is bool whihc is true if s was a circle"
// or false if it isn't a circle
// c, ok := s.(circle) 								----> Main syntax

// func main() {
// 	r := rect{
// 		height: 4,
// 		width: 5,
// 	}

// 	c := circle{
// 		radius: 3,
// 	}

// 	auth := authInfo{
// 		username: "theahmad",
// 		password: "Th3G0aT",
// 	}

// 	fmt.Printf("Rect Area: %.2f, Parameter: %.2f | Auth String => %s\n", r.area(), r.perimeter(), auth.getBasicAuth())
// 	fmt.Printf("Circle Area: %.2f, Parameter: %.2f | Auth String => %s", c.area(), c.perimeter(), auth.getBasicAuth())
// }