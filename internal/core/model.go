package core

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Participant struct {
	Id        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Chat struct {
	Id        uuid.UUID
	ChannelId uuid.UUID

	Channel      Channel
	Participants []Participant
}

type Message struct {
	Id            uuid.UUID
	ChatId        uuid.UUID
	ParticipantId uuid.UUID

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
	Id        uuid.UUID
	MessageId uuid.UUID

	MimeType string
	Url      string
}

type Presence struct {
	Id uuid.UUID

	ReceivedAt time.Time
	CreatedAt  time.Time
}
