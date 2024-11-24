-- providers table
-- id should be bigserial
CREATE TABLE 
"public"."providers" (
    "id" bigserial not null,
    "name" varchar(255) not null,
    "city" varchar(255) not null,
    "region" varchar(255) not null,
    "line_of_business" varchar(255) not null,
    "keywords" varchar(255) not null,
    "userId" bigserial not null,
    constraint "providers_pkey" primary key ("id")
);

insert into "public"."providers" ("city", "keywords", "line_of_business", "name", "region", "userId") values 
('Helsinki', 'personal,training,gym,wellness,health', 'Health and Wellness', 'Healthy Personal Training', 'Uusimaa', 1), 
('Stockholm', 'yoga,wellness,lifestyle,health', 'Health and Wellness', 'Healthy Yoga Instructor', 'Stockholm', 1), 
('Oslo', 'agile,coaching,coach,software,development,teams', 'Information technology', 'Healthy Agile Coaching', 'Akershus', 1);

-- end of providers table

-- users table 

CREATE TABLE
  "public"."users" (
    "id" bigserial not null,
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