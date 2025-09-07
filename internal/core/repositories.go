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
	"context"

	"github.com/google/uuid"
)

type ChannelRepository interface {
    Create(ctx context.Context, channel *Channel) (*Channel, error)
    Save(ctx context.Context, channel *Channel) (*Channel, error)
    GetByID(ctx context.Context, id uuid.UUID) (*Channel, error)
}

type ChatRepository interface {
    GetByID(ctx context.Context, id uuid.UUID) (*Chat, error)
    Create(ctx context.Context, chat *Chat) (*Chat, error)
    Save(ctx context.Context, chat *Chat) (*Chat, error)
}

type ParticipantRepository interface {
    GetByID(ctx context.Context, id uuid.UUID) (*Participant, error)
    Create(ctx context.Context, participant *Participant) (*Participant, error)
    Save(ctx context.Context, participant *Participant) (*Participant, error)
    SendPresence(ctx context.Context, participant *Presence) (*Presence, error)
}
