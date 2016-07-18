package helpers

import(
    "myServer"
    "strconv"
    //"github.com/wolfgangmeyers/go-rake/rake"
    "github.com/jinzhu/gorm"
    "github.com/jzelinskie/geddit"
)

func GetPosts(reddit geddit.LoginSession, sub string) []geddit.Submission{
    db, err := gorm.Open("sqlite3", "lucy.db")
    subOpts := geddit.ListingOptions{
        Limit: 10,
    }
    posts, _ := reddit.SubredditSubmissions(sub, geddit.HotSubmissions, subOpts)
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
    return posts
}

func getComments(reddit geddit.LoginSession, post geddit.Submission) []geddit.Comment{
    comments, _ := reddit.Comments(post)
    return comments
}

func simpleBool(boolean string) string{
    val := " "
    if boolean == "True"{
        return "1"
    } else{
        return "0"
    }
}

func simpleSent(sentiment string) string {
    FILTER := rake("stoplist.txt")
    sent, err := strconv.ParseFloat(sentiment, 64)
    if sent > 0 {
        return "1"
    } else if sent == 0 {
        return "0"
    } else {
        return "-1"
    }
}

func hour(time string) string{
    intTime, err := strconv.Atoi(time)
    if err != nil{
        return  "0"
    }
    hour := time.Unix(intTime, 0).Hour()
    return hour
}

func rounder(num int) string{
    if num == 0{
        return '0'
    } else {
        rounded := round(num, -int(fllor(log10(abs(int(num))))))
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

func login() geddit.LoginSession{
    session, _ := geddit.NewLoginSession(
        "anmousyony",
        "buffalo12",
        "lucy",
    )
    return session
}