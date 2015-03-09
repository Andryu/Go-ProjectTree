package main

import (
  "./config"
  "./controllers"
  "fmt"
)

var (
  pool = "POOL......" + config.GOENV
)

func main() {
  fmt.Println(pool)
  fmt.Println("Call Main")

  res := controllers.RankingUser()
  fmt.Println(res)
  fmt.Println("End Main...")
}
