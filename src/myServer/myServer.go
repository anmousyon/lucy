package myServer

import(
    "io"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/jzelinskie/geddit"
    "strconv"
    "time"
    "bufio"
    "os"
    "fmt"
)

type post struct{
    gorm.Model
    label string
    title string
    site string
    user string
    sub string
    created string
    edited string
    sentiment float64
    karma int
    gold int
}

type comment struct{
    gorm.Model
    label string
    body string
    postid string
    user string
    sub string
    created string
    edited string
    sentiment float64
    karma int
    gold int
}

func hello(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello world!")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func StartServer(){
    db, _ := gorm.Open("sqlite3", "lucy.db")

    posts := getPosts("python")
    fmt.Println("got all posts")
    for _, item := range posts{
        fmt.Println("adding post")
        newPost := post{
            label: item.ID,
            title: item.Title,
            site: item.URL,
            user: item.Author,
            sub: item.Subreddit,
            created: strconv.FormatFloat(float64(item.DateCreated), 'E', -1, 64),
            edited: strconv.FormatFloat(float64(item.DateCreated), 'E', -1, 32),
            sentiment: 0.0,
            karma: item.Score,
            gold: 0,
        }
        db.NewRecord(newPost)
    }

    fmt.Println("starting server")
    server := http.Server{
        Addr: ":8000",
        Handler: &myHandler{},
    }
    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = hello
    server.ListenAndServe()
}

type myHandler struct{}

func(*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if h, ok := mux[r.URL.String()]; ok{
        h(w, r)
        return
    }

    io.WriteString(w, "My server: "+r.URL.String())
}

func getPosts(sub string) []*geddit.Submission{
    session, err := geddit.NewLoginSession(
        "anmousyony",
        "buffalo12",
        "lucyprogram",
    )

    if err != nil{
        fmt.Println(err)
    }

    subOpts := geddit.ListingOptions{
        Limit: 10,
    }
    fmt.Println("get posts?")

    submissions, _ := session.DefaultFrontpage(geddit.DefaultPopularity, subOpts)
    submissions, _ = session.SubredditSubmissions("python", geddit.NewSubmissions, subOpts)
    /*
    dbPosts := db.Find(&Post)
    var ids []string
    for _, post := range dbPosts{
        ids.append(post.ID)
    }
    for i, post := range posts{
        for j, id := range ids{
            posts.remove(post)
        }
    }
    */
    fmt.Println("got posts")
    return submissions
}
/*
func getComments(reddit geddit.LoginSession, post geddit.Submission) []*geddit.Comment{
    comments, _ := reddit.Comments(post)
    return comments
}
*/
func simpleBool(boolean string) string{
    if boolean == "True"{
        return "1"
    } else{
        return "0"
    }
}
/*
func simpleSent(sentiment string) string {
    FILTER := rake.RAKE("stoplist.txt")
    sent, err := strconv.ParseFloat(sentiment, 64)
    if sent > 0 {
        return "1"
    } else if sent == 0 {
        return "0"
    } else {
        return "-1"
    }
}
*/
func hour(postTime string) string{
    intTime, err := strconv.Atoi(postTime)
    if err != nil{
        return  "0"
    }
    hour := time.Unix(int64(intTime), 0).Hour()
    strHour := strconv.Itoa(hour)
    return strHour
}
/*
func rounder(num int) string{
    if num == 0{
        return "0"
    } else {
        rounded := math.Round(num, -int(math.Floor(math.Log10(math.Abs(int(num))))))
        strRounded := strconv.Itoa(rounded)
        return strRounded
    }
}

func terms(text string) string{
    text = FILTER.run(text)
    if test != nil{
        return text[0][0]
    } else {
        return ' '
    }
}
*/
func login() *geddit.LoginSession{
    session, _ := geddit.NewLoginSession(
        "anmousyony",
        "buffalo12",
        "lucy",
    )
    return session
}