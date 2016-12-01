func findTheDifference(s string, t string) string {
    var ret rune
    for _, c := range s {
        ret ^= c 
    }
    for _, c := range t {
        ret ^= c 
    }
    return string(ret) 
}