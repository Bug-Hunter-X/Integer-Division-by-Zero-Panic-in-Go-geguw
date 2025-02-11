func myFunc(x int) (int, error) {
    if x == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    result := 10 / x
    return result, nil
}

func main(){
  result, err := myFunc(0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
  result, err = myFunc(2)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
}