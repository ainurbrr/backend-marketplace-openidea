CREATE TABLE IF NOT EXISTS "users" (
    "id" serial PRIMARY KEY not null,
    "username" varchar(100) not null UNIQUE,
    "name" varchar(255) not null,
    "password" varchar(255) not null,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)