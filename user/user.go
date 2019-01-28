package user

import (
	"fmt"
	"os/user"
)

// GetUser bb
func GetUser() {
	usr, err := user.Current()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", usr.Uid, usr.Gid, usr.Username, usr.Name, usr.HomeDir)
}
