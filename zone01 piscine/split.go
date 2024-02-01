package piscine 

func Split(s, sep string) []string {
  first := 0
  var result []string
  for i := 0; i + len(sep) <= len(s); i++ {
    if sep == s[i:i+len(sep)] {
      result = append(result, s[first:i])
      first = i + len(sep)
      i = first
    }
  }
  if first < len(s) {
    result = append(result, s[first:])
  }
  return result
}