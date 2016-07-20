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
        fmt.Println("error on create", err)
    }
}

//FillDB with all posts
func FillDB(){
    db, err := sql.Open("sqlite3", "lucy.db")
    if err != nil{
        fmt.Println("error on open", err)
    }
    tx, err := db.Begin()
    if err != nil{
        fmt.Println("error on begin", err)
    }
    createTable(db)
    posts := Posts("python")
    for _, i := range posts{
        _, err := db.Exec(
            "Insert into post (label, title, site, sub, user, created, edited, sentiment, karma, gold) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
            i.ID,
            i.Title,
            i.URL,
            i.Author,
            i.Subreddit,
            strconv.Itoa(int(i.DateCreated)),
            strconv.Itoa(int(i.DateCreated)),
            0.0,
            round(i.Score),
            0,
        )
        if err != nil{
            fmt.Println("error on insert", err)
        }
    }
    tx.Commit()
}

//Posts from database
func Posts(sub string) (posts []*geddit.Submission){
    s, err := Login()
    if err != nil{
        fmt.Println("error on login", err)
    }
    opts := geddit.ListingOptions{
        Limit: 10,
    }
    time.Sleep(2000 * time.Millisecond)
    posts, _ = s.SubredditSubmissions(sub, geddit.NewSubmissions, opts)
    return posts
}

//SimplifyBool for booleans in database
func SimplifyBool(b string) (simple string){
    if b == "True"{
        return "1"
    }
    return "0"
}

//Hour of timestamp
func Hour(ts string) (hours string){
    t, err := strconv.Atoi(ts)
    if err != nil{
        return  "0"
    }
    h := strconv.Itoa(time.Unix(int64(t), 0).Hour())
    return h
}

//Round a number to most significant digit
func round(x int) (rounded int){
    f := float64(x)
    sz := float64(int(math.Floor(math.Log10(f))))
    rad := math.Pow(10, sz)
    rnd := math.Abs(float64(int(f/rad))*rad)
    return int(rnd)
}

//Login to reddit
func Login() (session *geddit.LoginSession, err error){
    s, err := geddit.NewLoginSession(
        "anmousyony",
        "buffalo12",
        "lucy v0.01 for machine learning research by /u/anmousyony",
    )
    return s, err
}