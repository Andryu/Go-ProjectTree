package services

import(
    "fmt"
    "../models"
)

type CommonParams struct{
    Result string
    Status string
    Message string
}

func CreateRanking() *CommonParams{
    u := models.NewUser(0, "Player1")
    u.Validate()
    if u.Err() != nil{
        fmt.Println("Service User is valid")
        return &CommonParams{"failed", "400", u.Err().Error()}
    }
    u.Create()

    // ranking

    return &CommonParams{"success", "200", ""}
}
