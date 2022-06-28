package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    // Read from OS Standard Input Stream
    stdin := bufio.NewReader(os.Stdin)

    var a int
    var b int

    n, err := fmt.Scanln(&a, &b)
    if err != nil {
        // When Error occurs the remaining stream data is left in stdin
        fmt.Println(err)
        // Find a new line to read until
        // Read stream data until new line
        stdin.ReadString('\n')
    } else {
        fmt.Println(n, a, b)
    }

    n, err = fmt.Scanln(&a, &b)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(n, a, b)
    }
}
