package piscine

func FirstRune(s string) rune {
  for _, value := range s {
    return value
  }
  return rune(s[0])
}