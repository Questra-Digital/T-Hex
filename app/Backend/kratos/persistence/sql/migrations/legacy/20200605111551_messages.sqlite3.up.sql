ALTER TABLE "selfservice_verification_requests" ADD COLUMN "messages" TEXT;
ALTER TABLE "selfservice_login_requests" ADD COLUMN "messages" TEXT;
ALTER TABLE "selfservice_registration_requests" ADD COLUMN "messages" TEXT;