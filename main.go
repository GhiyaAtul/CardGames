package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


cards:= [13]{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" We Will be Playing Blackjack")
	fmt.Println("===============================")
	fmt.Printf("===============================\n\n")

L:

	for {
		fmt.Println("Type anyone of the below commands:")
		fmt.Println("Start")
		fmt.Println("Deal")
		fmt.Println("Stop")
		fmt.Println("End")
		fmt.Println("Enter Choice followed by Enter Key")
		choice, _ := reader.ReadString('\n')
		switchChoice := strings.ToUpper(strings.TrimSpace(choice))

		switch switchChoice {
		case "START":
			fmt.Println("Started the play")
			break
		case "DEAL":
			fmt.Println("Deal a Card")
			break
		case "STOP":
			fmt.Println("STOPPED the Play")
			break
		case "END":
			fmt.Println("Ended the Play")
			break L
		default:
			fmt.Println("======Default")
			break
		}
	}

}


func startGame(){


}


func playerCallHit(){


}

func playerCallStand(){


}

func dealCard(){


}

func calculatePlayerCardValue(){

}

func calculateDealerCardValue(){


}

func updateDeck(){


}

func EndGame(){


}