CREATE TABLE IF NOT EXISTS "users" (
    "id" serial PRIMARY KEY not null,
    "username" varchar(100) not null UNIQUE,
    "name" varchar(255) not null,
    "password" varchar(255) not null,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE condition AS ENUM ('new', 'second');

CREATE TABLE IF NOT EXISTS "products" (
    "id" serial PRIMARY KEY not null,
    "name" varchar(60) not null,
    "price" int not null,
    "image_url" varchar(255) not null,
    "stock" int not null,
    "condition" condition not null,
    "tags" varchar(255)[] not null,
    "is_purchaseable" boolean not null,
    "purchase_count"  int null,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "bank_accounts" (
    "id" serial PRIMARY KEY not null,
    "user_id" int REFERENCES "users" ("id") ON DELETE CASCADE,
    "bank_name" VARCHAR(50) NOT NULL,
    "bank_account_name" VARCHAR(50) NOT NULL,
    "bank_account_number" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);