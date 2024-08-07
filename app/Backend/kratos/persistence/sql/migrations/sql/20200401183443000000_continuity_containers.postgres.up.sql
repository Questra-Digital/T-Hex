CREATE TABLE "continuity_containers" (
"id" UUID NOT NULL,
PRIMARY KEY("id"),
"identity_id" UUID,
"name" VARCHAR (255) NOT NULL,
"payload" jsonb,
"expires_at" timestamp NOT NULL,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL,
FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") ON DELETE cascade
);