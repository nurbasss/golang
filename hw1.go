package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(Anagram("HUu","uHhu"))
	fmt.Println(FindEvens([]int{2,5,5,8,9,12,1,0}))
	fmt.Println(SliceProduct([]int{1, 2, 3}))
	fmt.Println(Unique([]int{1,1,2,2,3,5,4,2}))
	fmt.Println(InvertMap(map[string]int{
		"rast": 1,
		"cs": 3,
		"fort": 2,
	}))
	fmt.Println(TopCharacters("go makes me cry", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
        log.Fatal(err)
    }
    nums := reg.ReplaceAllString(phone, "")
	res := "("+nums[0:3]+") " + nums[3:6] + "-" + nums[6:10]
	return res
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	arr1 := strings.Split(strings.ToLower(s1), "")
	arr2 := strings.Split(strings.ToLower(s2), "")
	sort.Strings(arr1)
	sort.Strings(arr2)
	
	return strings.Join(arr1,"")==strings.Join(arr2,"")
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var res []int
	for i := 0; i < len(e); i++ {
		if e[i] % 2 == 0 {
			res = append(res, e[i])
		}
	}	
	return res
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	res := 1
	for _, i := range e {
		res *= i
	}
	return res
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	uniqueNums := make(map[int]bool)
    res := []int{}	
    for _, num := range e {
        if _, value := uniqueNums[num]; !value {
            uniqueNums[num] = true
            res = append(res, num)
        }
    }    
	return res
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	res := make(map[int]string)
	for key, value := range kv {
		res[value] = key
	}
	return res
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	occur := make(map[rune]int)
	for _, r := range s {
		occur[r]++
	}
	for r, occ := range occur {
		if occ <= k {
			delete(occur, r)
		}
	}
	return occur
}