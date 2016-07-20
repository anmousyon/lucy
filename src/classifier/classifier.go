package classifier

import (
    "fmt"
    "github.com/sjwhitworth/golearn/base"
    "github.com/sjwhitworth/golearn/trees"
	"github.com/sjwhitworth/golearn/evaluation"
    "github.com/gonum/matrix/mat64"
)

//FitClassify the data and print prediction
func FitClassify(enc *mat64.Dense, cls []float64){
    inst := base.InstancesFromMat64(len(cls), 10, enc)
    dTrain, dTest := base.InstancesTrainTestSplit(inst, 0.10)
    tree := trees.NewID3DecisionTree(0.6)
    err := tree.Fit(dTrain)
    if err != nil{
        fmt.Println("fit error: ", err)
    }
    pred, err := tree.Predict(dTest)
    if err != nil{
        fmt.Println("prediction error: ", err)
    }
	cf, err := evaluation.GetConfusionMatrix(dTest, pred)
	fmt.Println(evaluation.GetSummary(cf))
}