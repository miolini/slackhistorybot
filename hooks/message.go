package hooks

import (
	"errors"
	"github.com/mholt/binding"
	"math"
	"net/http"
	"time"
)

type Message struct {
	Token          string    `json:"token" datastore:"-,noindex"`
	TeamId         string    `json:"team_id"`
	TeamDomain     string    `json:"team_domain"`
	ChannelId      string    `json:"channel_id"`
	ChannelName    string    `json:"channel_name"`
	TimestampFloat float64   `json:"-" datastore:"-,noindex" search:"-"`
	Timestamp      time.Time `json:"timestamp"`
	UserId         string    `json:"user_id"`
	UserName       string    `json:"user_name"`
	Text           string    `json:"text"`
}

func (msg Message) EntityKind() string {
	return "Message"
}

func (msg *Message) BindRequest(req *http.Request) error {
	bindErr := binding.Bind(req, msg)
	if bindErr.Error() != "" {
		return errors.New(bindErr.Error())
	} else if msg.Text == "" {
		return errors.New("text field is empty")
	}
	secs := math.Ceil(msg.TimestampFloat)
	nsecs := (msg.TimestampFloat - secs) * 1e9
	msg.Timestamp = time.Unix(int64(secs), int64(nsecs))
	return nil
}

func (msg *Message) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&msg.Token:          "token",
		&msg.TeamId:         "team_id",
		&msg.TeamDomain:     "team_domain",
		&msg.ChannelId:      "channel_id",
		&msg.ChannelName:    "channel_name",
		&msg.TimestampFloat: "timestamp",
		&msg.UserId:         "user_id",
		&msg.UserName:       "user_name",
		&msg.Text:           "text",
	}
}
