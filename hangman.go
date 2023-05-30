package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wins := 0
	loses := 0
	numberOfLetters := 0

	reInitRandomValue(&numberOfLetters)
	again, hasWon :=playHangman(numberOfLetters)

	for	{
		if hasWon ==true {
			wins++
		}else{
			loses++ 
		}
		if again =="y" {
			printScore(wins,loses)
			reInitRandomValue(&numberOfLetters)
			again, hasWon =playHangman(numberOfLetters)

		}else if again =="n" {
			printScore(wins,loses)
            break
				}
	}
	fmt.Println("Thanks for playing")

}

func playHangman(numberOfLetters int) (playAgain string, iswinner bool) {
	stageOfHangman :=0
	gameType :=""
	hasGuessedALetter := false
	hasWon := false
	guess := ""
	guessedLetters :=""
	again :=""
	dashes :=""
	newDashes := ""
	fmt.Println("--------------------------------\n")
	fmt.Println("Welcome to Hangman\n")
	fmt.Println("--------------------------------\n")
	for{
		fmt.Println("Select game type")
		fmt.Println("Press e for Easy level")
		fmt.Println("Press h for Hard level")
		fmt.Scanln(&gameType)
		gameType =strings.ToLower(gameType)
		if gameType == "e" || gameType =="h" {
			clearScreen()
			break
		} else {
			fmt.Println("Press e or h only")
		}
	}

	word :=chooseRandomWord(numberOfLetters,gameType)

	fmt.Println()
	for{
		drawHangman(stageOfHangman)
		if stageOfHangman == 10 {
					fmt.Println("Your hangman died!\n")
					fmt.Println("The word was: %s\n",strings.ToUpper(word))
					return wantToPlayAgain(), false
				}

				if hasGuessedALetter == false {
					dashes = hideTheWord(len(word))
					fmt.Printf(" %s\n", dashes)
				} else {
					fmt.Printf(" %s\n", newDashes)
				}
				fmt.Printf("\n Guess a letter: ")
				fmt.Scanln(&guess)

				isALetter, someKindOfError := regexp.MatchString("^[a-zA-Z]", guess)
		if someKindOfError != nil {
			clearScreen()
			fmt.Printf("Something went horribly wrong. ")
			fmt.Printf("Exiting with error can not regex match %v", again)
			return
		}

		// we allow upper case but continue with lower case only
		guess = strings.ToLower(guess)

		if isALetter == false {
			clearScreen()
			fmt.Printf("That is not even a letter. Try again.\n")
		} else if len(guess) > 1 {
			clearScreen()
			fmt.Printf("You entered more than 1 character. Try again.\n")
		} else if strings.Contains(guessedLetters, guess) {
			clearScreen()
			fmt.Printf("You have already guessed that letter. Try again.\n")
		} else if strings.Contains(word, guess) {
			clearScreen()
			fmt.Printf("Great! The letter you guessed is in the word.\n")
			guessedLetters += guess

			if hasGuessedALetter == false {
				updatedDashes := revealDashes(word, guess, dashes)
				newDashes = updatedDashes
			} else {
				updatedDashes := revealDashes(word, guess, newDashes)
				newDashes = updatedDashes
			}

			// we have a hit, a good time to check
			// if all revealed letters are identical
			// if yes the player won
			hasGuessedALetter = true
			if newDashes == strings.ToUpper(word) {
				hasWon = true
			}

			// output in case of a winning a round
			if hasWon == true {
				clearScreen()
				fmt.Printf("### C O N G R A T U L A T I O N S ###\n")
				fmt.Println()
				fmt.Printf(" _   _\n")
				fmt.Printf("  \\O/\n")
				fmt.Printf("   |\n")
				fmt.Printf("   |\n")
				fmt.Printf("  / \\\n")
				fmt.Println()
				fmt.Printf("You won this round! The word was: %s\n", strings.ToUpper(word))
				fmt.Printf("You saved hangman in %v of 10 guesses.\n", stageOfHangman)
				return wantToPlayAgain(), true
			}
		} else {
			clearScreen()
			fmt.Printf("The letter you guessed is not in the word. :(\n")
			stageOfHangman++
			guessedLetters += guess
		}
	}
}

