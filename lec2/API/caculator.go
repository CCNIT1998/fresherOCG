package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Caculator(w http.ResponseWriter, r *http.Request) {
	methods := r.URL.Query()
	method := methods.Get("method")
	num1, _ := strconv.Atoi(methods.Get("num1"))
	num2, _ := strconv.Atoi(methods.Get("num2"))

	switch method {
	case "sum":
		fmt.Fprintln(w, "num1 = ",num1 ,"\nnum2 = ", num2)
		fmt.Fprintln(w, "Summation")
		fmt.Fprintln(w, num1,"+",num2,"=", num1+num2)
	case "sub":
		fmt.Fprintln(w, "num1 = ",num1 ,"\nnum2 = ", num2)
		fmt.Fprintln(w, "Subtraction")
		fmt.Fprintln(w, num1,"+",num2,"=", num1-num2)
	case "mul":
		fmt.Fprintln(w, "num1 = ",num1 ,"\nnum2 = ", num2)
		fmt.Fprintln(w, "Multiplication")
		fmt.Fprintln(w, num1,"*",num2,"=", num1*num2)
	case "div":
		fmt.Fprintln(w, "Division")
		if(num2 == 0){
			fmt.Fprintln(w, "num1 = ",num1 ,"\nnum2 = ", num2)
			fmt.Fprintln(w,"denominator should be different from 0")
		}else{
			fmt.Fprintln(w, "num1 = ",num1 ,"\nnum2 = ", num2)
			fmt.Fprintln(w, num1,"/",num2,"=", float64(num1)/float64(num2))
		}
	default:
		fmt.Fprintln(w, "url not found")
	}
}

func main() {
	fmt.Println("Starting server http://localhost:8090")
	fmt.Println("API sum: http://localhost:8090/caculator?method=sum&num1=3&num2=2")
	fmt.Println("API sub: http://localhost:8090/caculator?method=sub&num1=3&num2=2")
	fmt.Println("API mul: http://localhost:8090/caculator?method=mul&num1=3&num2=2")
	fmt.Println("API div: http://localhost:8090/caculator?method=div&num1=3&num2=2")
	defer func() {
		fmt.Println("Server is stopped")
	}()
	http.HandleFunc("/caculator", Caculator)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("Error when running server")
	}


}
