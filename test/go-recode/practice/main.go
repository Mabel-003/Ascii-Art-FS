package main

import (
	"fmt"
	"sort"
	"strconv"

	// "strconv"
	"strings"

	// "unicode"
	"slices"
)

// "strings"
// "strings"

// func isAnagram(s1, s2 string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}

// 	str1 := strings.Split(strings.ToLower(s1), "")
// 	str2 := strings.Split(strings.ToLower(s2), "")

// 	sort.Strings(str1)
// 	sort.Strings(str2)

// 	fmt.Println(str1)
// 	fmt.Println(str2)

// 	return slices.Equal(str1, str2)
// }

// func main() {
// 	fmt.Println(isAnagram("sShe", "hses"))
// }

// func chunk(nums []int, size int) [][]int {

// 	counter := 0
// 	var carry []int

// 	var result [][]int

// 	for _, number := range nums {
// 		if counter < size-1 {
// 			carry = append(carry, number)
// 			counter++
// 		} else if counter == size-1 {
// 			carry = append(carry, number)
// 			result = append(result, carry)
// 			counter = 0
// 			carry = []int{}
// 		}
// 	}

// 	return result
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

// func main() {
// 	fmt.Println(chunk([]int{1, 2, 3, 4, 5}, 2))
// }

//  func sum(nums ...int) int {
// 	r := 0
// 	for _, s := range nums{
// 		r += s
// 	}
// 	return r
// }

// func main(){
// 	 nums := []int{4,5,6}
// 	fmt.Println(sum(nums...))
// }

//  stripTags([]string{"hello","(hex)","30","(up)","world"}) -> ["hello","30","world"]
// func stripTags(words []string) []string {
//  result := []string{}
//  for _, token := range words {
// 		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")" ) {
// 			continue
// 		} else {
// 			result = append(result, token)
// 		}
//  }
//  return result

// }
// func main(){

//		fmt.Println(stripTags([]string{"hello","(hex)","30","(up)","world"}))
//	}
// func fixEllipsis(s string) string {
// 	fmt.Println(strings.ReplaceAll(s, " ...", "..."))
// 	return ""
// }
// func main() {
// 	fmt.Println(fixEllipsis("wait ..."))
// 	fmt.Println(splitPunctuation("hello,"))

// 	array := []string{"Hello", ","}
// 	fmt.Printf("%q\n", array)
// }

// func splitPunctuation(s string) []string {
// 	result := strings.Fields(strings.ReplaceAll(s, ",", " ,"))

// 	for index, word := range result {
// 		result[index] = "\"" + word + "\""
// 	}

// func applyCap(s string) string {
// 	words := strings.Fields(s)

// 	for i := 0; i < len(words); i++ {
// 		word := words[i]
// 		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
// 	}
// 	return strings.Join(words, " ")
// }
// func main(){
// 	fmt.Println(applyCap("we love recoding"))
// }

// caesarCipher("abc", 3)   -> "def"
// func caesarCipher(s string, shift int) string {

// }
//

// func anagram(w1,w2 string) bool{
// 	if len(w1) != len(w2) {
// 		return false
// 	}

// 	str1 := strings.Split(w1, "")
// 	str2 := strings.Split(w2, "")

// 	sort.Strings(str1)
// 	sort.Strings(str2)

