CREATE TABLE "identity_verifiable_addresses" (
"id" UUID NOT NULL,
PRIMARY KEY("id"),
"code" VARCHAR (32) NOT NULL,
"status" VARCHAR (16) NOT NULL,
"via" VARCHAR (16) NOT NULL,
"verified" bool NOT NULL,
"value" VARCHAR (400) NOT NULL,
"verified_at" timestamp,
"expires_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
"identity_id" UUID NOT NULL,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL,
FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") ON DELETE cascade
);
CREATE UNIQUE INDEX "identity_verifiable_addresses_code_uq_idx" ON "identity_verifiable_addresses" (code);
CREATE INDEX "identity_verifiable_addresses_code_idx" ON "identity_verifiable_addresses" (code);
CREATE UNIQUE INDEX "identity_verifiable_addresses_status_via_uq_idx" ON "identity_verifiable_addresses" (via, value);
CREATE INDEX "identity_verifiable_addresses_status_via_idx" ON "identity_verifiable_addresses" (via, value);
CREATE TABLE "selfservice_verification_requests" (
"id" UUID NOT NULL,
PRIMARY KEY("id"),
"request_url" VARCHAR (2048) NOT NULL,
"issued_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
"expires_at" timestamp NOT NULL,
"form" jsonb NOT NULL,
"via" VARCHAR (16) NOT NULL,
"csrf_token" VARCHAR (255) NOT NULL,
"success" bool NOT NULL,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL
);