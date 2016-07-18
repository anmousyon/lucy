package mlHelpers

import(
    "fmt"
)

func extractString(data [][]string, col int) ([]string){
    column := make([]string, len(data))
    for row := range data{
        column[row] = data[row][col]
    }
    return column
}

func extractInt(encoded [][]int, col int) ([]int){
    column := make([]int, len(encoded))
    for row := range encoded{
        column[row] = encoded[row][col]
    }
    return column
}

func combine(cols [][]int) [][]int{
    data := make([][]int, len(cols[0]))
    for row := range cols[0]{
        dataRow := make([]int, len(cols))
        for col := range cols{
            dataRow[col] = cols[col][row]
        }
        data[row] = dataRow
    }
    return data
}
/*
func getData() [][]string{
    db, err := gorm.Open("sqlite3", "lucy.db")
    comments := db.Find(&Comment)
    data := make([]string, len(comments))
    for comment := range comments{
        row := [...]string{
            helpers.terms(comment.body),
            comment.post,
            comment.user,
            comment.sub,
            helpers.hour(comment.created),
            helpers.hour(comment.edited),
            helpers.simpleSent(comment.sentiment),
            helpers.rounder(comment.karma),
            helpers.simpleBool(comment.gold),
        }
        data = append(row)
    }
    return data
}
*/
func encode(col []string) ([]int, map[string]int){
    unique := map[string]int{}
    encoded := make([]int, len(col), len(col))
    counter := 0
    for _, x := range col{
        match := false
        for y := range unique{
            if x == y{
                match = true
                break
            }
        }
        if !match{
            unique[x] = counter
            counter++
        }
    }
    for index, x := range col{
        for y, z := range unique{
            if x == y{
                encoded[index] = z
                break
            }
        }
    }
    return encoded, unique
}

//Prep the data for golearn
func Prep(data [][]string) ([][]int, []int, []map[string]int){
    cols := make([][]int, len(data[0]))
    encoders := make([]map[string]int, len(data[0]))
    for x := range data[0]{
        col := extractString(data, x)
        encoded, encoder := encode(col)
        cols[x] = encoded
        encoders[x] = encoder
    }
    encodedData := combine(cols)
    for x := range encodedData{
        for y := range encodedData[0]{
            fmt.Print(encodedData[x][y])
        }
        fmt.Println("")
    }
    classes := extractInt(encodedData, 2)
    return encodedData, classes, encoders
}