package word

import "testing"

func BenchmarkIsPalindrome(t *testing.B){
	for i:=0;i<t.N;i++{
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}