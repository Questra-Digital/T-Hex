DROP INDEX IF EXISTS "identity_verifiable_addresses_code_uq_idx";COMMIT TRANSACTION;BEGIN TRANSACTION;
DROP INDEX IF EXISTS "identity_verifiable_addresses_code_idx";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_verifiable_addresses" DROP COLUMN "code";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_verifiable_addresses" DROP COLUMN "expires_at";COMMIT TRANSACTION;BEGIN TRANSACTION;