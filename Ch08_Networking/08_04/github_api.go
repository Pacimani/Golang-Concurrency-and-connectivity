// Calling Github API

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"` // for public repos in JSON
}

func UserInfo(login string) (*User, error) {

	u := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	user := User{Login: login}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {

	user, err := UserInfo("Pacimani")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", user)
}
