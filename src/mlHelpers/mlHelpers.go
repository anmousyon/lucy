package mlHelpers

import (
    "fmt"
    "database/sql"
    //database driver for sqlite3
    _ "github.com/mattn/go-sqlite3"
    "strconv"
    "github.com/gonum/matrix/mat64"
)

//Data from the database
func Data() (data [][]string){
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
        var sent float64
        var karma int
        var gold int
        err := rows.Scan(&label, &title, &site, &user, &sub, &created, &edited, &sent, &karma, &gold)
        if err != nil {
		    panic(err)
	    }
        sentStr :=strconv.FormatFloat(sent, 'E', -1, 64)
        karmaStr := strconv.Itoa(karma)
        goldStr := strconv.Itoa(gold)
        row := []string{label, title, site, user, sub, created, edited, sentStr, karmaStr, goldStr}
        data = append(data, row)
    }
    fmt.Println(data)
    return data
}

func extractString(data [][]string, i int) (column []string){
    col := make([]string, len(data))
    for row := range data{
        col[row] = data[row][i]
    }
    return col
}

//ExtractEncoded column from a slice of slices
func ExtractEncoded(enc [][]float64, i int) (column []float64){
    col := make([]float64, len(enc))
    for row := range enc{
        col[row] = enc[row][i]
    }
    return col
}

func combine(cols [][]float64) (matrix *mat64.Dense){
    m := mat64.NewDense(len(cols[0]), len(cols), nil)
    for row := range cols[0]{
        for col := range cols{
            m.Set(col, row, cols[col][row])
        }
    }
    return m
}

func encode(col []string) (encoded []float64, encoders map[string]float64){
    unq := map[string]float64{}
    enc := make([]float64, len(col), len(col))
    n := 0.0
    for _, i := range col{
        match := false
        for v := range unq{
            if i == v{
                match = true
                break
            }
        }
        if !match{
            unq[i] = n
            n++
        }
    }
    for x, i := range col{
        for k, v := range unq{
            if i == k{
                enc[x] = v
                break
            }
        }
    }
    return enc, unq
}

//Prep the data for golearn
func Prep(data [][]string) (encoded *mat64.Dense, classes []float64, encoders []map[string]float64){
    cols := make([][]float64, len(data[0]))
    encs := make([]map[string]float64, len(data[0]))
    for x := range data[0]{
        col := extractString(data, x)
        cols[x], encs[x] = encode(col)
    }
    enc := combine(cols)
    cls := make([]float64, len(cols[0]))
    mat64.Col(cls, 6, enc)
    return enc, cls, encs
}