func wantToPlayAgain() string{
	for{
		again := ""
		fmt.Println("Wanna play again? [y/n]\n")
		fmt.Scanln(&again)
		isYorN, someKindOfError := regexp.MatchString("^Y|y|N|n",again)
		if someKindOfError != nil {
			fmt.Printf("Something went horribly wrong. ")
			fmt.Printf("Exiting with error can not regex match %v", again)
			panic(someKindOfError)
		}

		if isYorN == false {
			fmt.Printf("You didn't type [y] or [n}. Try again.\n")
		} else if len(again) > 1 {
			fmt.Printf("You entered more than 1 character. Try again.\n")
		} else if strings.ToLower(again) == "y" {
			return "y"
		} else {
			return "n"
		}
	}
}	


func chooseRandomWord(numberOfLetters int,gameType string) string {
	switch gameType {
	case "e":
		var lettersData []byte
		var err error
		if numberOfLetters ==4 {
			lettersData, err = os.ReadFile("words/simple4letters.txt")
		} else if  numberOfLetters ==5 {
			lettersData, err = os.ReadFile("words/simple5letters.txt")
		} else if  numberOfLetters >=6 {
			lettersData, err = os.ReadFile("words/simple6letters.txt")
		} 
		if err != nil {
			panic(err)
		}
		dataString := string(lettersData)
		someWords := strings.Split(dataString, " ")
		randomNumber :=rand.Intn(len(someWords)-1)
		choseWord := someWords[randomNumber]
		return choseWord

	case "h":
		var lettersData []byte
		var err error
		if numberOfLetters == 4 {
			lettersData, err = os.ReadFile("words/all4letters.txt")
		} else if numberOfLetters == 5 {
			lettersData, err = os.ReadFile("words/all5letters.txt")
		} else if numberOfLetters == 6 {
			lettersData, err = os.ReadFile("words/all6letters.txt")
		} else if numberOfLetters == 7 {
			lettersData, err = os.ReadFile("words/all7letters.txt")
		} else if numberOfLetters == 8 {
			lettersData, err = os.ReadFile("words/all8letters.txt")
		} else if numberOfLetters == 9 {
			lettersData, err = os.ReadFile("words/all9letters.txt")
		} else if numberOfLetters == 10 {
			lettersData, err = os.ReadFile("words/all10letters.txt")
		} else if numberOfLetters == 11 {
			lettersData, err = os.ReadFile("words/all11letters.txt")
		} else if numberOfLetters == 12 {
			lettersData, err = os.ReadFile("words/all12letters.txt")
		} else if numberOfLetters == 13 {
			lettersData, err = os.ReadFile("words/all13letters.txt")
		} else if numberOfLetters == 14 {
			lettersData, err = os.ReadFile("words/all14letters.txt")
		} else if numberOfLetters == 15 {
			lettersData, err = os.ReadFile("words/all15letters.txt")
		}
		if err != nil {
			panic(err)
		}
		dataString := string(lettersData)
		someWords := strings.Split(dataString, " ")
		randomNumber := rand.Intn(len(someWords) - 1)
		chosenWord := someWords[randomNumber]
		return chosenWord
	}
	return "this you will only see in case of a weird bug"
}

func reInitRandomValue(toInit *int){
	*toInit = rand.Intn(11)+4
}

func hideTheWord(wordLength int) string {
	dashes := ""
	for i := 0; i < wordLength; i++ {
		dashes += "_"
	}
	return dashes
}

func revealDashes(word string, guess string, dashes string) string {
	newDashes := ""
	for i, r := range dashes {
		if c := string(r); c != "_" {
			newDashes += c
		} else {
			var letter = string(word[i])
			if guess == letter {
				newDashes += strings.ToUpper(guess)
			} else {
				newDashes += "_"
			}
		}
	}
	return newDashes
}

func printScore(wins,loses int){
	clearScreen()
	fmt.Println("--------------------------------\n")
	fmt.Println("ScoreCard \n")
	fmt.Println("Wins:",wins,"\t")
	fmt.Println("Loses:",loses,"\n")
	fmt.Println("--------------------------------\n")
}
func clearScreen(){
	if runtime.GOOS != "windows"{
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd","/c","cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func drawHangman(stageOfHangman int)	{
	switch stageOfHangman {
	case 0:
		fmt.Printf("   +---+\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 1:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 2:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 3:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf("  /|   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 4:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 5:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\  |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 6:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 7:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 8:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("  /    |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 9:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("  / \\  |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 10:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   0   |\n")
		fmt.Printf("  /|\\  |\n")
		fmt.Printf(" ° | ° |\n")
		fmt.Printf("  / \\  |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" R.I.P.|\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	}

}
