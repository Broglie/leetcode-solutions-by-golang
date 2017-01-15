func hammingDistance(x int, y int) int {
    var dest = 0
    for temp := x ^ y; temp != 0; temp &= temp - 1 {
        dest++
    }
    return dest
}