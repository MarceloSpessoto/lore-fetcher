package fetcher

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
  "strings"
  "time"
	"golang.org/x/text/encoding/ianaindex"
  "github.com/MarceloSpessoto/lore-fetcher/internal/types"
)

type Feed struct {
  Entries []struct {
    Name string `xml:"author>name"`
    Email string `xml:"author>email"`
    Title string `xml:"title"`
    Link struct {
      Href string `xml:"href,attr"`
    } `xml:"link"`
  } `xml:"entry"`
} 

type Fetcher struct {
  feed Feed
  patchStatus map[string]bool
  MailingList string
  FetchInterval int  
}

func NewFetcher() *Fetcher {
  var fetcher Fetcher
  fetcher.patchStatus = make(map[string]bool)
  return &fetcher
}

func parsePatchTag(href string) string{
  hrefComponents := strings.Split(href, "/")
  patchTagComponents := strings.Split(hrefComponents[4], "-")
  return patchTagComponents[0]
}

func (fetcher *Fetcher) FetchDaemon(fetchBuffer chan types.Patch){
  for {
    fetcher.GetPatches()
    if len(fetcher.feed.Entries) == 0 {
      time.Sleep(20 * time.Second)
      continue
    }
    fmt.Println("Most recent patch from ", fetcher.MailingList, ":\n", fetcher.feed.Entries[0].Title)
    firstPatchHref := fetcher.feed.Entries[0].Link.Href
    patchTag := parsePatchTag(firstPatchHref)
    fetcher.patchStatus[patchTag] = true
    break
  }

  for {
    time.Sleep(30 * time.Second)

    fmt.Println("[", time.Now(), "]: Searching for new patches in", fetcher.MailingList)
    fetcher.GetPatches()
    fetcher.processPatches(fetchBuffer)
  }
}

func (fetcher *Fetcher) FetchBatch(fetchBuffer chan types.Patch){
  fetcher.GetPatches()
  fetcher.processPatches(fetchBuffer)
}

func (fetcher *Fetcher) GetPatches() {
  var fetchUrl string = fmt.Sprintf("https://lore.kernel.org/%s/?q=rt:..+AND+NOT+s:Re&x=A", fetcher.MailingList)
  resp, err := http.Get(fetchUrl)
  if err != nil {
    fmt.Println(err)
    return 
  }

  // Obtain XML content from HTTP Request Body and put it into a decoder.
  xmlStream := xml.NewDecoder(resp.Body)
  defer resp.Body.Close()

  // Enable parsing of XML files with encodings different from UTF-8,
  // such as US-ASCII, used by the Lore.
  xmlStream.CharsetReader = func(charset string, reader io.Reader) (io.Reader, error) {
    enc, err := ianaindex.IANA.Encoding(charset)
    if err != nil {
      return nil, fmt.Errorf("charset %s: %s", charset, err.Error())
    }
    return enc.NewDecoder().Reader(reader), nil
  }

  var feed Feed 
  err = xmlStream.Decode(&feed)
  if err != nil {
    fmt.Println(err)
  }
  fetcher.feed = feed
}

func (fetcher *Fetcher) processPatches(fetchBuffer chan types.Patch){
  for i := 0; i < len(fetcher.feed.Entries); i++ {
    patchHref := fetcher.feed.Entries[i].Link.Href
    patchTag := parsePatchTag(patchHref)
    fmt.Println(patchHref)
    fmt.Println(patchTag)

    if _, ok := fetcher.patchStatus[patchTag]; !ok {
      var patch types.Patch
      patch.Title = fetcher.feed.Entries[i].Title
      patch.AuthorName = fetcher.feed.Entries[i].Name
      patch.AuthorEmail = fetcher.feed.Entries[i].Email
      patch.PatchHref = patchHref
      patch.PatchTag = patchTag
      fetcher.patchStatus[patchTag] = true
      fetchBuffer <- patch
      fmt.Println("[", time.Now(), "]: Sending patch '", patch.Title, "' to CI Pipeline")
    } else {
      break
    }
  }
}
