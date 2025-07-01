package main

// func sum(c chan int, s []int) {
// 	total := 0
// 	for _, v := range s {
// 		total += v
// 	}
// 	c <- total
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int)
// 	defer close(c)
// 	go sum(c, s[:len(s)/2])
// 	go sum(c, s[len(s)/2:])
// 	x, y := <-c, <-c
// 	println("Sum of first half:", x)
// 	println("Sum of second half:", y)
// 	println("Total sum:", x+y)
// }




// package main

// import "fmt"

// func fibonacci(c, quit chan int){
// 	x, y := 0, 1
// 	for{
// 		select{
// 			case c <- x:
// 				x, y := y, x+y
			
// 			case <-quit:
// 				fmt.Println("quit")
// 				return
// 		}
// 	}
// }

// func main(){
// 	c := make(chan int)
// 	quit := make(chan int)

// 	go func(){
// 		for i=0; i<10; i++{
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)

// }


