-- providers table
CREATE TABLE 
"public"."providers" (
    "id" serial not null,
    "name" varchar(255) not null,
    constraint "providers_pkey" primary key ("id")
);

insert into "public"."providers" ("id", "name") values (1, 'Healthy Personal Training');
insert into "public"."providers" ("id", "name") values (2, 'Healthy Yoga Instructor');
insert into "public"."providers" ("id", "name") values (3, 'Healthy Agile Coaching');


-- end of providers table

-- users table 

create table
  "public"."users" (
    "id" serial not null,
    "username" varchar(255) not null default NOW(),
    "email" varchar(255) not null,
    "password" text not null,
    "salt" varchar(255) not null,
    constraint "user_pkey" primary key ("id")
  );
  
  -- end user table

CREATE TABLE
  "public"."sessions" (
    "token" varchar(255) not null,
    "expires" timestamp without time zone not null default now(),
    "user_id" integer not null,
    constraint "sessions_pkey" primary key ("token","user_id")
);