func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    map1 := make(map[byte]int)
    map2 := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        map1[s[i]] = map1[s[i]]+1
        map2[t[i]] = map2[t[i]]+1
    }
    for k,v := range map1 {
        if map2[k] != v {
            return false
        }
    }
    return true
}
