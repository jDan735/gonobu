package gonobu

import (
	"math/rand"
	"time"
	"fmt"
	"os"
)


func Start(){
	Welcome()
	enter := Enter()

	Game(enter)
}


func Game(enter int){
    objects := [3]string{"rock", "scissors", "paper"};

	rand.Seed(time.Now().UnixNano())
	var bot int = rand.Intn(3)

	result := Battle(enter, bot)
	fmt.Println(result.Color, result.Name + "!\033[0m", "Bot has chosen", objects[bot])
}


type Result struct {
	Type  int8
	Name  string
	Color string
}


func Battle(enter int, bot int) (result Result){
    matrix := [3][3]int8{{2, 0, 1}, {1, 2, 0}, {0, 1, 2}}
    events := [3]string{"Win", "Lose", "Draw"}
    colors := [3]string{"\033[32m", "\033[31m", "\033[33m"}

    for i := 0; i < 3; i++ {
    	if bot == i {
    		event := matrix[enter][bot]
		    result = Result{
		    	event,
		    	events[event],
		    	colors[event],
		    }
    	}
    }

    return
}


func Welcome(){
	fmt.Println("\033[1m\033[34m                         _")
	fmt.Println("  __ _  ___  _ __   ___ | |__  _   _")
	fmt.Println(" / _` |/ _ \\| '_ \\ / _ \\| '_ \\| | | |")
	fmt.Println("| (_| | (_) | | | | (_) | |_) | |_| |")
	fmt.Println(" \\__, |\\___/|_| |_|\\___/|_.__/ \\__,_|")
	fmt.Println(" |___/\033[0m           Kanobu written in go \n")

	fmt.Println("\033[37m1.\033[0m Rock")
	fmt.Println("2. Scissors")
	fmt.Println("3. Paper")
}


func Enter() (enter int){
	fmt.Printf(">>> ")
	fmt.Scanf("%d\n", &enter)

	fmt.Println()

	enter = enter - 1

	if enter < 0 {
		fmt.Println("[\033[31mERROR\033[0m] Enter < 0")
		fmt.Println("[\033[33mNOTE\033[0m]  Enter number from 1 until 3.")
		os.Exit(0)
	} else if enter > 2 {
		fmt.Println("[\033[31mERROR\033[0m] Enter < 2")
		fmt.Println("[\033[33mNOTE\033[0m]  Enter number from 1 until 3.")
		os.Exit(0)
	}

	return
}
