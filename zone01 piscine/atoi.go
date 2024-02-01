package piscine 

func Atoi(s string) int {
  sign := 1
  res := 0
  for i := 0; i < len(s); i++ {
    if i == 0 && (s[i] == '-' || s[i] == '+') {
      if s[i] == '-' {
        sign = -sign
      }
    }else if s[i] >= '0' && s[i] <= '9' {
      res *= 10
      res += int(s[i] - '0')
    } else {
      return 0
    }
  }
  return res * sign
}