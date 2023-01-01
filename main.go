package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
rand.Seed(time.Now().UnixNano())

eludedGuards := rand.Intn(100)

var isHeistOn bool
isHeistOn = true

if eludedGuards >= 50 {
  fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
} else {
  isHeistOn = false
  fmt.Println("Plan a better disguise next time?")
}

rand.Seed(time.Now().UnixNano())
openedVault := rand.Intn(100)


if isHeistOn = true; openedVault >= 70 {
  fmt.Println("Grab and GO!")
} else if isHeistOn {  
      isHeistOn = false
      fmt.Println("The vault can't be opened!")
  } 

  rand.Seed(time.Now().UnixNano())
  leftSafely := rand.Intn(5)

if isHeistOn {
  switch leftSafely {
    case 0:
      isHeistOn = false
      fmt.Println("Looks like you tripped an alarm... run?")

    case 1:
      isHeistOn = false
      fmt.Println("Turns out vault doors don't open from the inside...")
  
    case 2:
      isHeistOn = false
      fmt.Println("A security guard caught glimpse of one of your accomplices! Run!")

    case 3:
      isHeistOn = true
      fmt.Println("Vault has been opened by your hacking skills and the guards have been distracted by the fire alarm one of your accomplices turned on. Take the money and get out!")
    default:
      fmt.Println("Start the getaway car!")
  }
}

if isHeistOn {
  var amtStolen = 10000 + rand.Intn(1000000)

  fmt.Println("Amount stolen: $", amtStolen, "Not bad!")
}

fmt.Println("isHeistOn is currently:", isHeistOn)

}
