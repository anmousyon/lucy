package mlHelpers

import (
    "fmt"
    "database/sql"
    //database driver for sqlite3
    _ "github.com/mattn/go-sqlite3"
    "strconv"
    "github.com/gonum/matrix/mat64"
)

//GetData from the database
func GetData() [][]string{
    data := [][]string{}
    db, err := sql.Open("sqlite3", "lucy.db")
    rows, err := db.Query("select * from post")
    if err != nil{
        panic(err)
    }
    defer db.Close()
    for rows.Next(){
        var label string
        var title string
        var site string
        var user string
        var sub string
        var created string
        var edited string
        var sentiment float64
        var karma int
        var gold int
        err := rows.Scan(&label, &title, &site, &user, &sub, &created, &edited, &sentiment, &karma, &gold)
        if err != nil {
		    panic(err)
	    }
        sentimentString :=strconv.FormatFloat(sentiment, 'E', -1, 64)
        karmaString := strconv.Itoa(karma)
        goldString := strconv.Itoa(gold)
        row := []string{label, title, site, user, sub, created, edited, sentimentString, karmaString, goldString}
        data = append(data, row)
    }
    fmt.Println(data)
    return data
}

func extractString(data [][]string, col int) ([]string){
    column := make([]string, len(data))
    for row := range data{
        column[row] = data[row][col]
    }
    return column
}

//ExtractEncoded column from a slice of slices
func ExtractEncoded(encoded [][]float64, col int) ([]float64){
    column := make([]float64, len(encoded))
    for row := range encoded{
        column[row] = encoded[row][col]
    }
    return column
}

func combine(columns [][]float64) *mat64.Dense{
    dataMatrix := mat64.NewDense(len(columns[0]), len(columns), nil)
    for row := range columns[0]{
        for column := range columns{
            dataMatrix.Set(column, row, columns[column][row])
        }
    }
    return dataMatrix
}

func encode(column []string) ([]float64, map[string]float64){
    unique := map[string]float64{}
    encoded := make([]float64, len(column), len(column))
    counter := 0.0
    for _, item := range column{
        match := false
        for value := range unique{
            if item == value{
                match = true
                break
            }
        }
        if !match{
            unique[item] = counter
            counter++
        }
    }
    for index, item := range column{
        for key, value := range unique{
            if item == key{
                encoded[index] = value
                break
            }
        }
    }
    return encoded, unique
}

//Prep the data for golearn
func Prep(data [][]string) (*mat64.Dense, []float64, []map[string]float64){
    columns := make([][]float64, len(data[0]))
    encoders := make([]map[string]float64, len(data[0]))
    for col := range data[0]{
        column := extractString(data, col)
        columns[col], encoders[col] = encode(column)
    }
    encodedData := combine(columns)
    classes := make([]float64, len(columns[0]))
    mat64.Col(classes, 6, encodedData)
    return encodedData, classes, encoders
}