BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.culture_bases
(
  id              UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug            VARCHAR(255) NOT NULL,
  origin          VARCHAR(255) NOT NULL,
  creator_user_id UUID REFERENCES users(id),
  created_at      TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at     TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX culture_bases_slug_uindex ON culture_bases (slug);

CREATE TABLE public.culture_subbases
(
  id              UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug            VARCHAR(255) NOT NULL,
  base_slug       VARCHAR(255) REFERENCES culture_bases(slug) NOT NULL,
  origin          VARCHAR(255) NOT NULL,
  creator_user_id UUID REFERENCES users(id),
  created_at      TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at     TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX culture_subbases_slug_uindex ON culture_subbases (slug);

CREATE TABLE public.cultures
(
  id                   UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  slug                 VARCHAR(255) NOT NULL,
  base_slug            VARCHAR(255) REFERENCES culture_bases(slug) NOT NULL,
  subbase_slug         VARCHAR(255) REFERENCES culture_subbases(slug) NOT NULL,
  ethos_slug           VARCHAR(255) NOT NULL,
  lang_slug            VARCHAR(255) REFERENCES lang_languages(slug) NOT NULL,
  dom_gender_name      VARCHAR(255) REFERENCES genders(name) NOT NULL,
  dom_gender_influence VARCHAR(255) NOT NULL,
  martial_custom       VARCHAR(255) REFERENCES gender_acceptances(name) NOT NULL,
  origin               VARCHAR(255) NOT NULL,
  creator_user_id      UUID REFERENCES users(id),
  created_at           TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  modified_at          TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX cultures_slug_uindex ON cultures (slug);

CREATE TABLE public.cultures_parents_relationships
(
  id                  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  parent_culture_slug VARCHAR(255) REFERENCES cultures(slug) NOT NULL,
  child_culture_slug  VARCHAR(255) REFERENCES cultures(slug) NOT NULL
);

CREATE TABLE public.cultures_traditions_relationships
(
  id              UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  culture_slug    VARCHAR(255) REFERENCES cultures(slug) NOT NULL,
  tradition_slug  VARCHAR(255) REFERENCES traditions(slug) NOT NULL
);

COMMIT;
