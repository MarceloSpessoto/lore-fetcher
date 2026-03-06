package mailingList

import (
	"crypto/tls"
	"log"
	"os"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"lore-fetcher/internal/core/domain"
)

const maxPatches = 20

type MailingListRepository struct {
	server   string
	username string
	password string
	mailbox  string
	useTLS   bool
}

func NewMailingListRepository() *MailingListRepository {
	server := os.Getenv("IMAP_SERVER")
	if server == "" {
		server = "localhost:993"
	}
	mailbox := os.Getenv("IMAP_MAILBOX")
	if mailbox == "" {
		mailbox = "INBOX"
	}
	return &MailingListRepository{
		server:   server,
		username: os.Getenv("IMAP_USERNAME"),
		password: os.Getenv("IMAP_PASSWORD"),
		mailbox:  mailbox,
		useTLS:   os.Getenv("IMAP_TLS") != "false",
	}
}

func (r *MailingListRepository) connect() (*client.Client, error) {
	if r.useTLS {
		return client.DialTLS(r.server, &tls.Config{})
	}
	return client.Dial(r.server)
}

func (r *MailingListRepository) GetRecentPatches() []domain.Patch {
	c, err := r.connect()
	if err != nil {
		log.Println("Failed to connect to IMAP server:", err)
		return []domain.Patch{}
	}
	defer c.Logout()

	if err = c.Login(r.username, r.password); err != nil {
		log.Println("Failed to login to IMAP server:", err)
		return []domain.Patch{}
	}

	if _, err = c.Select(r.mailbox, true); err != nil {
		log.Println("Failed to select mailbox:", err)
		return []domain.Patch{}
	}

	criteria := imap.NewSearchCriteria()
	criteria.Header.Add("Subject", "[PATCH]")
	uids, err := c.UidSearch(criteria)
	if err != nil {
		log.Println("Failed to search IMAP messages:", err)
		return []domain.Patch{}
	}
	if len(uids) == 0 {
		return []domain.Patch{}
	}

	// Keep the most recent N UIDs (UIDs are in ascending order)
	if len(uids) > maxPatches {
		uids = uids[len(uids)-maxPatches:]
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(uids...)

	messages := make(chan *imap.Message, len(uids))
	if err = c.UidFetch(seqSet, []imap.FetchItem{imap.FetchEnvelope, imap.FetchUid}, messages); err != nil {
		log.Println("Failed to fetch IMAP messages:", err)
		return []domain.Patch{}
	}

	patches := make([]domain.Patch, 0, maxPatches)
	for msg := range messages {
		if msg.Envelope == nil || len(msg.Envelope.From) == 0 {
			continue
		}
		env := msg.Envelope
		from := env.From[0]
		patches = append(patches, domain.Patch{
			Title:       env.Subject,
			AuthorName:  from.PersonalName,
			AuthorEmail: from.MailboxName + "@" + from.HostName,
			PatchHref:   strings.Trim(env.MessageId, "<>"),
		})
	}

	// Reverse so the most recent patch is first, matching the lore repository behaviour
	for i, j := 0, len(patches)-1; i < j; i, j = i+1, j-1 {
		patches[i], patches[j] = patches[j], patches[i]
	}
	return patches
}
