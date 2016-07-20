package classifier

import (
    "fmt"
    "github.com/sjwhitworth/golearn/base"
    "github.com/sjwhitworth/golearn/trees"
	"github.com/sjwhitworth/golearn/evaluation"
    "github.com/gonum/matrix/mat64"
)

//FitClassify the data and print prediction
func FitClassify(encoded *mat64.Dense, classes []float64){
    readyData := base.InstancesFromMat64(len(classes), 10, encoded)
    trainData, testData := base.InstancesTrainTestSplit(readyData, 0.10)
    tree := trees.NewID3DecisionTree(0.6)
    err := tree.Fit(trainData)
    if err != nil{
        panic(err)
    }
    predictions, err := tree.Predict(testData)
    if err != nil{
        panic(err)
    }
	cf, err := evaluation.GetConfusionMatrix(testData, predictions)
	fmt.Println(evaluation.GetSummary(cf))
}