BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.users
(
  id           UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  username     VARCHAR(255) NOT NULL,
  lang         VARCHAR(255) NOT NULL,
  country      VARCHAR(255) NOT NULL,
  first_name   VARCHAR(255) NOT NULL,
  last_name    VARCHAR(255) NOT NULL,
  email        VARCHAR(255) NOT NULL,
  birthday     TIMESTAMP(0) WITH TIME ZONE,
  password     VARCHAR(255) NOT NULL,
  dynamic_salt VARCHAR(255) NOT NULL,
  status       VARCHAR(255) NOT NULL,
  created_at   TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at  TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX users_username_uindex ON users (username);

CREATE TABLE public.lang_families
(
  id              UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug            VARCHAR(255) NOT NULL,
  origin          VARCHAR(255) NOT NULL,
  creator_user_id UUID REFERENCES users(id),
  created_at      TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at     TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX lang_families_slug_uindex ON lang_families (slug);

CREATE TABLE public.lang_subfamilies
(
  id                 UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug               VARCHAR(255) NOT NULL,
  family_slug        VARCHAR(255) REFERENCES lang_families(slug) NOT NULL,
  origin             VARCHAR(255) NOT NULL,
  extended_subfamily JSONB NOT NULL DEFAULT '{}'::JSONB,
  creator_user_id    UUID REFERENCES users(id),
  created_at         TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at        TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX lang_subfamilies_slug_uindex ON lang_subfamilies (slug);

CREATE TABLE public.lang_languages
(
  id               UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug             VARCHAR(255) NOT NULL,
  family_slug      VARCHAR(255) REFERENCES lang_families(slug) NOT NULL,
  subfamily_slug   VARCHAR(255) REFERENCES lang_subfamilies(slug) NOT NULL,
  wordbase_slug    VARCHAR(255) NOT NULL,
  female_own_names VARCHAR(255)[] NOT NULL,
  male_own_names   VARCHAR(255)[] NOT NULL,
  words            VARCHAR(255)[] NOT NULL,
  min              INTEGER NOT NULL,
  max              INTEGER NOT NULL,
  dupl             VARCHAR(255) NOT NULL,
  m                VARCHAR(255) NOT NULL,
  origin           VARCHAR(255) NOT NULL,
  creator_user_id  UUID REFERENCES users(id),
  created_at       TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at      TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX lang_languages_slug_uindex ON lang_languages (slug);
CREATE UNIQUE INDEX lang_languages_wordbase_slug_uindex ON lang_languages (wordbase_slug);

COMMIT;
