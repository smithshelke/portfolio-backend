-- Create "features" table
CREATE TABLE "public"."features" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "created_by" uuid NULL,
  "priority" text NULL,
  "status" text NULL,
  PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "role" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "created_by" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "feature_owners" table
CREATE TABLE "public"."feature_owners" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "feature_id" uuid NOT NULL,
  "user_name" text NULL,
  "user_role" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "feature_owners_feature_id_fkey" FOREIGN KEY ("feature_id") REFERENCES "public"."features" ("id") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "feature_owners_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create "tasks" table
CREATE TABLE "public"."tasks" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "created_by" uuid NULL,
  "feature_id" uuid NOT NULL,
  "feature_name" text NULL,
  "priority" text NULL,
  "status" text NULL,
  "git_data" jsonb NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "tasks_feature_id_fkey" FOREIGN KEY ("feature_id") REFERENCES "public"."features" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
