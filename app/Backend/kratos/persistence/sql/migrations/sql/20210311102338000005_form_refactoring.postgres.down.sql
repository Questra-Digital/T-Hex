CREATE TABLE "selfservice_recovery_flow_methods" (
"id" UUID NOT NULL,
PRIMARY KEY("id"),
"method" VARCHAR (32) NOT NULL,
"selfservice_recovery_flow_id" UUID NOT NULL,
"config" jsonb NOT NULL,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL,
FOREIGN KEY ("selfservice_recovery_flow_id") REFERENCES "selfservice_recovery_flow_methods" ("id") ON DELETE cascade
);