package main

 import "fmt"
 import "myownfunctions"  // <-- remember to include the directory/package name - in this case "Project"

 func main(){
     vara := myownfunctions.SayHello() //<---- function name must be capital!
     varb := myownfunctions.SayWorld()
     fmt.Println(vara+" "+varb)
 }
