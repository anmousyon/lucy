package main

import(
    "myServer"
    "mlHelpers"
    "classifier"
    "helpers"
)

func main(){
    err := helpers.FillDB()
    if err != nil{
        panic(err)
    }
    data := mlHelpers.GetData()
    encoded, classes, encoders := mlHelpers.Prep(data)
    for _ = range encoders{
        break
    }
    classifier.FitClassify(encoded, classes)
    myServer.StartServer()
}