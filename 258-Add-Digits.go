/* 方法一：
 * 解决方法请参见：https://en.wikipedia.org/wiki/Digital_root
 */
func addDigits(num int) int {
    return 1 + (num - 1) % 9
}

// 方法二
func addDigits(num int) int {
    if num < 10 {
        return num
    }
    ret := 0
    for {
        for num > 0 {
           ret += num % 10
           num /= 10
        }
        if ret < 10 {
            break
        } else {
            num = ret
            ret = 0
        }
    }
    return ret
}