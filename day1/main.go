package main

import (
    "fmt"
    "os"
    "io"
    "net/http"

    "github.com/joho/godotenv"
)


func getInput(day int) string {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    inputFile := fmt.Sprintf("day%d.txt", day)
    _, err = os.Stat(inputFile)
    if err == nil {
        content, _ := os.ReadFile(inputFile)
        return string(content)
    }


	sessionCookie := os.Getenv("SESSION")

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, _ := http.NewRequest("GET", url, nil)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	_ = os.WriteFile(inputFile, body, os.ModePerm)

	return string(body)
}


func main() {
    fmt.Println(getInput(1))
}
