// 方法一
func fizzBuzz(n int) []string {
    var ret []string
    for i := 1; i <= n; i++ {
        switch {
            case i%15 == 0:
                ret = append(ret, "FizzBuzz")
            case i%3 == 0:
                ret = append(ret, "Fizz")
            case i%5 == 0:
                ret = append(ret, "Buzz")
            default:
                ret = append(ret, strconv.Itoa(i))
        }
    }
    return ret
}

// 方法二
func fizzBuzz(n int) []string {
    var ret []string
    for i := 1; i <= n; i++ {
        if i % 15 == 0 {
            ret = append(ret, "FizzBuzz")
        } else if i % 3 == 0 {
            ret = append(ret, "Fizz")
        } else if i % 5 == 0 {
            ret = append(ret, "Buzz")
        } else {
            ret = append(ret, strconv.Itoa(i))
        }
    }
    return ret
}