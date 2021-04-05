-- Events
DROP TABLE IF EXISTS public.events;

CREATE TABLE public.events (
  id bigserial,
  event text NOT NULL COLLATE pg_catalog."default",
  data text COLLATE pg_catalog."default",
  cts bigint NOT NULL,
  sts bigint NOT NULL,
  origin text COLLATE pg_catalog."default",
  referer text COLLATE pg_catalog."default",
  uid text COLLATE pg_catalog."default",
  CONSTRAINT events_pkey PRIMARY KEY (id)
);

-- Users
DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users (
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

CREATE INDEX uid_btree_index ON public.users USING btree(uid);
CREATE INDEX ua_hash_btree_index ON public.users USING btree(ua_hash);