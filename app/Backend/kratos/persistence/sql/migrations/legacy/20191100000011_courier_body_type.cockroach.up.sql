ALTER TABLE "courier_messages" RENAME COLUMN "body" TO "_body_tmp";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "courier_messages" ADD COLUMN "body" text;COMMIT TRANSACTION;BEGIN TRANSACTION;
UPDATE "courier_messages" SET "body" = "_body_tmp";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "courier_messages" ALTER COLUMN "body" SET NOT NULL;COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "courier_messages" DROP COLUMN "_body_tmp";COMMIT TRANSACTION;BEGIN TRANSACTION;