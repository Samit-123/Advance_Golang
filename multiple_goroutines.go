package main
import(
         "fmt"
         "time"
)
func Aname(){
	arr1:=[4]string{"Rohit","Suman","Aman","Ria"}
	for t1:=0;t1<=3;t1++{
		time.Sleep(150*time.Millisecond)
		fmt.Printf("%s\n",arr1[t1])
	}
}
func Bname(){
	arr2:=[4]int{300,301,302,303}
	for t2:=0;t2<=3;t2++{
		time.Sleep(500*time.Millisecond)
		fmt.Printf("%d\n",arr2[t2])
	}
}	
func main(){
	fmt.Println("...............Main_Goroutine_start!...............")
	go Aname()
	go Bname()
	time.Sleep(3500 * time.Millisecond) 
	fmt.Println("...............Main_Goroutine_end!.................")
}
