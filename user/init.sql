CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial,
  "username" varchar(50) NOT NULL,
  "password" varchar(100) NOT NULL,
  "email" varchar(100) NOT NULL,
  "phone" varchar(20) DEFAULT null,
  "avatar" varchar(255) DEFAULT null,
  "nickname" varchar(50) DEFAULT null,
  "introduction" text DEFAULT null,
  "role" varchar(10) NOT NULL DEFAULT 'user',
  "status" smallint NOT NULL DEFAULT 1,
  "last_login_at" timestamptz DEFAULT null,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz,
  PRIMARY KEY ("id")
);