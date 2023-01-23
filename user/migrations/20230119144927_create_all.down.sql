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
-- reverse: create index "idx_services_service" to table: "services"
DROP INDEX "public"."idx_services_service";
-- reverse: create index "idx_services_deleted_at" to table: "services"
DROP INDEX "public"."idx_services_deleted_at";
-- reverse: create "services" table
DROP TABLE "public"."services";
-- reverse: create index "idx_roles_role" to table: "roles"
DROP INDEX "public"."idx_roles_role";
-- reverse: create index "idx_roles_deleted_at" to table: "roles"
DROP INDEX "public"."idx_roles_deleted_at";
-- reverse: create "roles" table
DROP TABLE "public"."roles";
-- reverse: create index "idx_methods_method" to table: "methods"
DROP INDEX "public"."idx_methods_method";
-- reverse: create index "idx_methods_deleted_at" to table: "methods"
DROP INDEX "public"."idx_methods_deleted_at";
-- reverse: create "methods" table
DROP TABLE "public"."methods";
