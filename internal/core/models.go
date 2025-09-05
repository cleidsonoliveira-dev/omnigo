package core

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id         uuid.UUID
	ExternalId string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Participant struct {
	Id         uuid.UUID
	ExternalId string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Chat struct {
	Id         uuid.UUID
	ChannelId  uuid.UUID
	ExternalId string

	Channel      Channel
	Participants []Participant
}

type Message struct {
	Id            uuid.UUID
	ChatId        uuid.UUID
	ParticipantId uuid.UUID
	ExternalId    string

	Content     string
	Attachments []MessageAttachment
	Edited      bool
	Removed     bool

	ReceivedAt         *time.Time
	SentAt             *time.Time
	EditedAt           *time.Time
	RemovedAt          *time.Time
	RemovalConfirmedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageAttachment struct {
	Id         uuid.UUID
	MessageId  uuid.UUID
	ExternalId string

	MimeType string
	Url      string
}

type ParticipantPresenceStatus string

const (
	ParticipantPresenceOnline  ParticipantPresenceStatus = "online"
	ParticipantPresenceOffline ParticipantPresenceStatus = "offline"
	ParticipantPresenceIdle    ParticipantPresenceStatus = "idle"
)

type ParticipantPresence struct {
	Id            uuid.UUID
	ParticipantId uuid.UUID

	Status     ParticipantPresenceStatus
	ReceivedAt time.Time
	CreatedAt  time.Time
}

type ChatPresenceStatus string

const (
	ChatPresenceTyping    ChatPresenceStatus = "typing"
	ChatPresenceRecording ChatPresenceStatus = "recording"
)

type ChatPresence struct {
	Id            uuid.UUID
	ChatId        uuid.UUID
	ParticipantId uuid.UUID

	Status     ChatPresenceStatus
	ReceivedAt time.Time
	CreatedAt  time.Time
}
