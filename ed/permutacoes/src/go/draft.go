package main

import (
    "fmt"
    "sort"
    "strings"
)        

func backtrack(chars []string, visited []bool, current []string) {
    
    if len(current) == len(chars) {
        fmt.Println(strings.Join(current, ""))
        return
    }

    for i := 0; i < len(chars); i++ {
    
        if visited[i] {
            continue
        }
        
        visited[i] = true
        current = append(current, chars[i])

        backtrack(chars, visited, current)

        current = current[:len(current)-1]
        visited[i] = false
    }
}

func main() {
    
    var s string

    if _, err := fmt.Scan(&s); err != nil {
        return 
    }

    chars := strings.Split (s, "")
    
    sort.Strings(chars)
    
    visited := make([]bool, len(chars))
    
    var current []string
    
    backtrack(chars, visited, current)
}