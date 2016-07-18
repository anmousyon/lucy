package classifier

import (
    "fmt"
    "github.com/sjwhitworth/golearn/base"
    "github.com/sjwhitworth/golearn/trees"
)

func fit(data [][]string){
    data, classes, encoders := mlHelpers.prep(data)
    trainData, testData := base.InstancesTrainTestSplit(data, 0.10)
    classifier := trees.id3()
    trainData, trainClasses := mlHelpers.extract(trainData, 6)
    testData, testClasses := mlHelpers.extract(testData, 6)
    classifier.fit(trainData, trainClasses)
    classify(encoders, testData, testClasses)
}

func classify(encoders []struct, test [][]string, testClasses []string){
    prediction := cls.Predict(testData)
    fmt.print("prediction")
    for item := range prediction{
        fmt.print(item)
    }
    fmt.print("actual")
    for item := range testData{
        fmt.print(item)
    }
}

func train(){
    data := mlHelpers.getData()
    fit(data)
}