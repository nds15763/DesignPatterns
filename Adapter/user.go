package Adapter

import "fmt"

type User struct{
	file File
}

func Init(file File) *User{
	return &User{file}
}

func (u *User)UserReadBook(){
	fmt.Println("UserReadBook")
	u.file.CsvRead()
}