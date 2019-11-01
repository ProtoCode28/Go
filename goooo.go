package main

//https://www.youtube.com/watch?v=H-IUE0idp7k

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)
var	question string



func main() {

	fmt.Println("To exit the program type exit as your question")

		for question != "exit" {

			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Type your question")
			scanner.Scan()
			question = scanner.Text()

			if question != "exit" {fmt.Println("The magic conch says") }

			rand.Seed(time.Now().UTC().UnixNano())
			n:= rand.Intn(20)

			switch n {

			case 0 :
				fmt.Println("It is certain")

			case 1 :
				fmt.Println("It is decidedly so")

			case 2 :
				fmt.Println("Without a doubt")

			case 3 :
				fmt.Println("Yes - definitely")

			case 4 :
				fmt.Println("You may rely on it")

			case 5 :
				fmt.Println("As I see it, yes")

			case 6 :
				fmt.Println("Most likely")

			case 7 :
				fmt.Println("Outlook good")

			case 8 :
				fmt.Println("Yes")

			case 9 :
				fmt.Println("Signs point to yes")

			case 10 :
				fmt.Println("Reply hazy, try again")

			case 11 :
				fmt.Println("Ask again later")

			case 12 :
				fmt.Println("Better not tell you now")

			case 13 :
				fmt.Println("Cannot predict now")

			case 14 :
				fmt.Println("Concentrate and ask again")

			case 15 :
				fmt.Println("Don't count on it")

			case 16 :
				fmt.Println("My reply is no")

			case 17 :
				fmt.Println("My sources say no")

			case 18 :
				fmt.Println("Outlook not so good")

			case 19 :
				fmt.Println("Very doubtful")


			}
}

}