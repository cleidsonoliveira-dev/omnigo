// Copyright 2025 The Omnigo Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
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
	ID          uuid.UUID
	ProviderID  string
	Name        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ParticipantRole string

const (
	ParticipantRoleCustomer ParticipantRole = "customer"
	ParticipantRoleAgent    ParticipantRole = "agent"
)

type Participant struct {
	ID         uuid.UUID
	ProviderID *string
	Role       ParticipantRole
	Name       string
	Metadata   map[string]any

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChatStatus string

const (
	ChatStatusActive   ChatStatus = "active"
	ChatStatusClosed   ChatStatus = "closed"
	ChatStatusArchived ChatStatus = "archived"
)

type Chat struct {
	ID         uuid.UUID
	ChannelID  uuid.UUID
	ProviderID *string
	Topic      string
	Status     ChatStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time

	ParticipantIDs []uuid.UUID
}

type ChatParticipant struct {
	ID            uuid.UUID
	ChatID        uuid.UUID
	ParticipantID uuid.UUID
	JoinedAt      time.Time
	LeftAt        *time.Time
}

type MessageDirection string

const (
	MessageDirectionInbound  MessageDirection = "inbound"
	MessageDirectionOutbound MessageDirection = "outbound"
)

// Messages
type Message struct {
	ID            uuid.UUID
	ChatID        uuid.UUID
	ParticipantID uuid.UUID

	Content     string
	Attachments []MessageAttachment

	Direction MessageDirection

	ReceivedAt *time.Time
	SentAt     *time.Time
	EditedAt   *time.Time
	RemovedAt  *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageAttachment struct {
	ID         uuid.UUID
	MessageID  uuid.UUID
	ProviderID string

	MimeType string
	URL      string
	Name     string
	Size     int64 // bytes
}

type PresenceStatus string

const (
	PresenceOnline  PresenceStatus = "online"
	PresenceOffline PresenceStatus = "offline"
	PresenceIdle    PresenceStatus = "idle"
	PresenceTyping  PresenceStatus = "typing"
)

type Presence struct {
	ID            uuid.UUID
	ParticipantID uuid.UUID
	ChatID        *uuid.UUID
	Status        PresenceStatus
	LastSeenAt    time.Time
	CreatedAt     time.Time
}
