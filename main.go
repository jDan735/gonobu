package main

import (
	"math/rand"
    "flag"
	"time"
	"fmt"
	"os"
)


var VERSION string = "0.0.1"


type Options struct {
	Version bool
	Choice  int
}


func main(){
	options := ParseArguments(os.Args[1:])

	if options.Version {
		fmt.Println("v" + VERSION)
	} else {
		Logo()

		if options.Choice != 0 {
			if options.Choice--; options.Choice > 0 && options.Choice < 3 {
				Game(options.Choice - 1)
			} else {
				fmt.Println("[\033[31mERROR\033[0m] Choice > 0 && Choice < 3")
			}
		} else {
			Game(Enter())
		}
	}
}


func ParseArguments(arguments []string) (Options) {
	var defaultOptions Options
	options := defaultOptions

    fs := flag.NewFlagSet("main", flag.ExitOnError)

    fs.BoolVar(&options.Version, "v", false, "Show version")
    fs.IntVar(&options.Choice, "c", 0, "Select choice")

    fs.Parse(arguments)

    return options
}


func Game(user int){
    objects := [3]string{"rock", "scissors", "paper"};

	rand.Seed(time.Now().UnixNano())
	var bot int = rand.Intn(3)

	result := Battle(user, bot)
	fmt.Println(result.Color, result.Name + "!\033[0m", "Bot has chosen", objects[bot])
}


type Result struct {
	Type  int8
	Name  string
	Color string
}


func Battle(user int, bot int) (result Result){
    matrix := [3][3]int8{{2, 0, 1}, {1, 2, 0}, {0, 1, 2}}
    events := [3]string{"Win", "Lose", "Draw"}
    colors := [3]string{"\033[32m", "\033[31m", "\033[33m"}

    for i := 0; i < 3; i++ {
    	if bot == i {
    		state := matrix[user][bot]
		    result = Result{
		    	state,
		    	events[state],
		    	colors[state],
		    }
    	}
    }

    return
}


func Logo(){
	fmt.Println("\033[1m\033[34m                         _")
	fmt.Println("  __ _  ___  _ __   ___ | |__  _   _")
	fmt.Println(" / _` |/ _ \\| '_ \\ / _ \\| '_ \\| | | |")
	fmt.Println("| (_| | (_) | | | | (_) | |_) | |_| |")
	fmt.Println(" \\__, |\\___/|_| |_|\\___/|_.__/ \\__,_|")
	fmt.Println(" |___/\033[0m           Kanobu written in go \n")
}


func Enter() (user int){
	fmt.Println("\033[37m1.\033[0m Rock")
	fmt.Println("2. Scissors")
	fmt.Println("3. Paper")
	fmt.Printf(">>> ")

	fmt.Scanf("%d\n", &user)

	fmt.Println()

	user = user - 1

	if user < 0 {
		fmt.Println("[\033[31mERROR\033[0m] user < 0")
		fmt.Println("[\033[33mNOTE\033[0m]  user number from 1 until 3.")
		os.Exit(0)
	} else if user > 2 {
		fmt.Println("[\033[31mERROR\033[0m] user < 2")
		fmt.Println("[\033[33mNOTE\033[0m]  user number from 1 until 3.")
		os.Exit(0)
	}

	return
}
