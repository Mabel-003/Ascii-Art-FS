package main

import (
	"fmt"
	"strconv"
	"slices"
	"sort"
	"strings"
)

// func up(t string) string {
// 	return strings.ToUpper(t)
// }
// func low(t string) string {
// 	return strings.ToLower(t)
// }
// func cap(t string) string {
// 	if len(t) == 0 {
// 		return t
// 	} else {
// 		return strings.ToUpper(t[:1]) + strings.ToLower(t[1:])
// 	}
// }
// func bin(s string) string {
// 	n, _ := strconv.ParseInt(s, 2, 64)
// 	return strconv.FormatInt(n, 10)
// }
// func hex(s string) string {
// 	n, _ := strconv.ParseInt(s, 16, 64)
// 	return strconv.FormatInt(n, 10)
// }
// func article(text string) string {
// 	return strings.TrimSpace(text)
// }
// func punc(t string) string {
// 	// words := strings.Fields(t)
// 	// t = strings.Join(words, "")
// 	r := strings.NewReplacer(
// 		" ,", ",",
// 		" ;", ";",
// 		" '", "'",
// 		" .", ".",
// 		" *", "*",
// 	)
// 	return r.Replace(t)
// }
// func chunk(nums []int, size int) [][]int {
// 	if size <= 0 {
// 		return nil
// 	}
// 	var chunks [][]int
// 	for i := 0; i < len(nums); i += size {
// 		end := i + size
// 		if end > len(nums) {
// 			end = len(nums)
// 		}
// 		chunks = append(chunks, nums[i:end])
// 	}
// 	return chunks
// }
// func sum(nums ...int) int {
// 	var num int
// 	for _, n := range nums {
// 		num += n
// 	}
// 	return num
// }
// func isPalindrome(s string) bool {

// }

//	func main() {
//		fmt.Println(up("word"))
//		fmt.Println(low("WORD"))
//		fmt.Println(cap("word"))
//		fmt.Println(bin("1010"))
//		fmt.Println(hex("1E"))
//		fmt.Println(article("' awe '"))
//		fmt.Println(punc("word , available ."))
//		fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 2))
//		fmt.Println(sum(4,5,6))
//	}
func splitwords(s string) []string {
	ans := strings.Fields(s)
	// fmt.Printf("%q", ans)
	return ans
}
func fixArticles(s string) string {
	word := strings.Fields(s)
	 for i := 0; i < len(word); i++{
		if strings.ToLower(word[i]) == "a" && strings.ContainsRune("aeiouhAEIOUH", rune(word[i+1][0])){
			word[i] = "An" strings.ContainsAn()
		}
	 }

	 return strings.Join(word, " ")
}
func applyLow(s string) string {
	word := strings.Fields(s)

	for i := 1; i < len(word); i++{
		if word[i] == "(low)"{
			word[i-1] = strings.ToLower(word[i-1])
		}
	}
	var ans []string
	for _, letter := range word {
		if letter != "(low)"{
			ans = append(ans, letter)
		}
	}
	return strings.Join(ans, " ")
}	
func applyUp(s string) string {
	word := strings.Fields(s)

	for i := 1; i < len(word); i++{
		if word[i] == "(up)"{
			word[i-1] = strings.ToUpper(word[i-1])
		}
	}
	var ans []string
	for _, letter := range word{
		if letter != "(up)"{
			ans = append(ans, letter)
		}
	}
	return strings.Join(ans, " ")
}
func bin(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++{
		if word[i] == "(bin)"{
			num, err := strconv.ParseInt(word[i-1], 2, 64)
			if err == nil {
				word[i-1] = strconv.Itoa(int(num))
			}
		}
	}
	var ans []string
	for _, letter := range word {
		if letter != "(bin)"{
			ans = append(ans, letter)
		}
	}
	return strings.Join(ans, " ")
}
func CountVowels(s string) int {
	count := 0
	for _, vowels := range "aeiou"{
		count += strings.Count(s, string(vowels))
	}
	return count
}
func reverse(s string) string {
	ans := ""
	for i := len(s)-1; i >= 0; i--{
		ans = ans + string(s[i])
	}
	return ans
}
func ispalindrome(s string) bool {
	letter := []rune(s)
	for i,j := 0, len(letter)-1; i<j; i,j = i+1, j-1{
		if letter[i] != letter[j]{
			return false
		}
	}
	return true
}
func swapCase(s string) string {
	letter := []rune(s)
	var ans strings.Builder

	// for _, char := range letter{
	// 	if char >= 'a' && char <= 'z'{
	// 		ans += strings.ToUpper(string(char))
	// 	} else if char >= 'A' && char <= 'Z'{
	// 		ans += strings.ToLower(string(char))
	// 	} else {
	// 		ans += string(char)
	// 	}
	// }
	for _, char := range letter {
		switch {
		case char >= 'a' && char <= 'z':
			ans.WriteString(strings.ToUpper(string(char)))
		case char >= 'A' && char <= 'Z':
			ans.WriteString(strings.ToLower(string(char)))
		default :
			ans.WriteString(string(char))		
		}
	}
	return ans.String()
}
func countChars(s string) int {
	count := 0
	for range s{
		count++
	}
	return count
}
func convertHex(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++{
		if word[i] == "(hex)" {
			num, err := strconv.ParseInt(word[i-1], 16, 32)
			if err == nil{
				word[i-1] = strconv.Itoa(int(num))
			}
		}
	}
	var result []string
	for _, letter := range word{
		if letter != "(hex)"{
			result = append(result, letter)
		}
	}
	return strings.Join(result, " ")
}
func applyCap(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++{
		word[i] = strings.ToUpper(string(word[i][0])) + strings.ToLower((word[i][1:]))
	}
	return strings.Join(word, " ")
}
func lastNWords(words []string, n int) []string{
	start := len(words) - n
	for i := start; i < len(words); i++{
		words[i] = strings.ToUpper(words[i])
	}
	return words
}
func IsAnagram(a, b string) bool{
	if len(a) != len(b){
		return false
	}

	str1 := strings.Split(a, "")
	str2 := strings.Split(b, "")

	sort.Strings(str1)
	sort.Strings(str2)

	return slices.Equal(str1, str2)
}
//stripTags([]string{"hello","(hex)","30","(up)","world"}) -> ["hello","30","world"]
func stripTags(words []string) []string {
	result := []string{}
	for _, token := range words {
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")"){
			continue
		} else{
			result = append(result, token)
		}
	}
	return result
}
func main() {
	fmt.Printf("%q\n", splitwords("go is fun"))
	fmt.Println(fixArticles("A answer a hlan"))
	fmt.Println(applyLow("WORLD CLASS (low)"))
	fmt.Println(bin("1001 (bin)"))
	fmt.Println(CountVowels("aejifhonf"))
	fmt.Println(reverse("rahcar"))
	fmt.Println(swapCase("FFo"))
	fmt.Println(countChars("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(convertHex("1e (hex)"))
	fmt.Println(applyUp("hello there world (up)"))
	fmt.Println(ispalindrome("racecar"))
	fmt.Println(applyCap("we love recoding"))
	words := []string{"this", "is", "so", "nice"}
	result := lastNWords(words, 3)
	fmt.Println(result)
	fmt.Println(IsAnagram("listen", "silent"))
	fmt.Println(stripTags([]string{"hello","(hex)","30","(up)","world"}))
}
