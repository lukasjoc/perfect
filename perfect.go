package main

import (
    "fmt"
    // "errors"
)

/*
What is a perfect square?
---
Add up all rows and cols and diagonals
if the same number appears then it is a 
magic square
*/

//type Cell struct {
//    value uint64
//}
//
//type Row struct {
//    Data []Cell
//    Sum uint64
//    Size uint64
//}
//
//type Square struct {
//    Shape [][]Row
//    Size uint64
//    Limit uint64
//}
//
func nextPerm(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPerm(orig, p []int) []int {
    result := append([]int{}, orig...)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}
//
//// generate a new square
//func (s *Square) generate() error {
//    if s.Size <= 0 {
//        return errors.New("Please provide with size")
//    }
//    if s.Limit <= 0 {
//        return errors.New("Please provide with number limit")
//    }
//
//    // TODO: Generate Rows based of Size and Limit
//    // Size: 2
//    //0, 0
//    //0, 0
//
//    // How to sum a magic Square:
//    //diag 00 + 11 = ??
//    //diag 01 + 00 = ??
//    //col1 00 + 01 = ??
//    //col2 10 + 11 = ??
//    //row1 00 + 11 = ??
//    var uint64 i
//    for i = 0 ; i <= s.Limit; i++ {
//    }
//    return nil
//}

// TODO: need to find all permutations
// then group them together into groups of permutations
func main() {
    //sq := &Square{
    //    Size: 2,
    //    Limit: 2,
    //}

    // sq.generate()

    orig := []int{0, 1, 2, 3}
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
        fmt.Println(getPerm(orig, p))
    }

    // fmt.Println(sq)
}
