package piscine

func LastRune(s string) rune {
  var char rune
  for _, value := range s {
    char = value
  }
  return char
}