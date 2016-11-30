/*
 * LeetCode可能有bug，这个答案不能accept
 * 出错的输入是：
 * Input: "Snug & raw was I ere I saw war & guns."
 * Output: ".snug \u0026 raw was I ere I saw war \u0026 gunS"
 * Expected: ".snug & raw was I ere I saw war & gunS"
 *
 * 在go1.7中执行结果是： ".snug & raw was I ere I saw war & gunS"
 * 结果正确，说明这确实是LeetCode的问题
 */
func reverseString(s string) string {
    var ret []byte
    for i := len(s)-1; i >= 0; i-- {
        ret = append(ret, byte(s[i]))
    }
    return string(ret)
}