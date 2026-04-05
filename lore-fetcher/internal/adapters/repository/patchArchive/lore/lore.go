package lore

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/text/encoding/ianaindex"
	"lore-fetcher/internal/core/domain"
)

type Feed struct {
	Entries []struct {
		Name  string `xml:"author>name"`
		Email string `xml:"author>email"`
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
	} `xml:"entry"`
}

type LoreRepository struct {
	feed        Feed
	subsystem   string
	loreBaseURL string
}

func NewLoreRepository() *LoreRepository {
	subsystem := os.Getenv("LORE_SUBSYSTEM")
	if subsystem == "" {
		subsystem = "all"
	}
	loreBaseURL := os.Getenv("LORE_BASE_URL")
	if loreBaseURL == "" {
		loreBaseURL = "https://lore.kernel.org"
	}
	return &LoreRepository{
		subsystem:   subsystem,
		loreBaseURL: strings.TrimRight(loreBaseURL, "/"),
	}
}

func (lr *LoreRepository) GetRecentPatches() []domain.Patch {
	fetchUrl := lr.loreBaseURL + "/" + lr.subsystem + "/?q=rt:..+AND+NOT+s:Re&x=A"
	resp, err := http.Get(fetchUrl)
	if err != nil {
		log.Println("Failed to fetch patches:", err)
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
			return nil, err
		}
		return enc.NewDecoder().Reader(reader), nil
	}

	var feed Feed
	if err = xmlStream.Decode(&feed); err != nil {
		log.Println("Failed to decode patches feed:", err)
	}

	patches := make([]domain.Patch, len(feed.Entries))
	for i, entry := range feed.Entries {
		patches[i] = domain.Patch{
			AuthorName:  entry.Name,
			AuthorEmail: entry.Email,
			Title:       entry.Title,
			PatchHref:   entry.Link.Href,
		}
	}
	return patches
}
