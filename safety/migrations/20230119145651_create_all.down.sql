-- reverse: create "workspaces" table
DROP TABLE "public"."workspaces";
-- reverse: create index "idx_certificates_deleted_at" to table: "certificates"
DROP INDEX "public"."idx_certificates_deleted_at";
-- reverse: create "certificates" table
DROP TABLE "public"."certificates";
-- reverse: create index "idx_attendances_deleted_at" to table: "attendances"
DROP INDEX "public"."idx_attendances_deleted_at";
-- reverse: create "attendances" table
DROP TABLE "public"."attendances";
-- reverse: create index "idx_users_verification_code" to table: "users"
DROP INDEX "public"."idx_users_verification_code";
-- reverse: create index "idx_users_username" to table: "users"
DROP INDEX "public"."idx_users_username";
-- reverse: create index "idx_users_password_reset_token" to table: "users"
DROP INDEX "public"."idx_users_password_reset_token";
-- reverse: create index "idx_users_email" to table: "users"
DROP INDEX "public"."idx_users_email";
-- reverse: create "users" table
DROP TABLE "public"."users";
-- reverse: create index "idx_schedules_deleted_at" to table: "schedules"
DROP INDEX "public"."idx_schedules_deleted_at";
-- reverse: create index "idx_schedules_date" to table: "schedules"
DROP INDEX "public"."idx_schedules_date";
-- reverse: create "schedules" table
DROP TABLE "public"."schedules";
-- reverse: create index "idx_offices_deleted_at" to table: "offices"
DROP INDEX "public"."idx_offices_deleted_at";
-- reverse: create "offices" table
DROP TABLE "public"."offices";
