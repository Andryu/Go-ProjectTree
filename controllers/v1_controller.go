package controllers
import (
    "fmt"
    "../services"
)

func RankingUser() interface{}{
    fmt.Println("Services....")
    return services.CreateRanking()
}
