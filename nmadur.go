package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
    slc := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    sl := []string{}
    count := 0
    for _, i := range words{
        a := ""
        for _, j := range i{
            n := int(j) -97
            a += slc[n]
        }
        for j := 0; j < len(sl); j++{
            if sl[i] == a{
                count++
            }
        }
        sl = append(sl, a)
    }
    return count
}
func main() {
	a := []string{"adf", "adf", "asd"}
	fmt.Println(uniqueMorseRepresentations(a))
}