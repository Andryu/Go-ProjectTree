package models

import(
    "fmt"
    "errors"
)

type User struct{
    Id int
    Name string
    err error
}

var (
    ErrIdInIsvalid = errors.New("User Id is Invalid")
    ErrIdIsLessThan = errors.New("User Id is less than 1")
    ErrNameIsInvalid = errors.New("User Name is Invalid")
)

func NewUser(id int, name string) *User{ // params
    return &User{ id, name, nil }
}

func (u *User) Err() error{
    return u.err
}

func (u *User) Create(){
    if u.err != nil{
        return
    }
    fmt.Println(u.Id)
    fmt.Println("User Create !!")
}

func (u *User) Validate(){
    if u.err != nil{
        return
    }
    if u.Id >= 1{
        u.err = ErrIdIsLessThan
        return
    }
    if u.Name == "" {
        u.err = ErrNameIsInvalid
        return
    }
}
