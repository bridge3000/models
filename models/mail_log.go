package models

import (
	"../Adoter_Asset"
	"gorm.io/gorm"
)

type MailLog struct {
	gorm.Model    `xml:"-"`
	Title         string                          `xml:"title"`
	Content       string                          `xml:"content"`
	Sender        string                          `xml:"-"`
	Caller        string                          `xml:"caller"`
	Attachment    string                          `xml:"attachment"`
	Admin_id      uint                            `xml:"-"`
	AdminName     string                          `xml:"-"`
	Response_code Adoter_Asset.COMMAND_ERROR_CODE `xml:"-"`
	Zone_id       int                             `xml:"-"`
	Zone_name     string                          `xml:"-"`
	Guid          uint64                          `xml:"-"`
	Player_guid   uint64                          `xml:"-"`
	Hero_guid     uint64                          `xml:"-"`
	Mail_id       uint32                          `xml:"-"`
	Valid_time    uint32                          `xml:"valid_time"`
	Result        int                             `xml:"-"`
	Error_message string                          `xml:"-"`
	Player_id     uint64                          `xml:"-"`
	PlayerName    string                          `xml:"-"`
	ChannelType   int                             `xml:"channel_type"`
}
