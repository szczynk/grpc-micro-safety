-- create "casbin_rule" table
CREATE TABLE "public"."casbin_rule" ("id" bigserial NOT NULL, "ptype" character varying(100) NULL, "v0" character varying(100) NULL, "v1" character varying(100) NULL, "v2" character varying(100) NULL, "v3" character varying(100) NULL, "v4" character varying(100) NULL, "v5" character varying(100) NULL, PRIMARY KEY ("id"));
-- create index "idx_casbin_rule" to table: "casbin_rule"
CREATE UNIQUE INDEX "idx_casbin_rule" ON "public"."casbin_rule" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5");

INSERT INTO "casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES
	('1', 'g', '028de3d9-2f46-4cb2-92f0-2ca723e31862', 'admin', NULL, NULL, NULL, NULL),
	('2', 'p', 'admin', 'pb.UserService', 'RefreshAccessToken', NULL, NULL, NULL),
	('3', 'p', 'admin', 'pb.UserService', 'UpdateMe', NULL, NULL, NULL),
	('4', 'p', 'admin', 'pb.UserService', 'ChangeEmail', NULL, NULL, NULL),
	('5', 'p', 'admin', 'pb.UserService', 'Logout', NULL, NULL, NULL),
	('6', 'p', 'admin', 'pb.UserService', 'GetMe', NULL, NULL, NULL),
	('7', 'p', 'admin', 'pb.UserService', 'CreateUser', NULL, NULL, NULL),
	('8', 'p', 'admin', 'pb.UserService', 'UpdateUserById', NULL, NULL, NULL),
	('9', 'p', 'admin', 'pb.UserService', 'DeleteUserById', NULL, NULL, NULL),
	('10', 'p', 'admin', 'pb.UserService', 'FindUsers', NULL, NULL, NULL),
	('11', 'p', 'admin', 'pb.UserService', 'FindUserById', NULL, NULL, NULL),
	('12', 'p', 'admin', 'pb.RoleService', 'CreateRole', NULL, NULL, NULL),
	('13', 'p', 'admin', 'pb.RoleService', 'DeleteRoleById', NULL, NULL, NULL),
	('14', 'p', 'admin', 'pb.RoleService', 'FindRoles', NULL, NULL, NULL),
	('15', 'p', 'admin', 'pb.ServiceService', 'CreateService', NULL, NULL, NULL),
	('16', 'p', 'admin', 'pb.ServiceService', 'DeleteServiceById', NULL, NULL, NULL),
	('17', 'p', 'admin', 'pb.ServiceService', 'FindServices', NULL, NULL, NULL),
	('18', 'p', 'admin', 'pb.MethodService', 'CreateMethod', NULL, NULL, NULL),
	('19', 'p', 'admin', 'pb.MethodService', 'DeleteMethodById', NULL, NULL, NULL),
	('20', 'p', 'admin', 'pb.MethodService', 'FindMethods', NULL, NULL, NULL),
	('21', 'p', 'admin', 'pb.PolicyService', 'CreatePolicy', NULL, NULL, NULL),
	('22', 'p', 'admin', 'pb.PolicyService', 'DeletePolicy', NULL, NULL, NULL),
	('23', 'p', 'admin', 'pb.PolicyService', 'FindPolicies', NULL, NULL, NULL),
	('24', 'p', 'user', 'pb.UserService', 'RefreshAccessToken', NULL, NULL, NULL),
	('25', 'p', 'user', 'pb.UserService', 'UpdateMe', NULL, NULL, NULL),
	('26', 'p', 'user', 'pb.UserService', 'ChangeEmail', NULL, NULL, NULL),
	('27', 'p', 'user', 'pb.UserService', 'Logout', NULL, NULL, NULL),
	('28', 'p', 'user', 'pb.UserService', 'GetMe', NULL, NULL, NULL),
	('29', 'p', 'admin', 'pb.AttendanceService', 'CreateAttendance', NULL, NULL, NULL),
	('30', 'p', 'admin', 'pb.AttendanceService', 'UpdateAttendanceById', NULL, NULL, NULL),
	('31', 'p', 'admin', 'pb.AttendanceService', 'DeleteAttendanceById', NULL, NULL, NULL),
	('32', 'p', 'admin', 'pb.AttendanceService', 'FindAttendances', NULL, NULL, NULL),
	('33', 'p', 'admin', 'pb.AttendanceService', 'FindAttendanceById', NULL, NULL, NULL),
	('34', 'p', 'admin', 'pb.CertificateService', 'CreateCertificate', NULL, NULL, NULL),
	('35', 'p', 'admin', 'pb.CertificateService', 'UpdateCertificateById', NULL, NULL, NULL),
	('36', 'p', 'admin', 'pb.CertificateService', 'DeleteCertificateById', NULL, NULL, NULL),
	('37', 'p', 'admin', 'pb.CertificateService', 'FindCertificates', NULL, NULL, NULL),
	('38', 'p', 'admin', 'pb.CertificateService', 'FindCertificateById', NULL, NULL, NULL),
	('39', 'p', 'admin', 'pb.CheckService', 'CheckIn', NULL, NULL, NULL),
	('40', 'p', 'admin', 'pb.CheckService', 'CheckOut', NULL, NULL, NULL),
	('41', 'p', 'admin', 'pb.CheckService', 'FindChecks', NULL, NULL, NULL),
	('42', 'p', 'admin', 'pb.CheckService', 'FindCheckById', NULL, NULL, NULL),
	('43', 'p', 'admin', 'pb.OfficeService', 'CreateOffice', NULL, NULL, NULL),
	('44', 'p', 'admin', 'pb.OfficeService', 'UpdateOfficeById', NULL, NULL, NULL),
	('45', 'p', 'admin', 'pb.OfficeService', 'DeleteOfficeById', NULL, NULL, NULL),
	('46', 'p', 'admin', 'pb.OfficeService', 'FindOffices', NULL, NULL, NULL),
	('47', 'p', 'admin', 'pb.OfficeService', 'FindOfficeById', NULL, NULL, NULL),
	('48', 'p', 'admin', 'pb.ScheduleService', 'CreateSchedule', NULL, NULL, NULL),
	('49', 'p', 'admin', 'pb.ScheduleService', 'UpdateScheduleById', NULL, NULL, NULL),
	('50', 'p', 'admin', 'pb.ScheduleService', 'DeleteScheduleById', NULL, NULL, NULL),
	('51', 'p', 'admin', 'pb.ScheduleService', 'FindSchedules', NULL, NULL, NULL),
	('52', 'p', 'admin', 'pb.ScheduleService', 'FindScheduleById', NULL, NULL, NULL),
	('53', 'p', 'admin', 'pb.WorkspaceService', 'CreateWorkspace', NULL, NULL, NULL),
	('54', 'p', 'admin', 'pb.WorkspaceService', 'DeleteWorkspaceById', NULL, NULL, NULL),
	('55', 'p', 'admin', 'pb.WorkspaceService', 'FindWorkspaces', NULL, NULL, NULL),
	('56', 'p', 'admin', 'pb.ImageService', 'UploadImage', NULL, NULL, NULL),
	('57', 'p', 'admin', 'pb.ImageService', 'GetImageId', NULL, NULL, NULL),
	('58', 'p', 'user', 'pb.AttendanceService', 'CreateAttendance', NULL, NULL, NULL),
	('59', 'p', 'user', 'pb.AttendanceService', 'FindAttendances', NULL, NULL, NULL),
	('60', 'p', 'user', 'pb.AttendanceService', 'FindAttendanceById', NULL, NULL, NULL),
	('61', 'p', 'user', 'pb.CertificateService', 'CreateCertificate', NULL, NULL, NULL),
	('62', 'p', 'user', 'pb.CertificateService', 'UpdateCertificateById', NULL, NULL, NULL),
	('63', 'p', 'user', 'pb.CertificateService', 'FindCertificates', NULL, NULL, NULL),
	('64', 'p', 'user', 'pb.CertificateService', 'FindCertificateById', NULL, NULL, NULL),
	('65', 'p', 'user', 'pb.CheckService', 'CheckIn', NULL, NULL, NULL),
	('66', 'p', 'user', 'pb.CheckService', 'CheckOut', NULL, NULL, NULL),
	('67', 'p', 'user', 'pb.CheckService', 'FindChecks', NULL, NULL, NULL),
	('68', 'p', 'user', 'pb.CheckService', 'FindCheckById', NULL, NULL, NULL),
	('69', 'p', 'user', 'pb.OfficeService', 'FindOffices', NULL, NULL, NULL),
	('70', 'p', 'user', 'pb.OfficeService', 'FindOfficeById', NULL, NULL, NULL),
	('71', 'p', 'user', 'pb.ScheduleService', 'FindSchedules', NULL, NULL, NULL),
	('72', 'p', 'user', 'pb.ScheduleService', 'FindScheduleById', NULL, NULL, NULL),
	('73', 'p', 'user', 'pb.WorkspaceService', 'FindWorkspaces', NULL, NULL, NULL),
	('74', 'p', 'user', 'pb.ImageService', 'UploadImage', NULL, NULL, NULL),
	('75', 'p', 'user', 'pb.ImageService', 'GetImageId', NULL, NULL, NULL) RETURNING "id";