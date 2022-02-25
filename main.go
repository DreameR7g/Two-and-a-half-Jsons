package main

import (
	"encoding/json"
	"fmt"
)

type likesStuff struct {
	Name string `json:"name"`

	Likes string `json:"likes"`

	Dislikes string `json:"dislikes"`

	Skills string `json:"skills"`
}

func main() {
	stuff := make(map[string]likesStuff)

	stuff["Dreamer"] = likesStuff{
		Name:     "Dreamer Rigorstorm",
		Likes:    "Friends, Raid, Go, Kisa",
		Dislikes: "People who bully LaLafells",
		Skills:   "Making people laugh, fixing things",
	}
	stuff["Siryu"] = likesStuff{
		Name:     "Siryu Unslay",
		Likes:    "Raiding, Food, Cherry",
		Dislikes: "Being smol",
		Skills:   "Tanking, Coding, Teaching",
	}
	stuff["Kisa"] = likesStuff{
		Name:     "Kisa Rozen",
		Likes:    "Yellow, gardening, bunnies",
		Dislikes: "ORANGE!!",
		Skills:   "Doodling, Singing, Knitting",
	}
	stuff["James"] = likesStuff{
		Name:     "James Carmike",
		Likes:    "Music, Hugs, Friends",
		Dislikes: "Disorganization",
		Skills:   "Being cute, Analyzing things, ????",
	}

	output, err := json.Marshal(stuff)
	if err != nil {
		fmt.Println("This error is:", err)
		return
	}

	fmt.Println(output)

	arrays()

}

func arrays() {
	boop := [6]int{2, 3, 4, 5, 6, 7}

	output, err := json.Marshal(boop)
	if err != nil {
		fmt.Println("This error is:", err)
		return
	}
	fmt.Println(output)
}
