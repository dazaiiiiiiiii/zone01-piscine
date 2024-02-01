package piscine

func check(a byte) bool {
  return ((a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9'))
}
func Capitalize(s string) string {
  sure := true
  str := []byte(s)
  for i := 0; i < len(s); i++ {
    if sure && check(s[i]) {
      if (s[i] >= 'a' && s[i] <= 'z') {
        str[i] -= 32
      }
      sure = false
    } else if sure == false &&str[i] >= 'A' && str[i] <= 'Z'{
       str[i] += 32
    } else if !(check(s[i])) {
      sure = true
    }
  }
  return string(str)
}