-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS repository_id_seq;

-- Table Definition
CREATE TABLE "public"."repository" (
    "id" int8 NOT NULL DEFAULT nextval('repository_id_seq'::regclass),
    "name" text NOT NULL,
    "url" text NOT NULL,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS result_id_seq;

-- Table Definition
CREATE TABLE "public"."result" (
    "id" int8 NOT NULL DEFAULT nextval('result_id_seq'::regclass),
    "status" text NOT NULL,
    "repository_name" text NOT NULL,
    "repository_url" text NOT NULL,
    "queued_at" timestamp,
    "scanning_at" timestamp,
    "finished_at" timestamp,
    "findings" jsonb,
    PRIMARY KEY ("id")
);