package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Perms default file permissions
const Perms = 0777

// CheckArgsLength checks the length of path elements separated by '/'
func CheckArgsLength(args []string, i int, w http.ResponseWriter) bool {
	l := len(args)

	for _, s := range args {
		if s == "" {
			l--
			break
		}
	}

	if l < i {
		fmt.Fprintln(w, "insufficient args")
		return false
	}

	return true
}

// MsgDir message directory
func MsgDir(usr string) string {
	ret, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/gomsg/msg/%s/", ret, usr)
}

// UsrDir users directory
func UsrDir(usr string) string {
	ret, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/gomsg/usr/%s.dat", ret, usr)
}

// SplitPath splits a path
func SplitPath(path string) []string {
	return strings.Split(path[1:], "/")[1:]
}

// ErrorText returns an error text
func ErrorText(text string, err error) string {
	var s string
	if err != nil {
		s = ": " + err.Error()
	}
	return text + s
}

// AlertSuccessful alerts when successful
func AlertSuccessful(w http.ResponseWriter) {
	fmt.Fprintln(w, "operation successful")
}

// IsPathError returns whether an error is a *fs.PathError
func IsPathError(err error) bool {
	return reflect.TypeOf(err).String() == "*fs.PathError"
}
