CREATE TABLE "selfservice_registration_flow_methods" (
"id" TEXT PRIMARY KEY,
"method" TEXT NOT NULL,
"selfservice_registration_flow_id" char(36) NOT NULL,
"config" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (selfservice_registration_flow_id) REFERENCES selfservice_registration_flow_methods (id) ON DELETE cascade
);