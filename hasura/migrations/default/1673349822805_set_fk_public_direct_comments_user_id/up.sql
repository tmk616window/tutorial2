alter table "public"."direct_comments"
  add constraint "direct_comments_user_id_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
