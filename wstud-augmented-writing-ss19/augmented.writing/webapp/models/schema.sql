Create database augmented_writing;

CREATE TABLE IF NOT EXISTS "users" (
    "id" serial NOT NULL PRIMARY KEY,
    "first_name" text NOT NULL DEFAULT '' ,
    "last_name" text NOT NULL DEFAULT '' ,
    "user_name" text NOT NULL DEFAULT ''  UNIQUE,
    "email" text NOT NULL DEFAULT ''  UNIQUE,
    "password" text NOT NULL DEFAULT ''
);