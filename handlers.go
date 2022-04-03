package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// TestHandler is a test handler
func TestHandler(w http.ResponseWriter, r *http.Request) {
	AlertSuccessful(w)
	fmt.Fprintln(w, "test")
}

// SignUpHandler is the handler for signing up
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	info := SplitPath(r.URL.Path)
	if !CheckArgsLength(info, 2, w) {
		return
	}

	filename := UsrDir(info[0])

	if _, err := os.ReadFile(filename); !IsPathError(err) {
		fmt.Fprintf(w, ErrorText("user already exists, or read error", err))
		return
	}

	if err := os.WriteFile(filename, []byte(info[1]), Perms); err != nil {
		fmt.Fprintln(w, ErrorText("write error", err))
	} else {
		AlertSuccessful(w)
	}
}

// DeleteAccountHandler is the handler for deleting an account
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	info := SplitPath(r.URL.Path)
	if !CheckArgsLength(info, 2, w) {
		return
	}

	filename := UsrDir(info[0])

	if IsPasswordCorrect(info[0], info[1], w) {
		if err, err0 := os.Remove(filename), os.RemoveAll(MsgDir(info[0])); err != nil || err0 != nil {
			fmt.Fprintln(w, ErrorText("remove error", err))
			return
		}
		AlertSuccessful(w)
	}
}

// SendHandler is the handler for sending
func SendHandler(w http.ResponseWriter, r *http.Request) {
	info := SplitPath(r.URL.Path)
	if !CheckArgsLength(info, 4, w) ||
		!IsPasswordCorrect(info[0], info[1], w) ||
		!UserExists(info[0], w) ||
		!UserExists(info[2], w) {
		return
	}

	msg := fmt.Sprintf("FROM %s\n\n%s\n", info[0], info[3])
	var k int
	dir := MsgDir(info[2])
	filename := dir + strconv.Itoa(k) + ".dat"

	for {
		if _, err := os.ReadFile(filename); err != nil {
			break
		}
		k++
		filename = dir + strconv.Itoa(k) + ".dat"
	}

	os.Mkdir(dir, Perms)
	if err := os.WriteFile(filename, []byte(msg), Perms); err != nil {
		fmt.Fprintln(w, ErrorText("write error", err))
	} else {
		AlertSuccessful(w)
	}
}

// ReceiveHandler is the handler for receiving
func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	info := SplitPath(r.URL.Path)
	if !CheckArgsLength(info, 2, w) ||
		!IsPasswordCorrect(info[0], info[1], w) {
		return
	}

	path := MsgDir(info[0])
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(w, ErrorText("read error", err))
	}

	str := "\n"
	for _, f := range files {
		data, err := os.ReadFile(path + f.Name())
		if err != nil {
			fmt.Fprintln(w, ErrorText("read error", err))
			continue
		}
		str += string(data) + "\n"
	}

	AlertSuccessful(w)
	fmt.Fprintln(w, str)
}
