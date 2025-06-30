package main

// "fmt"

// "errors"
// "strings"

func main() {
	// printme("Hello, World!")
	// var numu int = 10
	// var deno int = 0
	// q, d, err := divide(numu, deno)
	// fmt.Println(q)
	// if err!= nil{
	// 	fmt.Println("Error: ", err)
	// }else {
	// 	fmt.Printf("Quotient: %d, Remainder: %d\n", q, d)
	// }
	// var ptr *int
	// q := 10
	// ptr = &q
	// fmt.Printf("%d\n", *ptr)
	// 	s := "😀"
	// fmt.Println(len(s))
	// text := "hello world"

	// upper := strings.ToUpper(text)
	// fmt.Println(upper)

	// fmt.Println(strings.Contains(text, "world"))

	// replaced := strings.Replace(text, "world", "Go", 1)
	// fmt.Println(replaced)

	// words := strings.Split(text, " ")
	// fmt.Println(words)

	// trimmed := strings.TrimSpace("  spaced  ")
	// fmt.Println(trimmed)
	// num := 1
	// switch num {
	// case 1:
	// 	fmt.Println("One")
	// 	fallthrough // By default, Go does not fall through to the next case. If you want it to continue into the next case, use fallthrough:
	// case 2:
	// 	fmt.Println("Two")
	//}

	// nums := []int{1, 2, 3, 4}
	// for ind, val:=range nums{
	// 	fmt.Println(ind, val)
	// }

	// nums := []int{1, 2, 3, 4}
	// for _, val:=range nums{
	// 	fmt.Println(val)
	// }

	// var input string
	// for input != "exit"{
	// 	fmt.Println("Type the input")
	// 	fmt.Scanln(&input)
	// }
	// fmt.Println("GoodBye")

	// score := map[int]int{
	// 	5:5,
	// 	20:20,
	// }
	// for _, val:= range score{
	// 	fmt.Println(val)
	// }

	// a, b := split(1000)
	// fmt.Println(a,b)
	// fmt.Println(total(1,2,3))
	// fmt.Println(total(10,20,30))

	// nums := []int{4, 5, 6}
	// fmt.Println(total(nums...))

	// fmt.Println(fibnacci(10))
	// test := func(x int) int{
	// 	return x*x
	// }
	// test2(test)
	// returnfunc("Subash")()

	// var arr = [...]int{1,2,3,4,5,6,7}
	// for _, val := range arr{
	// 	fmt.Println(val)
	// }

	// 	mychannel := make(chan string)

	// 	go func() {
	// 		mychannel <- "data"
	// 	}()
	// 	msg := <-mychannel
	// 	fmt.Println(msg)

	// 	nums := []int{1, 2, 3, 4, 5}
	// 	datachannel := slicetoChannel(nums)
	// 	final := sq(datachannel)
	// 	for n := range final {
	// 		fmt.Println(n)
	// 	}


	
}

// func slicetoChannel(nums []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func sq(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func returnfunc(x string) func(){
// 	return func(){
// 		fmt.Println(x)
// 	}
// }

// func test2(text func(x int) int){
// 	fmt.Println(text(10))
// }

// func isEven(n int) bool {
// 	return n%2 == 0
// }

// func fibnacci(n int) int {
// 	if n <= 1	{
// 		return n
// 	}
// 	return fibnacci(n-1) + fibnacci(n-2)
// }

// func total(nums ...int)(x int){
// 	for _, val:=range nums{
// 		x += val
// 	}
// 	return
// }

// func split(sum int) (x int, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }
// func printme(s string){
// 	fmt.Println(s)
// }

// func divide(nume int, deno int)(int, int, error) {
// 	if(deno == 0){
// 		err := errors.New("denominator cannot be zero")
// 		return 0, 0, err
// 	}
// 	var q int = nume / deno;
// 	var r int = nume % deno;
// 	return q, r, nil
// }
