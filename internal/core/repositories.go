package core

import (
	"context"

	"github.com/google/uuid"
)

type ChannelRepository interface {
	create(ctx context.Context, channel Channel) (Channel, error)
	save(ctx context.Context, channel Channel) (Channel, error)
	getById(ctx context.Context, id uuid.UUID) (Channel, error)
}

type ChatRepository interface {
	getById(ctx context.Context, id uuid.UUID) (Chat, error)
	create(ctx context.Context, chat Chat) (Chat, error)
	save(ctx context.Context, chat Chat) (Chat, error)
	sendPresence(ctx context.Context, presence ChatPresence) (Chat, error)
}

type ParticipantRepository interface {
	getById(ctx context.Context, id uuid.UUID) (Participant, error)
	create(ctx context.Context, participant Chat) (Participant, error)
	save(ctx context.Context, participant Chat) (Participant, error)
	sendPresence(ctx context.Context, participant ParticipantPresence) (Chat, error)
}
