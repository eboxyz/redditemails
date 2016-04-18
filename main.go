package main

import(
    "log"
    "net/http"
    "fmt"
    "encoding/json"
    "errors"
    // "github.com/sendgrid/sendgrid-go"
    // "bytes"
)


//func main to send out scheduled emails
// func main() {
//
//     sg := sendgrid.NewSendGridClient("eboxyz", "SENDGRID_API_KEY")
//     message := sendgrid.NewMail()
//
//     message.AddTo("ed.got.a.gun@gmail.com")
//     message.AddToName("Edward")
//     message.AddSubject("Daily League of Leggos")
//     message.AddFrom("eboxyz@sendgrid.com")
//
//     message.AddHTML(Email())
//
//     if rep := sg.Send(message); rep == nil{
//         fmt.Println("Email sent!")
//         fmt Println("Closing...")
//     } else {
//         fmt.Println(rep)
//     }
// }

//func main to grab reddit titles
func main() {

    items, err := Get("leagueoflegends")
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range items {
        fmt.Println(item.Title)
    }

}

func Get(subreddit string)([]Item, error){
    url := fmt.Sprintf("http://reddit.com/r/%s.json", subreddit)
    r, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    defer r.Body.Close()
    if r.StatusCode != http.StatusOK {
        return nil, errors.New(r.Status)
    }

    resp := new(Response)
    err = json.NewDecoder(r.Body).Decode(resp)
    if err != nil{
        return nil, err
    }

    items := make([]Item, len(resp.Data.Children))
    for i, child := range resp.Data.Children {
        items[i] = child.Data
    }
    return items, nil
}

// func (i Item) String() string{
//     com := ""
//     switch i.LinkScore{
//     case 0:
//         //nothing
//     case 1:
//         com = " Score: 1"
//     default:
//         com = fmt.Sprintf(" (Score: %d)", i.LinkScore)
//     }
//     return fmt.Sprintf("%s\n%s", i.Title, i.URL)
// }

// func Email() string{
//     var buffer bytes.Buffer
//
//     items, err := Get("leagueoflegends")
//     if err != nil{
//         log.Fatal(err)
//     }
//     //Need to build strings from items
//     for _, item := range items{
//         buffer.WriteString(item.String())
//     }
//
//     return buffer.String()
// }

  type Item struct {
      Title string
      URL string
      LinkScore int `json: "score"`
  }

  type Response struct {
      Data struct{
          Children []struct{
              Data Item
          }
      }
  }
