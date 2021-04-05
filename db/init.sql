DROP TABLE IF EXISTS public.events;

CREATE TABLE public.events (
  id bigserial,
  event text NOT NULL COLLATE pg_catalog."default",
  data text COLLATE pg_catalog."default",
  cts bigint NOT NULL,
  sts bigint NOT NULL,
  origin text COLLATE pg_catalog."default",
  referer text COLLATE pg_catalog."default",
  ua_hash text COLLATE pg_catalog."default",
  uid text COLLATE pg_catalog."default",
  CONSTRAINT events_pkey PRIMARY KEY (id)
);


DROP TABLE IF EXISTS public.useragents;

CREATE TABLE public.useragents (
  id bigserial,
  ua text NOT NULL COLLATE pg_catalog."default",
  ua_hash text NOT NULL UNIQUE,
  CONSTRAINT useragents_pkey PRIMARY KEY (id)
);

CREATE INDEX ua_hash_btree_index ON public.useragents USING btree(ua_hash);