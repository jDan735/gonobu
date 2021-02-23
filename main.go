package main

import (
    "gopkg.in/yaml.v2"

    "math/rand"
    "io/ioutil"
    "strconv"
    "strings"
    "flag"
    "time"
    "fmt"
    "os"
)


var VERSION string = "0.1.0"
var locale Locale
var objects [3]string


type Options struct {
    Version bool
    Choice  int
    Lang    string
}


func main(){
    options := ParseArguments(os.Args[1:])

    locale = GetLocale(options.Lang)

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


type Locale struct {
    Lang struct {Name string; Code string; Cases [3]string}
    Game struct {Description string}
    Results struct {Win string; Lose string; Draw string}
    Objects struct {Rock string; Scissors string; Paper string}
    Bot struct {Have string}
}


func GetLocale(lang string) (locale Locale){
    file, _ := ioutil.ReadFile("locale/" + lang + ".yml")
    locale = Locale{}
    yaml.Unmarshal(file, &locale)

    objects = [3]string{
        locale.Objects.Rock,
        locale.Objects.Scissors,
        locale.Objects.Paper,
    }

    return
}


func ParseArguments(arguments []string) (Options) {
    var defaultOptions Options

    defaultLang := strings.Split(os.Getenv("LANG"), ".")[0]
    options := defaultOptions

    fs := flag.NewFlagSet("main", flag.ExitOnError)

    fs.StringVar(&options.Lang, "l", defaultLang, "Select language")
    fs.BoolVar(&options.Version, "v", false, "Show version")
    fs.IntVar(&options.Choice, "c", 0, "Select choice")

    fs.Parse(arguments)

    return options
}


func Game(user int){
    rand.Seed(time.Now().UnixNano())
    var bot int = rand.Intn(3)

    result := Battle(user, bot)
    fmt.Println(
        result.Color,
        result.Name + "!\033[0m",
        locale.Bot.Have + locale.Lang.Cases[bot],
        objects[bot],
    )
}


type Result struct {
    Type  int8
    Name  string
    Color string
}


func Battle(user int, bot int) (result Result){
    matrix := [3][3]int8{{2, 0, 1}, {1, 2, 0}, {0, 1, 2}}
    events := [3]string{
        locale.Results.Win,
        locale.Results.Lose,
        locale.Results.Draw,
    }

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
    LOGO_LEN   := 37              // Max len of line in logo
    LETTER_LEN := 8               // Len of bottom part of g

    fmt.Println("\033[1m\033[34m                         _")
    fmt.Println("  __ _  ___  _ __   ___ | |__  _   _")
    fmt.Println(" / _` |/ _ \\| '_ \\ / _ \\| '_ \\| | | |")
    fmt.Println("| (_| | (_) | | | | (_) | |_) | |_| |")
    fmt.Println(" \\__, |\\___/|_| |_|\\___/|_.__/ \\__,_|")

    description := locale.Game.Description
    descriptionLen := len(strings.Split(description, ""))

    spaceLen := LOGO_LEN - (LETTER_LEN + descriptionLen)
    space := ""

    for i := 0; i <= spaceLen; i++ {
        space += " "
    }

    fmt.Println(" |___/ \033[0m" + space + locale.Game.Description, "\n")
}


func Enter() (user int){
    for i := 0; i < 3; i++ {
        fmt.Println(
            strconv.Itoa(i + 1) + ".",
            strings.Title(objects[i]),
        )
    }

    fmt.Printf(">>> ")
    fmt.Scanf("%d\n", &user)
    fmt.Println()

    user--

    if user < 0 {
        fmt.Println("[\033[31mERROR\033[0m] user < 0")
        os.Exit(0)
    } else if user > 2 {
        fmt.Println("[\033[31mERROR\033[0m] user < 2")
        os.Exit(0)
    }

    return
}
