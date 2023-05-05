alter table "public"."relationships" add constraint "relationships_follow_id_following_id_room_id_key" unique ("follow_id", "following_id", "room_id");
