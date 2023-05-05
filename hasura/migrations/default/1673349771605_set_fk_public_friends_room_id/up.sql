alter table "public"."friends"
  add constraint "friends_room_id_fkey"
  foreign key ("room_id")
  references "public"."rooms"
  ("id") on update restrict on delete restrict;
