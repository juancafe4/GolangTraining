package main

import (
  "fmt"
)

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
  switch x.(type) { // this is an assert; asserting, "x is of this type"
  case int:
    fmt.Println("int")
  case string:
    fmt.Println("string")
  case contact:
    fmt.Println("contact")
  default:
    fmt.Println("unknown")

  }
}

func main() {
  switch "Mhi" {
  case "Daniel":
    fmt.Println("Wassup Daniel")
  case "Medhi":
    fmt.Println("Wassup Medhi")
  case "Jenny":
    fmt.Println("Wassup Jenny")
  default:
    fmt.Println("Have you no friends?")
  }

  switch "Marcus" {
  case "Tim":
    fmt.Println("Wassup Tim")
  case "Jenny":
    fmt.Println("Wassup Jenny")
  case "Marcus":
    fmt.Println("Wassup Marcus")
    fallthrough
  case "Medhi":
    fmt.Println("Wassup Medhi")
    fallthrough
  case "Julian":
    fmt.Println("Wassup Julian")
  case "Sushant":
    fmt.Println("Wassup Sushant")
  }

  switch "Jenny" {
  case "Tim", "Jenny":
    fmt.Println("Wassup Tim, or, err, Jenny")
  case "Marcus", "Medhi":
    fmt.Println("Both of your names start with M")
  case "Julian", "Sushant":
    fmt.Println("Wassup Julian / Sushant")
  }

  switch {
  case len(myFriendsName) == 2:
    fmt.Println("Wassup my friend with name of length 2")
  case myFriendsName == "Tim":
    fmt.Println("Wassup Tim")
  case myFriendsName == "Jenny":
    fmt.Println("Wassup Jenny")
  case myFriendsName == "Marcus", myFriendsName == "Medhi":
    fmt.Println("Your name is either Marcus or Medhi")
  case myFriendsName == "Julian":
    fmt.Println("Wassup Julian")
  case myFriendsName == "Sushant":
    fmt.Println("Wassup Sushant")
  default:
    fmt.Println("nothing matched; this is the default")
  }

}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/