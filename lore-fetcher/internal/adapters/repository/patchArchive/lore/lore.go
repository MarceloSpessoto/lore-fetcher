package lore
import (
	"io"
	"encoding/xml"
	"net/http"
	"lore-fetcher/internal/core/domain"
	"golang.org/x/text/encoding/ianaindex"
	"fmt"
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

type LoreRepository struct {
	feed Feed
}

func NewLoreRepository() *LoreRepository {
	return &LoreRepository{}
}

func (lr *LoreRepository) GetRecentPatches() []domain.Patch {
	fetchUrl := "https://lore.kernel.org/all/?q=rt:..+AND+NOT+s:Re&x=A"
	resp, err := http.Get(fetchUrl)
	if err != nil {
		fmt.Println(err)
		return make([]domain.Patch, 0)
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

	patches := make([]domain.Patch, len(feed.Entries))
	for i, entry := range feed.Entries {
		patches[i] = domain.Patch{
			AuthorName:  entry.Name,
			AuthorEmail: entry.Email,
			Title: entry.Title,
			PatchHref:  entry.Link.Href,
		}
	}
	return patches
}
