-- Copyright 2025 The Omnigo Authors.
-- Licensed under the Apache License, Version 2.0.
-- See: http://www.apache.org/licenses/LICENSE-2.0

-- +goose Up
-- Migration to create core chat schema
CREATE SCHEMA IF NOT EXISTS core;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Enums
CREATE TYPE core.participant_role AS ENUM ('customer', 'agent');
CREATE TYPE core.chat_status AS ENUM ('active', 'closed', 'archived');
CREATE TYPE core.message_direction AS ENUM ('inbound', 'outbound');
CREATE TYPE core.presence_status AS ENUM ('online', 'offline', 'idle', 'typing');


-- Channels
CREATE TABLE core.channels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    provider_id TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Participants
CREATE TABLE core.participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    provider_id TEXT UNIQUE,
    role core.participant_role NOT NULL,
    name TEXT NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Chats
CREATE TABLE core.chats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    channel_id UUID NOT NULL REFERENCES core.channels(id) ON DELETE CASCADE,
    provider_id TEXT UNIQUE,
    topic TEXT,
    status core.chat_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_chats_channel_id ON core.chats(channel_id);

-- Join table for chat participants
CREATE TABLE core.chat_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_id UUID NOT NULL REFERENCES core.chats(id) ON DELETE CASCADE,
    participant_id UUID NOT NULL REFERENCES core.participants(id) ON DELETE CASCADE,
    joined_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    left_at TIMESTAMP WITH TIME ZONE,
    UNIQUE (chat_id, participant_id)
);

CREATE INDEX IF NOT EXISTS idx_chat_participants_participant_id ON core.chat_participants(participant_id);

-- Messages
CREATE TABLE core.messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_id UUID NOT NULL REFERENCES core.chats(id) ON DELETE CASCADE,
    participant_id UUID NOT NULL REFERENCES core.participants(id) ON DELETE CASCADE,
    provider_id TEXT UNIQUE,

    content TEXT,
    direction core.message_direction NOT NULL,

    received_at TIMESTAMP WITH TIME ZONE,
    sent_at TIMESTAMP WITH TIME ZONE,
    edited_at TIMESTAMP WITH TIME ZONE,
    removed_at TIMESTAMP WITH TIME ZONE,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_messages_chat_id_sent_at ON core.messages(chat_id, sent_at);

-- Message Attachments
CREATE TABLE core.message_attachments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    message_id UUID NOT NULL REFERENCES core.messages(id) ON DELETE CASCADE,
    provider_id TEXT,
    mime_type TEXT NOT NULL,
    url TEXT NOT NULL,
    name TEXT,
    size BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- Participant Presence
CREATE TABLE core.presence (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    participant_id UUID NOT NULL REFERENCES core.participants(id) ON DELETE CASCADE,
    chat_id UUID REFERENCES core.chats(id) ON DELETE CASCADE,
    status core.presence_status NOT NULL,
    last_seen_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_presence_participant_id ON core.presence(participant_id);
CREATE INDEX IF NOT EXISTS idx_presence_chat_id ON core.presence(chat_id);

-- +goose Down
-- Drop dependent tables first (order matters)
DROP TABLE IF EXISTS core.presence CASCADE;
DROP TABLE IF EXISTS core.message_attachments CASCADE;
DROP TABLE IF EXISTS core.messages CASCADE;
DROP TABLE IF EXISTS core.chat_participants CASCADE;
DROP TABLE IF EXISTS core.chats CASCADE;
DROP TABLE IF EXISTS core.participants CASCADE;
DROP TABLE IF EXISTS core.channels CASCADE;

-- Drop enums after tables
DROP TYPE IF EXISTS core.presence_status;
DROP TYPE IF EXISTS core.message_direction;
DROP TYPE IF EXISTS core.chat_status;
DROP TYPE IF EXISTS core.participant_role;

-- Finally drop schema
DROP SCHEMA IF EXISTS core CASCADE;