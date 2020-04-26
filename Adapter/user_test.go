package Adapter

import "testing"

func TestUser_UserReadBook(t *testing.T) {
	file := File{ftype: "Csv"}
	user := User{file}
	user.UserReadBook()
}