//		return slices.Equal(str1, str2)
//	}
//
//	func main(){
//		fmt.Println(anagram("her", "rhe"))
//		fmt.Println(anagram("left", "heft"))
//	}
func printChar(s string) {
	for i := 0; i <= len(s)-1; i++{
		fmt.Println(string(s[i]))
	}
}
func countChars(s string) int {
	count := 0

	for i := 0; i <= len(s)-1; i++{

		count++
	}
	return count
	//return strings.Count(s, "l")
}
func toUpperCase(s string) {
	fmt.Println(strings.ToUpper(s))
}
func toLowerCase(s string) {
	fmt.Println(strings.ToLower(s))
}
func stri(s []string) []string {
	result := make([]string, len(s))
	for i, str := range s {
		result[i] = "\"" + str + "\""
	}
	return result
}
func reverse(s string) string {
	// ans := ""
	// for i := len(s)-1; i >= 0; i-- {
	// 	ans = ans + string(s[i])
	// }
	// return ans
	
	
	words := []rune(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return string(words)
}
func ispalindrome(s string) bool {
	words := []rune(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		if words[i] != words[j]{
			return false
		}
	}
	return true
}
func swapCase(s string) string {
	letter := []rune(s)
	ans := ""
	for _, char := range letter {
		if char >= 'a' && char <= 'z'{
			ans += strings.ToUpper(string(char))
		} else if char >= 'A' && char <= 'Z'{
			ans += strings.ToLower(string(char))
		} else {
			ans += string(char)
		}
	}
	return ans
}
func CountVowels(s string) int {
	count := 0

	s = strings.ToLower(s)
	for _, vowel := range "aeiou"{
		count += strings.Count(s, string(vowel))
	}
	return count
	// for _, r := range s {
	// 	if unicode.ToLower(r) == 'a'{
	// 		count++
	// 	} else if unicode.ToLower(r) == 'e'{
	// 		count++
	// 	} else if unicode.ToLower(r) == 'i'{
	// 		count++
	// 	}else if unicode.ToLower(r) == 'o'{
	// 		count++
	// 	}else if unicode.ToLower(r) == 'u'{
	// 		count++
	// 	// } else {
	// 	// 	count += int(unicode.ToLower(r))
	// 	 }
	// }
	// return count
	// for _, r := range s{
	// 	if unicode.ToLower(r) >= 'a' && unicode.ToLower(r) <= 'z'{
	// 		 count ++ //int(unicode.ToLower(r))
	// 	}
	// }
	// return count
}
func RemoveSpaces(s string) string{
	// return strings.ReplaceAll(s, " ", "")
	// letter := []rune(s)
	// for i, ch := range letter{
	// 	if ch == ' ' {
	// 		letter = append(letter[:i],letter[i+1:]... )
			
	// 	}
	// }

	// return string(letter)


	ans := ""

	for _,ch := range s {

		if ch != ' ' && ch != ' ' {
			ans += string(ch)
		} 
	}
	return ans


}
func FirstIndex(s string, ch rune) int{
	return strings.Index(s, string(ch))
	// for index, r := range s {
	// 	if r == ch {
	// 		return index
	// 	}
	// }
	// return int(ch)
}
func CountWords(s string) int{
	count := 0
	words := strings.Fields(s)

	for range words{
		count++
	}
	return count
}

func IsAnagram(a, b string) bool{
	if len(a) != len(b){
		return false
	}
	str1 := strings.Split(a, "")
	str2 := strings.Split(b, "")

	sort.Strings(str1)
	sort.Strings(str2)

	return slices.Equal(str2, str1)
}
func Join(a, b string) string{
	str1 := strings.Fields(a)
	str2 := append(str1, b)
	return strings.Join(str2, "")
}

func ReplaceChar(s string, old, new rune) string{
	//return strings.ReplaceAll(s, string(old), string(new))

	ans := ""
	for _, r := range s{
		if r == old{
			ans += string(new)
		} else {
			ans += string(r)
		}
	}
	return ans
}
func Duplicates(s string){
	// Create a map to store the frequency of each character
	counts := make(map[rune]int)

	// Count occurrences of each character
	for _, char := range s {
		counts[char]++
	}

	// fmt.Printf("Duplicate characters in '%s':\n", s)
	 found := false

	 // Iterate through the map and print characters with count > 1
	 for char, count := range counts {
	 	if count > 1 {
	 		fmt.Printf("'%c' appears %d times\n", char, count)
	 		found = true
	 	}
	 }

	 if !found {
	 	fmt.Println("No duplicates found.")
	 }
}
func splitWords(s string) []string{
	return strings.Fields(s)
}
func fixArticles(s string) string{
	words := strings.Fields(s)
	for i := 0; i < len(words); i++{
		if strings.ToLower(words[i]) == "a" && strings.ContainsRune("aeiouhAEIOUH", rune(words[i+1][0])){
			words[i] = "An"
		}
	}
	return strings.Join(words, " ")
}    
func returnNum(s string, n int) string{
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}
// func applyWordN(s string) []string{
// 	word := strings.Fields(s)
// 	ans := []string{}
// 	for i := 0; i < len(word); i++{
// 		if word[i] == "(up" {
// 			ans = ans + strings.ToUpper(word[i])
// 		} 
// 	} 
// 	return ans
// }
func applyUp(s string) string {
	words := strings.Fields(s)

	for i := 1; i < len(words); i++{
		if words[i] == "(up)"{
			words[i-1] = strings.ToUpper(words[i-1])
		}
	}

	var result []string
	for _, word := range words{
	if word != "(up)"{
		result = append(result, word)
	}
}
	return strings.Join(result, " ")
}
func punc(t string) string {
	r := strings.NewReplacer(
		" ,", ",",
		" .", ",",
	)
	return r.Replace(t)
}
func applyLow(s string) string{
	words := strings.Fields(s)

	for i := 1; i < len(words); i++{
		if words[i] == "(low)"{
			words[i-1] = strings.ToLower(words[i-1])
		}
	}
	
	var ans []string
	for _, word := range words{
		if word != "(low)"{
			ans = append(ans, word)
		}
	}
	return strings.Join(ans, " ")
}
func quotes(s string) string{
	r := strings.NewReplacer(
		"' ", "'",
		" '", "'",
	)
	return r.Replace(s)
}
// func hex(s string) string{
// 	word, _ := strconv.ParseInt(s, 16, 64)
// 	return word
// }
func applyCap(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++{
		// word := words[i]
		words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
	}
	return strings.Join(words, " ")
}
func convertHex(s string) string {
	words := strings.Fields(s)

	for i := 1; i < len(words); i++{
		if words[i] == "(hex)"{
			num, err := strconv.ParseInt(words[i-1], 16, 32) 
				if err == nil{
					words[i-1] = strconv.Itoa(int(num))
				}
		}
	}
	var result []string

	for _, w := range words{
		if w != "(hex)"{
			result = append(result, w)
		}
	}
	return strings.Join(result, "")
}
func convertBin(s string) string {
	words := strings.Fields(s)
	for i := 1; i < len(words); i++{
		if words[i] == "(bin)"{
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil{
				words[i-1] = strconv.Itoa(int(num))
			}
		}
	}
	var result []string
	for _, w := range words {
		if w != "(bin)"{
			result = append(result, w)
		}
	}

	return strings.Join(result, " ")
}
func lastNWords(words []string, n int) []string{
	start := len(words) - n
	for i := start; i < len(words); i++{
		words[i] = strings.ToUpper(words[i])
	}
	return words
}
func main() {
	printChar("hello")
	fmt.Println(countChars("hello"))
	toUpperCase("hello")
	toLowerCase("HELLO")
	s := []string{"go", "is", "fun"}
	fmt.Println(stri(s))
	fmt.Println(reverse("hello"))
	fmt.Println(ispalindrome("amala"))
	fmt.Println(swapCase("!@GoLang12345"))
	fmt.Println(CountVowels("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(RemoveSpaces("go is fun"))
	fmt.Println(FirstIndex("golang", 'l'))
	fmt.Println(CountWords("go is very fast"))
	fmt.Println(IsAnagram("lsten", "silent"))
	fmt.Println(Join("go", "lang"))
	fmt.Println(ReplaceChar("banana", 'r', 'y'))
	Duplicates("programming")
	fmt.Printf("%q\n", splitWords("go is fun"))
	fmt.Println(fixArticles("There it was. A amazing rock. A honest man. A book."))
	fmt.Println(returnNum("Jesus", 5))
// 	fmt.Println(applyWordN("this is nice (up , 2)"))
	fmt.Println(applyUp("hello there world (up)"))
	fmt.Println(punc("so , what now ."))
	fmt.Println(applyLow("hello there WORLD (low)"))
	fmt.Println(quotes("' snd '"))
	fmt.Println(applyCap("we love recoding"))
	fmt.Println(convertHex("1e (hex)"))
	fmt.Println(convertBin("i love 1001 (bin)"))
	words := []string{"this", "is", "so", "nice"}
	result := lastNWords(words, 3)
	fmt.Println(result)
}