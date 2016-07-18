package jude

import(
    "database/sql"
    "github.com/mattn/go-sqlite3"
    "github.com/jinzhu/gorm"
)

func extract(data [][]string, col int) [][]string, []string{
    column := make(string, len(data))
    for row := range len(data){
        column = append(data[row][col])
        data[row] = append(data[row][:col], data[row][col+1:]...)
    }
    return data, column
}

func combine(cols [][]string) [][]string{
    data = make([]string, len(cols[0]))
    for row := range len(cols){
        dataRow = make(string, len(cols))
        for col := range len(cols[0]){
            dataRow = append(cols[col][row])
        }
        data = append(dataRow)
    }
    return data
}

func getData() [][]string{
    db, err := gorm.Open("sqlite3", "lucy.db")
    comments := db.Find(&comments)
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
            helpers.simpleBool(comment.gold)
        }
        data = append(row)
    }
    return data
}

func encode(col []string) []int, map[string]int{}{
    unique := map[string]int{}
    encoded := []int{}
    counter := 0
    for x := range col{
        for y, _ := range unique{
            if x == y{
                break
            }
        }
        unique[x] = counter
        counter++
    }
    for x := range col{
        for y, z := range unique{
            if x == y{
                encoded.append(z)
                break
            }
        }
    }
    return encoded, unique
}

func prep([][]string data) [][]string, []string, map[string]int{}{
    cols := [][]string
    encoders := map[string]int{}
    for x := range len(data[0]){
        data, col = extract(data, 0))
        col, encoder = encode(col)
        cols.append(col)
        encoders.append(encoder)
    }
    data = combine(cols)
    data, classes = extract(data, 6)
    return data, classes, encoders
}