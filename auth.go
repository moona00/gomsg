package main

import (
	"fmt"
	"net/http"
	"os"
)

// UserExists checks whether an user exists
func UserExists(usr string, w http.ResponseWriter) bool {
	_, err := os.ReadFile(UsrDir(usr))

	if err == nil {
		return true
	} else if IsPathError(err) {
		fmt.Fprintf(w, "user %s does not exist\n", usr)
		return false
	} else {
		fmt.Fprintln(w, ErrorText("read error", err))
		return false
	}
}

// IsPasswordCorrect checks whether the pass is correct
func IsPasswordCorrect(usr, pwd string, w http.ResponseWriter) bool {
	data, err := os.ReadFile(UsrDir(usr))

	if err != nil {
		if IsPathError(err) {
			fmt.Fprintln(w, "user does not exist")
		} else {
			fmt.Fprintln(w, ErrorText("read error", err))
		}
		return false
	} else if string(data) != pwd {
		fmt.Fprintln(w, "incorrect password")
		return false
	}

	return true
}
