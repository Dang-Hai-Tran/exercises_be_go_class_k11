package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func saveToFile(username, password, userprofile string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// Write the data in the format: username|password|userprofile
	_, err = file.WriteString(fmt.Sprintf("%s|%s|%s\n", username, password, userprofile))
	return err
}

func saveProfileImage(_ http.ResponseWriter, req *http.Request) error {
	// Parse the multipart form for the file upload
	err := req.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		return err
	}
	// Retrieve the file
	file, handler, err := req.FormFile("imageprofile")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a directory called 'profile' if it doesn't exist
	if _, err := os.Stat("profile"); os.IsNotExist(err) {
		err = os.Mkdir("profile", 0755)
		if err != nil {
			return err
		}
	}
	// Create the file in the 'profile' directory
	dst, err := os.Create(filepath.Join("profile", handler.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	return err
}

func signin(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse form data
	err := req.ParseMultipartForm(10 << 10) // Max 10 Kb ~ 10000 chars
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	userprofile := req.FormValue("userprofile")
	fmt.Printf("Received %s|%s|%s\n", username, password, userprofile)
	// Save username|password|userprofile to a file
	err = saveToFile(username, password, userprofile)
	if err != nil {
		http.Error(w, "Error saving data", http.StatusInternalServerError)
		return
	}
	// Save the profile image
	err = saveProfileImage(w, req)
	if err != nil {
		http.Error(w, "Error saving profile image", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Signin successful, profile saved.\n")
}

func main() {
	http.HandleFunc("/signin", signin)
	fmt.Println("Server started at localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}
