package config

import (
	"fmt"
	"os"
)

type UserContext struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadFromEnv() *UserContext {
	// Load username & password from os.env
	username := os.Getenv("IMAGINE_USERNAME")
	password := os.Getenv("IMAGINE_PASSWORD")

	if username == "" || password == "" {
		fmt.Println("Please use environment variables [IMAGINE_USERNAME, IMAGINE_PASSWORD] to set username and password ")
		os.Exit(1)
	}

	return &UserContext{
		Username: username,
		Password: password,
	}
}
