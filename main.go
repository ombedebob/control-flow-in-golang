package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var choices = []string{
	"rock",
	"paper",
	"scissors",
}

func main() {
	for {
		//get player and computer choice
		rand.Seed(time.Now().UnixNano())
		computer_choice := choices[rand.Intn(len(choices))]

		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("pick a choice: (rock,paper,scissors): ")
		input_choice, _ := scanner.ReadString('\n')
		player_choice := strings.Trim(input_choice, "\r\n")

		for index := range choices {
			switch {
			case player_choice == choices[index]:
				fmt.Printf("player choice: %s\n", player_choice)
				fmt.Printf("computer choice: %s\n", computer_choice)
			}
		}

		//determine winner
		if player_choice == computer_choice {
			fmt.Println("Its a tie")
		}

		if player_choice == "rock" {
			if computer_choice == "paper" {
				fmt.Println("you loose! Paper covers rock")
			} else {
				if computer_choice == "scissors" {
					fmt.Println("You win! rock smashes scissors")
				}
			}
		}

		if player_choice == "paper" {
			if computer_choice == "scissors" {
				fmt.Println("you loose! scissor cuts paper")
			} else {
				if computer_choice == "rock" {
					fmt.Println("You win! paper covers rock")
				}
			}
		}

		if player_choice == "scissors" {
			if computer_choice == "paper" {
				fmt.Println("you win! scissor cuts paper")
			} else {
				if computer_choice == "rock" {
					fmt.Println("You loose! rock smashes scissors")
				}
			}
		}

		//quit game and break out of loop
		fmt.Println("Press Enter to try again,'q' to quit")
		input, _ := scanner.ReadString('\n')
		quit_game := strings.Trim(input, "\r\n")
		if quit_game == "q" {
			return
		}
	}
}
