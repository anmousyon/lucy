package main

import(
    //"myServer"
    "mlHelpers"
    "fmt"
)

func main(){
    data := make([][]string, 5)
    data [0] = []string{"a", "b", "a", "d"}
    data [1] = []string{"b", "c", "b", "e"}
    data [2] = []string{"c", "b", "a", "f"}
    data [3] = []string{"d", "e", "a", "g"}
    data [4] = []string{"a", "f", "b", "h"}
    for x := range data{
        for y := range data[0]{
            fmt.Print(data[x][y])
        }
        fmt.Println("")
    }
    encoded, classes, encoders := mlHelpers.Prep(data)
    for row := range encoded{
        for col := range encoded[0]{
            encoded[row][col] = encoded[row][col]
        }
    }
    for class := range encoders{
        classes[class] = classes[class]
    }
    //myServer.StartServer()
}