-- create "offices" table
CREATE TABLE "public"."offices" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "name" text NULL, "detail" text NULL, PRIMARY KEY ("id"));
-- create index "idx_offices_deleted_at" to table: "offices"
CREATE INDEX "idx_offices_deleted_at" ON "public"."offices" ("deleted_at");
-- create "schedules" table
CREATE TABLE "public"."schedules" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "office_id" bigint NULL, "total_capacity" bigint NULL, "capacity" bigint NULL, "date" timestamptz NULL, PRIMARY KEY ("id"), CONSTRAINT "fk_schedules_office" FOREIGN KEY ("office_id") REFERENCES "public"."offices" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "idx_schedules_date" to table: "schedules"
CREATE UNIQUE INDEX "idx_schedules_date" ON "public"."schedules" ("date");
-- create index "idx_schedules_deleted_at" to table: "schedules"
CREATE INDEX "idx_schedules_deleted_at" ON "public"."schedules" ("deleted_at");
-- create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL, "username" text NOT NULL, "email" text NOT NULL, "password" text NOT NULL, "role" character varying(255) NOT NULL, "avatar" text NULL, "verification_code" text NULL, "verified" boolean NULL DEFAULT false, "verified_at" timestamptz NULL, "password_reset_token" text NULL, "password_reset_at" timestamptz NULL, "otp_enabled" boolean NULL DEFAULT false, "otp_verified" boolean NULL DEFAULT false, "otp_secret" text NULL, "otp_auth_url" text NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- create index "idx_users_password_reset_token" to table: "users"
CREATE INDEX "idx_users_password_reset_token" ON "public"."users" ("password_reset_token");
-- create index "idx_users_username" to table: "users"
CREATE UNIQUE INDEX "idx_users_username" ON "public"."users" ("username");
-- create index "idx_users_verification_code" to table: "users"
CREATE INDEX "idx_users_verification_code" ON "public"."users" ("verification_code");
-- create "attendances" table
CREATE TABLE "public"."attendances" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "user_id" uuid NOT NULL, "schedule_id" bigint NOT NULL, "image_url" text NULL, "description" text NULL, "admin_username" text NULL, "status" text NULL DEFAULT 'pending', "status_at" timestamptz NULL, "status_info" text NULL, "check_temperature" numeric NULL, "check_status" text NULL DEFAULT 'pending', "check_in" timestamptz NULL, "check_out" timestamptz NULL, PRIMARY KEY ("id", "user_id", "schedule_id"), CONSTRAINT "fk_attendances_schedule" FOREIGN KEY ("schedule_id") REFERENCES "public"."schedules" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "fk_attendances_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "idx_attendances_deleted_at" to table: "attendances"
CREATE INDEX "idx_attendances_deleted_at" ON "public"."attendances" ("deleted_at");
-- create "certificates" table
CREATE TABLE "public"."certificates" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "user_id" uuid NULL, "dose" bigint NULL, "image_url" text NULL, "description" text NULL, "admin_username" text NULL, "status" text NULL DEFAULT 'pending', "status_at" timestamptz NULL, "status_info" text NULL, PRIMARY KEY ("id"), CONSTRAINT "fk_certificates_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "idx_certificates_deleted_at" to table: "certificates"
CREATE INDEX "idx_certificates_deleted_at" ON "public"."certificates" ("deleted_at");
-- create "workspaces" table
CREATE TABLE "public"."workspaces" ("office_id" bigint NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("office_id", "user_id"), CONSTRAINT "fk_workspaces_office" FOREIGN KEY ("office_id") REFERENCES "public"."offices" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "fk_workspaces_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
