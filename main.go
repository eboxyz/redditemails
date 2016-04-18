package main

import(
    "log"
    "net/http"
    "fmt"
    "encoding/json"
    "errors"
)

func main() {

    items, err := Get("leagueoflegends")
    if err != nil {
        log.Fatal(err)
    }

    for _, item := range items {
        fmt.Println(item)
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

func (i Item) String() string{
    com := ""
    switch i.LinkScore{
    case 0:
        //nothing
    case 1:
        com = " Score: 1"
    default:
        com = fmt.Sprintf(" (Score: %d)", i.LinkScore)
    }
    return fmt.Sprintf("%s\n%s", i.Title, i.URL)
}


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
