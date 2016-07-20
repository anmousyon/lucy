package helpers

import(
    "time"
    "strconv"
    "math"
    "fmt"
    "database/sql"
    //database driver for sqlite3
    _ "github.com/mattn/go-sqlite3"
    //"github.com/wolfgangmeyers/go-rake/rake"
    "github.com/jzelinskie/geddit"
)


func createTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS post(label TEXT PRIMARY KEY, title TEXT, site TEXT, user TEXT, sub TEXT, created TEXT, edited TEXT, sentiment REAL, karma INTEGER, gold INTEGER);")
	if err != nil {
        panic(err)
    }
}

//FillDB with all posts
func FillDB() error{
    db, err := sql.Open("sqlite3", "lucy.db")
    if err != nil{
        fmt.Println("couldnt open")
        panic(err)
    }
    tx, err := db.Begin()
    if err != nil{
        fmt.Println("couldnt begin")
        panic(err)
    }
    createTable(db)
    posts := getPosts("python")
    for _, item := range posts{
        _, err := db.Exec(
            "Insert into post (label, title, site, sub, user, created, edited, sentiment, karma, gold) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
            item.ID,
            item.Title,
            item.URL,
            item.Author,
            item.Subreddit,
            strconv.Itoa(int(item.DateCreated)),
            strconv.Itoa(int(item.DateCreated)),
            0.0,
            round(item.Score),
            0,
        )
        if err != nil{
            fmt.Println("couldnt add")
            panic(err)
        }
    }
    tx.Commit()
    return nil
}

//GetPosts from database
func getPosts(sub string) []*geddit.Submission{
    session, err := geddit.NewLoginSession(
        "anmousyony",
        "buffalo12",
        "lucy",
    )
    if err != nil{
        fmt.Println("couldnt login")
        panic(err)
    }
    subOpts := geddit.ListingOptions{
        Limit: 10,
    }
    time.Sleep(2000 * time.Millisecond)
    submissions, _ := session.SubredditSubmissions(sub, geddit.NewSubmissions, subOpts)
    return submissions
}

//SimplifyBool for booleans in database
func SimplifyBool(boolean string) string{
    if boolean == "True"{
        return "1"
    }
    return "0"
}

//Hour of timestamp
func Hour(postTime string) string{
    intTime, err := strconv.Atoi(postTime)
    if err != nil{
        return  "0"
    }
    hour := time.Unix(int64(intTime), 0).Hour()
    strHour := strconv.Itoa(hour)
    return strHour
}

//Round a number to most significant digit
func round(num int) int{
    val := float64(num)
    size := float64(int(math.Floor(math.Log10(val))))
    power := math.Pow(10, size)
    rounded := math.Abs(float64(int(val / power)) * power)
    return int(rounded)
}

//Login to reddit
func Login() (*geddit.LoginSession, error){
    session, err := geddit.NewLoginSession(
        "anmousyony",
        "buffalo12",
        "lucy",
    )
    if err != nil{
        fmt.Println("error: ", err)
    }
    return session, err
}