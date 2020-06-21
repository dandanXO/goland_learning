package strutil

func Reverse(s string) (ret string) {
    for _, v := range s {
        defer func(r rune) { ret += string(r) }(v)
    }
    return
}