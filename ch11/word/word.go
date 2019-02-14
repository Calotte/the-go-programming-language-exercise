package word

import "unicode"

func IsPalindrome(s string)bool{
	var letters []rune
	//letters:=make([]rune,0,len(s))
	for _,r := range s{
		if unicode.IsLetter(r){
			letters= append(letters, unicode.ToLower(r))
		}
	}
	n:=len(letters)/2
	for i:=range letters{
		if letters[i]!=letters[n-i-1]{
			return false
		}
	}
	return true
}
