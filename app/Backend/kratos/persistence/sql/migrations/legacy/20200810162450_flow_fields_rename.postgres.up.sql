ALTER TABLE "selfservice_login_flow_methods" RENAME COLUMN "selfservice_login_request_id" TO "selfservice_login_flow_id";
ALTER TABLE "selfservice_registration_flow_methods" RENAME COLUMN "selfservice_registration_request_id" TO "selfservice_registration_flow_id";
ALTER TABLE "selfservice_recovery_flow_methods" RENAME COLUMN "selfservice_recovery_request_id" TO "selfservice_recovery_flow_id";
ALTER TABLE "selfservice_settings_flow_methods" RENAME COLUMN "selfservice_settings_request_id" TO "selfservice_settings_flow_id";