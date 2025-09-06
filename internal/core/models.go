// Copyright 2025 The Omnigo Authors. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
