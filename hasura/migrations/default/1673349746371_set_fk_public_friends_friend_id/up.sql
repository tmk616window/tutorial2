alter table "public"."friends"
  add constraint "friends_friend_id_fkey"
  foreign key ("friend_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
