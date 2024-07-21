CREATE TABLE "_selfservice_errors_tmp" (
"id" TEXT PRIMARY KEY,
"errors" TEXT NOT NULL,
"seen_at" DATETIME,
"was_seen" bool NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
"csrf_token" TEXT NOT NULL DEFAULT '',
"nid" char(36)
);