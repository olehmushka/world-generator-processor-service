BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.traditions
(
  id                    UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug                  VARCHAR(255) NOT NULL,
  description           VARCHAR(255) NOT NULL,
  preferred_ethos_slugs VARCHAR(255)[],
  type                  VARCHAR(255) NOT NULL,
  omit_tradition_slugs  VARCHAR(255)[],
  omit_gender_dominance VARCHAR(255)[],
  omit_ethos_slugs      VARCHAR(255)[],
  origin                VARCHAR(255) NOT NULL,
  creator_user_id       UUID REFERENCES users(id),
  created_at            TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at           TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX traditions_slug_uindex ON traditions (slug);

COMMIT;
