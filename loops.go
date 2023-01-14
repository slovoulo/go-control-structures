//Create a simple program with sufficient error handling tha allows a user
//to perform 4 mathematical operations based on a number of their choosing

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)


var reader = bufio.NewReader(os.Stdin)

//Step1
func getUserChoice()(string,error){
    fmt.Println("1) Add up all numbers up to number X")
    fmt.Println("2) Build the factorial up to number X")
    fmt.Println("3) Sum up manually entered numbers")
    fmt.Println("4) Sum up a list of numbers")  //Loop through a list
    fmt.Println("5) Get the square of a number")

    fmt.Println("Choose an option from 1 to 5")
    fmt.Println("--------------------------")

    userInput, err := reader.ReadString('\n') 

    //Check if there is an error in the user input
    if (err!=nil){
        return "",err
    }
    //Remove all spaces from user input
    userInput = strings.TrimSpace(userInput)

    //Check if user input is anything different from 1-5
    //if different return empty string and custom error message
    if(
        userInput=="1"||
        userInput=="2"||
        userInput=="3"||
        userInput=="4"||
        userInput=="5"){

            return userInput,nil
            
        }
        //if user input passes error check  return the input and a nil error
        
        return "",errors.New("INVALID CHOICE")
   
   }

   //Mathematical functions
   func addUpAllNumbers(){
    sum:=0
     //Call function to get user number
     userNumber,err:=getUserNumber()
     //End operation if there is an error fetching input
     if(err!=nil){
        fmt.Println("An error occured")
        return  }
        //If no error proceed with entered number
        for i:=0; i<=userNumber; i++ {
            sum=sum+i
        }

        fmt.Printf("The sum is %v", sum)
    
   }
   func buildFactorial(){
      
    factorial:=1
  
     //Call function to get user number
     userNumber,error:=getUserNumber()
     //End operation if there is an error fetching input
   
     if(error!=nil){
        fmt.Println("An error occured")
        return }
        //If no error proceed with entered number
        for i:=1; i<=userNumber; i++ {
            factorial=factorial*i
        }

        fmt.Printf("The factorial is %v", factorial)
   }
   func manualSummation(){
    //Here we use a modification of for loops which is golangs way of handling while loops
    //The logic is to start with a value that is set to true, keep looping until that value becomes false
    //BONUS: Save user input values in a list and print them when displaying their sum
    sum:=0
    isEnteringNumbers:=true
    userSelections:=[]int{}
    
    //for as long as is entering numbers is true, keep running the function that prompts user for a number
    for isEnteringNumbers{
        userNumber,error:=getUserNumber()
        userSelections=append(userSelections, userNumber)
        sum=userNumber+sum

        //If user enters a wrong keystroke like a letter, stop asking for input
        
            isEnteringNumbers=error==nil
             
        
    }
        fmt.Printf("The total sum of %v is %v", userSelections, sum)
   }
   func sumList()error{
    //Looping through a list of numbers
    fmt.Println("Please enter a comma-separated list of numbers")
    //Get user input as a list
    userList, err := reader.ReadString('\n') 
    if (err!=nil){
        fmt.Println("Invalid input")
        return err  //stop function if there is an error
    }

    userList=strings.TrimSpace(userList)
    //Convert the list of user strings to slice
    noCommas:=strings.Split(userList, ",") //Enter the list of strings to be modified(userList) followed by the part of the strings you want removed(,)

    //Use the range keyword to loop through the array
    //range returns a key value pair
    sum:=0
    for key,value:=range noCommas{
        fmt.Printf("Key %v : value %v\n", key,value)
        //convert values to int
        intValues,_:=strconv.ParseInt(value,0,64) 

        sum=sum+int(intValues)

         
    }
    fmt.Printf("The sum of %v is %v\n", noCommas,sum)
    return nil


   }
   func getSquareRoot(){}

   func getUserNumber()(int,error){
    fmt.Println("Enter a number for arithmetic operation")
    //Get the number user entered
    userNumber, err := reader.ReadString('\n') 
    if (err!=nil){
        return 0,err //stop function if there is an error
    }

    userNumber=strings.TrimSpace(userNumber)
    userNumberInt, numerr:=strconv.ParseInt(userNumber,0,64)
    if(numerr!=nil){
        return 0, numerr
    } 
        //If there is no error convert the input from base 64 int to int
    return int(userNumberInt),nil
   }
   
func main() {

    //step 2
    //fetch user input
    userSelection,error:=getUserChoice()
    if(error!=nil){
        fmt.Println("Invalid choice, exiting")
        return
    }


    //Execute mathematical operations
    if (userSelection=="1"){
        addUpAllNumbers()
    }else if(userSelection=="2"){
        buildFactorial()
       
    } else if (userSelection=="3"){
        manualSummation()
            
    }else if (userSelection=="4"){
        //Looping through arrays
        sumList()
            
    }else {
            
    }

    
}