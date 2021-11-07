-- Schema
CREATE SCHEMA IF NOT EXISTS ingest;

-- Events
DROP TABLE IF EXISTS ingest.events;

CREATE TABLE ingest.events (
  id bigserial,
  cts bigint NOT NULL,
  uid text COLLATE pg_catalog."default",
  session text COLLATE pg_catalog."default",
  event text NOT NULL COLLATE pg_catalog."default",
  page text COLLATE pg_catalog."default",
  query text COLLATE pg_catalog."default",
  data text COLLATE pg_catalog."default",
  sts bigint NOT NULL,
  origin text COLLATE pg_catalog."default",
  referer text COLLATE pg_catalog."default",
  CONSTRAINT events_pkey PRIMARY KEY (id)
);

-- Users
DROP TABLE IF EXISTS ingest.users;

CREATE TABLE ingest.users (
  id bigserial,
  uid text NOT NULL COLLATE pg_catalog."default",
  ua text COLLATE pg_catalog."default",
  ua_hash text COLLATE pg_catalog."default",
  ip_addr inet,
  window_width int,
  window_height int,
  window_avail_width int,
  window_avail_height int,
  orientation text COLLATE pg_catalog."default",
  cts bigint NOT NULL,
  sts bigint NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS uid_btree_index ON ingest.users USING btree(uid);
CREATE INDEX IF NOT EXISTS ua_hash_btree_index ON ingest.users USING btree(ua_hash);

-- Users
DROP TABLE IF EXISTS ingest.sessions;

CREATE TABLE ingest.sessions (
  id bigserial,
  session text NOT NULL COLLATE pg_catalog."default",
  uid text NOT NULL COLLATE pg_catalog."default",
  ua text COLLATE pg_catalog."default",
  ua_hash text COLLATE pg_catalog."default",
  ip_addr inet,
  window_width int,
  window_height int,
  window_avail_width int,
  window_avail_height int,
  orientation text COLLATE pg_catalog."default",
  cts bigint NOT NULL,
  sts bigint NOT NULL,
  CONSTRAINT sessions_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS session_btree_index ON ingest.sessions USING btree(session);
CREATE INDEX IF NOT EXISTS ua_hash_btree_index ON ingest.sessions USING btree(ua_hash);
