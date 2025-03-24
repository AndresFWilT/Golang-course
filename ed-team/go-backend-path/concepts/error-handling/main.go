package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found")

var food = map[int]string{
	1: "Pizza",
	2: "Hamburger",
	3: "Soap",
}

func main() {
	key, err := search("5")
	if errors.Is(err, errNotFound) {
		fmt.Println("Not Found Error controlled, returning a valid response..")
		return
	}
	if err != nil {
		fmt.Printf("search(): %v\n", err)
	}
	fmt.Println(key)
}

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}
	food, err := findFood(num)
	if err != nil {
		return "", fmt.Errorf("findFood(): %w", err)
	}
	return food, nil
}

func findFood(index int) (string, error) {
	searchValue, ok := food[index]
	if !ok {
		return "", errNotFound
	}
	return searchValue, nil
}
