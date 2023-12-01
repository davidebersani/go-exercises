package main

import "fmt"

type authenticationInfo struct {
	Username string
	Password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.Username, authInfo.Password)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("============================================")
}

func main() {
	test(authenticationInfo{
		Username: "davide",
		Password: "ds5a61d56sa",
	})
}
