BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.genders
(
  id              UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  name            VARCHAR(255) NOT NULL,
  origin          VARCHAR(255) NOT NULL,
  creator_user_id UUID REFERENCES users(id),
  created_at      TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at     TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX genders_name_uindex ON genders (name);

COMMIT;
