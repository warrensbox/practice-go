func selfDividingNumbers(left int, right int) []int {
    
    output := []int{}
  
 
    for ;left<=right;left++ {
   
        number := left
             fmt.Println("init number",number)
        for number > 0 {
            remaining := number % 10
            fmt.Println("remaining",remaining)
            if remaining == 0 || left%remaining != 0 {
                fmt.Println("break")
                break
            }
           
            number /= 10
            fmt.Println("number",number)
            fmt.Println("inner loop")
            fmt.Println("..........")
        }
        
        if number == 0 {
             output = append(output,left)
        }
       fmt.Println("outter loop")
        fmt.Println()
    }
    return output 
}