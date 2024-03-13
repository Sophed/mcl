package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Profile struct {
	UUID string `json:"id"`
	Name string `json:"name"`
}

const PROFILE_ENDPOINT = "https://api.mojang.com/users/profiles/minecraft/"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		println("Usage: mcl <username>")
		os.Exit(1)
		return
	}
	username := args[0]

	profile := getProfile(username)
	if profile == nil || profile.UUID == "" {
		println("Profile not found")
		os.Exit(1)
		return
	}

	printInfo(profile)

}

func getProfile(query string) *Profile {

	res, err := http.Get(PROFILE_ENDPOINT + query)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var profile Profile
	json.Unmarshal(body, &profile)
	return &profile

}
