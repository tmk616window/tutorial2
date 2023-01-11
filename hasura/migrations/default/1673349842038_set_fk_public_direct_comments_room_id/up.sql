alter table "public"."direct_comments"
  add constraint "direct_comments_room_id_fkey"
  foreign key ("room_id")
  references "public"."rooms"
  ("id") on update restrict on delete restrict;
