
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64)  float64 {
    var z float64 = 1.0
    last := 0.0
    epsilon := 0.0001
    for ix := 0 ; ix < 10 ; ix ++ {
        z -= ( z * z - x) / (2 * z)
        //fmt.Println(z)
        if last != 0 {
           if math.Abs(last - z) < epsilon {
               //fmt.Println("Exiting early")
               return z
           }
        }
        last = z
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(3))
    fmt.Println(Sqrt(4))
    fmt.Println(Sqrt(5))
    fmt.Println(Sqrt(6))
    fmt.Println(Sqrt(2))
}
