package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

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


func day1_1(input string){
    lines := strings.Split(strings.TrimSpace(input), "\n")

    var left []int
    var right []int

    for _, line := range lines {
        parts := strings.Fields(line)

        leftNum, _ := strconv.Atoi(parts[0])
        rightNum, _ := strconv.Atoi(parts[1])

        left = append(left, leftNum)
        right = append(right, rightNum)
    }

    sort.Ints(left)
    sort.Ints(right)

    totalDistance := 0

    for i := 0; i < len(left); i++ {
        distance := left[i] - right[i]
        if distance < 0 {
            distance *= -1
        }

        totalDistance += distance
    }
    fmt.Println("Day 1 output: ", totalDistance)
}


func main() {
    input := getInput(1)
    day1_1(input)
}
