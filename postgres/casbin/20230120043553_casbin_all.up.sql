-- create "casbin_rule" table
CREATE TABLE "public"."casbin_rule" ("id" bigserial NOT NULL, "ptype" character varying(100) NULL, "v0" character varying(100) NULL, "v1" character varying(100) NULL, "v2" character varying(100) NULL, "v3" character varying(100) NULL, "v4" character varying(100) NULL, "v5" character varying(100) NULL, PRIMARY KEY ("id"));
-- create index "idx_casbin_rule" to table: "casbin_rule"
CREATE UNIQUE INDEX "idx_casbin_rule" ON "public"."casbin_rule" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5");

INSERT INTO "casbin_rule" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES
	('g', '028de3d9-2f46-4cb2-92f0-2ca723e31862', 'admin', NULL, NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'RefreshAccessToken', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'UpdateMe', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'ChangeEmail', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'Logout', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'GetMe', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'CreateUser', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'UpdateUserById', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'DeleteUserById', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'FindUsers', NULL, NULL, NULL),
	('p', 'admin', 'pb.UserService', 'FindUserById', NULL, NULL, NULL),
	('p', 'admin', 'pb.RoleService', 'CreateRole', NULL, NULL, NULL),
	('p', 'admin', 'pb.RoleService', 'DeleteRoleById', NULL, NULL, NULL),
	('p', 'admin', 'pb.RoleService', 'FindRoles', NULL, NULL, NULL),
	('p', 'admin', 'pb.ServiceService', 'CreateService', NULL, NULL, NULL),
	('p', 'admin', 'pb.ServiceService', 'DeleteServiceById', NULL, NULL, NULL),
	('p', 'admin', 'pb.ServiceService', 'FindServices', NULL, NULL, NULL),
	('p', 'admin', 'pb.MethodService', 'CreateMethod', NULL, NULL, NULL),
	('p', 'admin', 'pb.MethodService', 'DeleteMethodById', NULL, NULL, NULL),
	('p', 'admin', 'pb.MethodService', 'FindMethods', NULL, NULL, NULL),
	('p', 'admin', 'pb.PolicyService', 'CreatePolicy', NULL, NULL, NULL),
	('p', 'admin', 'pb.PolicyService', 'DeletePolicy', NULL, NULL, NULL),
	('p', 'admin', 'pb.PolicyService', 'FindPolicies', NULL, NULL, NULL),
	('p', 'user', 'pb.UserService', 'RefreshAccessToken', NULL, NULL, NULL),
	('p', 'user', 'pb.UserService', 'UpdateMe', NULL, NULL, NULL),
	('p', 'user', 'pb.UserService', 'ChangeEmail', NULL, NULL, NULL),
	('p', 'user', 'pb.UserService', 'Logout', NULL, NULL, NULL),
	('p', 'user', 'pb.UserService', 'GetMe', NULL, NULL, NULL),
	('p', 'admin', 'pb.AttendanceService', 'CreateAttendance', NULL, NULL, NULL),
	('p', 'admin', 'pb.AttendanceService', 'UpdateAttendanceById', NULL, NULL, NULL),
	('p', 'admin', 'pb.AttendanceService', 'DeleteAttendanceById', NULL, NULL, NULL),
	('p', 'admin', 'pb.AttendanceService', 'FindAttendances', NULL, NULL, NULL),
	('p', 'admin', 'pb.AttendanceService', 'FindAttendanceById', NULL, NULL, NULL),
	('p', 'admin', 'pb.CertificateService', 'CreateCertificate', NULL, NULL, NULL),
	('p', 'admin', 'pb.CertificateService', 'UpdateCertificateById', NULL, NULL, NULL),
	('p', 'admin', 'pb.CertificateService', 'DeleteCertificateById', NULL, NULL, NULL),
	('p', 'admin', 'pb.CertificateService', 'FindCertificates', NULL, NULL, NULL),
	('p', 'admin', 'pb.CertificateService', 'FindCertificateById', NULL, NULL, NULL),
	('p', 'admin', 'pb.CheckService', 'CheckIn', NULL, NULL, NULL),
	('p', 'admin', 'pb.CheckService', 'CheckOut', NULL, NULL, NULL),
	('p', 'admin', 'pb.CheckService', 'FindChecks', NULL, NULL, NULL),
	('p', 'admin', 'pb.CheckService', 'FindCheckById', NULL, NULL, NULL),
	('p', 'admin', 'pb.OfficeService', 'CreateOffice', NULL, NULL, NULL),
	('p', 'admin', 'pb.OfficeService', 'UpdateOfficeById', NULL, NULL, NULL),
	('p', 'admin', 'pb.OfficeService', 'DeleteOfficeById', NULL, NULL, NULL),
	('p', 'admin', 'pb.OfficeService', 'FindOffices', NULL, NULL, NULL),
	('p', 'admin', 'pb.OfficeService', 'FindOfficeById', NULL, NULL, NULL),
	('p', 'admin', 'pb.ScheduleService', 'CreateSchedule', NULL, NULL, NULL),
	('p', 'admin', 'pb.ScheduleService', 'UpdateScheduleById', NULL, NULL, NULL),
	('p', 'admin', 'pb.ScheduleService', 'DeleteScheduleById', NULL, NULL, NULL),
	('p', 'admin', 'pb.ScheduleService', 'FindSchedules', NULL, NULL, NULL),
	('p', 'admin', 'pb.ScheduleService', 'FindScheduleById', NULL, NULL, NULL),
	('p', 'admin', 'pb.WorkspaceService', 'CreateWorkspace', NULL, NULL, NULL),
	('p', 'admin', 'pb.WorkspaceService', 'DeleteWorkspaceById', NULL, NULL, NULL),
	('p', 'admin', 'pb.WorkspaceService', 'FindWorkspaces', NULL, NULL, NULL),
	('p', 'admin', 'pb.ImageService', 'UploadImage', NULL, NULL, NULL),
	('p', 'admin', 'pb.ImageService', 'GetImageId', NULL, NULL, NULL),
	('p', 'user', 'pb.AttendanceService', 'CreateAttendance', NULL, NULL, NULL),
	('p', 'user', 'pb.AttendanceService', 'FindAttendances', NULL, NULL, NULL),
	('p', 'user', 'pb.AttendanceService', 'FindAttendanceById', NULL, NULL, NULL),
	('p', 'user', 'pb.CertificateService', 'CreateCertificate', NULL, NULL, NULL),
	('p', 'user', 'pb.CertificateService', 'UpdateCertificateById', NULL, NULL, NULL),
	('p', 'user', 'pb.CertificateService', 'FindCertificates', NULL, NULL, NULL),
	('p', 'user', 'pb.CertificateService', 'FindCertificateById', NULL, NULL, NULL),
	('p', 'user', 'pb.CheckService', 'CheckIn', NULL, NULL, NULL),
	('p', 'user', 'pb.CheckService', 'CheckOut', NULL, NULL, NULL),
	('p', 'user', 'pb.CheckService', 'FindChecks', NULL, NULL, NULL),
	('p', 'user', 'pb.CheckService', 'FindCheckById', NULL, NULL, NULL),
	('p', 'user', 'pb.OfficeService', 'FindOffices', NULL, NULL, NULL),
	('p', 'user', 'pb.OfficeService', 'FindOfficeById', NULL, NULL, NULL),
	('p', 'user', 'pb.ScheduleService', 'FindSchedules', NULL, NULL, NULL),
	('p', 'user', 'pb.ScheduleService', 'FindScheduleById', NULL, NULL, NULL),
	('p', 'user', 'pb.WorkspaceService', 'FindWorkspaces', NULL, NULL, NULL),
	('p', 'user', 'pb.ImageService', 'UploadImage', NULL, NULL, NULL),
	('p', 'user', 'pb.ImageService', 'GetImageId', NULL, NULL, NULL) RETURNING "id";