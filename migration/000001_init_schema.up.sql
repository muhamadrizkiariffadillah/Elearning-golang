CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar,
  "username" varchar UNIQUE,
  "email" varchar UNIQUE,
  "hash_password" varchar,
  "role" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "courses" (
  "id" bigserial PRIMARY KEY,
  "course_name" varchar,
  "course_image_url" varchar,
  "short_description" text,
  "price" int,
  "discount_percent" smallint,
  "final_price" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "sub_course" (
  "id" bigserial PRIMARY KEY,
  "course_id" int,
  "metadata_url" varchar,
  "description" text,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "membership_price" (
  "id" bigserial PRIMARY KEY,
  "duration" int,
  "benefits" text,
  "price" int,
  "discount_percent" smallint,
  "final_price" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "membership" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "start_at" timestamp,
  "end_at" timestamp
);

CREATE TABLE "user_courses" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "course_id" int,
  "created_at" timestamp
);

CREATE TABLE "user_progress" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "course_id" int,
  "sub_course_id" int,
  "is_complete" bool,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "course_id" int,
  "membership_id" int,
  "payment_goal" varchar,
  "code" varchar,
  "total_payment" int,
  "payment_url" varchar,
  "status" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

COMMENT ON COLUMN "users"."role" IS 'default is student';

COMMENT ON COLUMN "courses"."discount_percent" IS '1-100';

COMMENT ON COLUMN "membership_price"."duration" IS '3/6/12';

COMMENT ON COLUMN "membership_price"."benefits" IS 'a,b,c';

COMMENT ON COLUMN "membership_price"."discount_percent" IS '1-100';

COMMENT ON COLUMN "membership"."id" IS 'create when a user signup';

COMMENT ON COLUMN "user_progress"."is_complete" IS 'false';

COMMENT ON COLUMN "transactions"."payment_goal" IS 'membership/course';

COMMENT ON COLUMN "transactions"."code" IS 'CRS-CRSID-USERID/MBR-DURATION-USERID';

COMMENT ON COLUMN "transactions"."status" IS 'paid/pending/fail';

ALTER TABLE "sub_course" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "membership" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_courses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_courses" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "user_progress" ADD FOREIGN KEY ("sub_course_id") REFERENCES "sub_course" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("membership_id") REFERENCES "membership" ("id");