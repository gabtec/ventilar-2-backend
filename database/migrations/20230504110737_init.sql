-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Create "users" table
CREATE TABLE "public"."users" ("id" bigserial NOT NULL, "name" text NULL, "role" text NULL, "mec" bigint NULL, "password_hash" text NULL, "ward_id" bigint NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "users_mec_key" to table: "users"
CREATE UNIQUE INDEX "users_mec_key" ON "public"."users" ("mec");
-- Create "wards" table
CREATE TABLE "public"."wards" ("id" bigserial NOT NULL, "name" text NULL, "belongs_to" text NULL, "is_park" boolean NULL DEFAULT false, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, PRIMARY KEY ("id"));
