package models

type User struct {
    ID    int
    Name  string
    Email string
}

func (u *User) GetFullName() string {
    return u.Name
}
