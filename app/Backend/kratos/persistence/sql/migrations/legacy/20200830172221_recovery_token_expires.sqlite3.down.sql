DELETE FROM identity_recovery_tokens WHERE selfservice_recovery_flow_id IS NULL;
DROP INDEX IF EXISTS "identity_recovery_addresses_code_uq_idx";
DROP INDEX IF EXISTS "identity_recovery_addresses_code_idx";
CREATE TABLE "_identity_recovery_tokens_tmp" (
"id" TEXT PRIMARY KEY,
"token" TEXT NOT NULL,
"used" bool NOT NULL DEFAULT 'false',
"used_at" DATETIME,
"identity_recovery_address_id" char(36) NOT NULL,
"selfservice_recovery_flow_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"expires_at" DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00',
"issued_at" DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00',
FOREIGN KEY (identity_recovery_address_id) REFERENCES identity_recovery_addresses (id) ON UPDATE NO ACTION ON DELETE CASCADE,
FOREIGN KEY (selfservice_recovery_flow_id) REFERENCES selfservice_recovery_flows (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE UNIQUE INDEX "identity_recovery_addresses_code_uq_idx" ON "_identity_recovery_tokens_tmp" (token);
CREATE INDEX "identity_recovery_addresses_code_idx" ON "_identity_recovery_tokens_tmp" (token);
INSERT INTO "_identity_recovery_tokens_tmp" (id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, expires_at, issued_at) SELECT id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, expires_at, issued_at FROM "identity_recovery_tokens";
DROP TABLE "identity_recovery_tokens";
ALTER TABLE "_identity_recovery_tokens_tmp" RENAME TO "identity_recovery_tokens";
DROP INDEX IF EXISTS "identity_recovery_addresses_code_uq_idx";
DROP INDEX IF EXISTS "identity_recovery_addresses_code_idx";
CREATE TABLE "_identity_recovery_tokens_tmp" (
"id" TEXT PRIMARY KEY,
"token" TEXT NOT NULL,
"used" bool NOT NULL DEFAULT 'false',
"used_at" DATETIME,
"identity_recovery_address_id" char(36) NOT NULL,
"selfservice_recovery_flow_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"issued_at" DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00',
FOREIGN KEY (identity_recovery_address_id) REFERENCES identity_recovery_addresses (id) ON UPDATE NO ACTION ON DELETE CASCADE,
FOREIGN KEY (selfservice_recovery_flow_id) REFERENCES selfservice_recovery_flows (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE UNIQUE INDEX "identity_recovery_addresses_code_uq_idx" ON "_identity_recovery_tokens_tmp" (token);
CREATE INDEX "identity_recovery_addresses_code_idx" ON "_identity_recovery_tokens_tmp" (token);
INSERT INTO "_identity_recovery_tokens_tmp" (id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, issued_at) SELECT id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, issued_at FROM "identity_recovery_tokens";

DROP TABLE "identity_recovery_tokens";
ALTER TABLE "_identity_recovery_tokens_tmp" RENAME TO "identity_recovery_tokens";
DROP INDEX IF EXISTS "identity_recovery_addresses_code_uq_idx";
DROP INDEX IF EXISTS "identity_recovery_addresses_code_idx";
CREATE TABLE "_identity_recovery_tokens_tmp" (
"id" TEXT PRIMARY KEY,
"token" TEXT NOT NULL,
"used" bool NOT NULL DEFAULT 'false',
"used_at" DATETIME,
"identity_recovery_address_id" char(36) NOT NULL,
"selfservice_recovery_flow_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (identity_recovery_address_id) REFERENCES identity_recovery_addresses (id) ON UPDATE NO ACTION ON DELETE CASCADE,
FOREIGN KEY (selfservice_recovery_flow_id) REFERENCES selfservice_recovery_flows (id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE UNIQUE INDEX "identity_recovery_addresses_code_uq_idx" ON "_identity_recovery_tokens_tmp" (token);
CREATE INDEX "identity_recovery_addresses_code_idx" ON "_identity_recovery_tokens_tmp" (token);
INSERT INTO "_identity_recovery_tokens_tmp" (id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at) SELECT id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at FROM "identity_recovery_tokens";

DROP TABLE "identity_recovery_tokens";
ALTER TABLE "_identity_recovery_tokens_tmp" RENAME TO "identity_recovery_tokens";