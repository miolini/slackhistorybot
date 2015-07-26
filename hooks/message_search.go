package hooks

import (
	"appengine/search"
	"time"
)

type MessageSearch struct {
	TeamId      search.Atom
	TeamDomain  search.Atom
	ChannelId   search.Atom
	ChannelName search.Atom
	Timestamp   time.Time
	UserId      search.Atom
	UserName    search.Atom
	Text        search.HTML
}

func MessageSearchFromMessage(msg Message) MessageSearch {
	msgSearch := MessageSearch{}
	msgSearch.TeamId = search.Atom(msg.TeamId)
	msgSearch.TeamDomain = search.Atom(msg.TeamDomain)
	msgSearch.ChannelId = search.Atom(msg.ChannelId)
	msgSearch.ChannelName = search.Atom(msg.ChannelName)
	msgSearch.Timestamp = msg.Timestamp
	msgSearch.UserId = search.Atom(msg.UserId)
	msgSearch.UserName = search.Atom(msg.UserName)
	msgSearch.Text = search.HTML(msg.Text)
	return msgSearch
}
