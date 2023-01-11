CREATE TABLE "public"."friends" ("id" serial NOT NULL, "friend_id" integer NOT NULL, "user_id" integer NOT NULL, "room_id" integer NOT NULL, "is_block" boolean NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") );
CREATE OR REPLACE FUNCTION "public"."set_current_timestamp_updated_at"()
RETURNS TRIGGER AS $$
DECLARE
  _new record;
BEGIN
  _new := NEW;
  _new."updated_at" = NOW();
  RETURN _new;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER "set_public_friends_updated_at"
BEFORE UPDATE ON "public"."friends"
FOR EACH ROW
EXECUTE PROCEDURE "public"."set_current_timestamp_updated_at"();
COMMENT ON TRIGGER "set_public_friends_updated_at" ON "public"."friends" 
IS 'trigger to set value of column "updated_at" to current timestamp on row update';
