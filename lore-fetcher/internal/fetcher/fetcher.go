package fetcher

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
	"golang.org/x/text/encoding/ianaindex"
  "lore-fetcher/internal/core/services"
  "lore-fetcher/internal/core/domain"
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
  service services.LFService
}

func NewFetcher(service services.LFService) *Fetcher {
  var fetcher Fetcher
  fetcher.patchStatus = make(map[string]bool)
  fetcher.service = service
  return &fetcher
}

func parsePatchTag(href string) string{
  hrefComponents := strings.Split(href, "/")
  patchTagComponents := strings.Split(hrefComponents[4], "-")
  return patchTagComponents[0]
}

func (fetcher *Fetcher) FetchDaemon(){
  for {
    fetcher.GetPatches()
    if len(fetcher.feed.Entries) == 0 {
      time.Sleep(20 * time.Second)
      continue
    }
    fmt.Println("Most recent patch from all:\n", fetcher.feed.Entries[0].Title)
    firstPatchHref := fetcher.feed.Entries[0].Link.Href
    patchTag := parsePatchTag(firstPatchHref)
    fetcher.patchStatus[patchTag] = true
    var patch domain.Patch
    patch.Title = fetcher.feed.Entries[0].Title
    patch.AuthorName = fetcher.feed.Entries[0].Name
    patch.AuthorEmail = fetcher.feed.Entries[0].Email
    patch.PatchHref = firstPatchHref
    fetcher.service.SavePatch(patch)
    fmt.Println("New patch found: ", patch.Title)
    fmt.Println(patch)
    break
  }

  for {
    time.Sleep(30 * time.Second)

    fmt.Println("[", time.Now(), "]: Searching for new patches in all")
    fetcher.GetPatches()
    fetcher.processPatches()
  }
}

func (fetcher *Fetcher) FetchBatch(){
  fetcher.GetPatches()
  fetcher.processPatches()
}

func (fetcher *Fetcher) GetPatches() {
  var fetchUrl string = "https://lore.kernel.org/all/?q=rt:..+AND+NOT+s:Re&x=A"
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

func (fetcher *Fetcher) processPatches(){
  for i := 0; i < len(fetcher.feed.Entries); i++ {
    patchHref := fetcher.feed.Entries[i].Link.Href
    patchTag := parsePatchTag(patchHref)

    if _, ok := fetcher.patchStatus[patchTag]; !ok {
      var patch domain.Patch
      patch.Title = fetcher.feed.Entries[i].Title
      patch.AuthorName = fetcher.feed.Entries[i].Name
      patch.AuthorEmail = fetcher.feed.Entries[i].Email
      patch.PatchHref = patchHref
      patch.PatchTag = patchTag
      fetcher.patchStatus[patchTag] = true
      if isPatch(patch.Title){
        fetcher.service.SavePatch(patch)
        fmt.Println("New patch found: ", patch.Title)
      }
    } else {
      break
    }
  }
}

// A temporary method to assert we're fetching a Patch message:
// checking if the title string contains the [PATCH] tag
func isPatch(patchTitle string) bool {
  fmt.Println(patchTitle)
  hasPattern, _ := regexp.Match(`.*\[.*PATCH.*\].*`, []byte(patchTitle))
  return hasPattern
}
