package main  
  
import "fmt"  
import (  
    "math/rand"  
    //"time"  
    "time"  
)  
func main() {  
    fmt.Print(rand.Intn(100))  
    fmt.Println()  
  
    fmt.Print(rand.Float64())    
    fmt.Println()  
      
    rand.Seed(time.Now().Unix()) 
    myrand := random(1, 20)  
  
    fmt.Println(myrand)  
  
}  
  
func random(min, max int) int {  
    return rand.Intn(max - min) + min  
}  