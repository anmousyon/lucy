package main

import(
    "myServer"
    "mlHelpers"
    "classifier"
    "helpers"
)

func main(){
    helpers.FillDB()
    d := mlHelpers.Data()
    enc, cls, encs := mlHelpers.Prep(d)
    for _ = range encs{
        break
    }
    classifier.FitClassify(enc, cls)
    myServer.StartServer()
}