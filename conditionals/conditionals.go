package conditionals

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GameDecision()error{
    err1 :="Not an integer"
    gamePrice,_ := reader.ReadString('\n') 

    

	//save user input in variables
	gamePrice = strings.TrimSpace(gamePrice)
	

	//Convert price string to float
	urgamePrice,err :=strconv.ParseInt(gamePrice,0,64)  
    if(err!=nil){
        fmt.Println(err1)
        fmt.Println(err)
        fmt.Println(errors.New("INVALID INPUT"))
        return nil //"return nil ensures the function doesnt continue running"
    }

    if(urgamePrice>40){
        fmt.Println("too expensive, wait for discounts")
    } else if(urgamePrice<3){
        fmt.Printf("%v Too cheap, not worth your time", urgamePrice)

    }else{
        fmt.Println("BUY IT NOW")
    }

    return err

}



func main() {
    //Prompt user for struct values
	fmt.Println("Should you buy this game?")
	fmt.Println("-----------------------")
	fmt.Println("Enter current game price: ")
	GameDecision()

	

}