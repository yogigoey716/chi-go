-- Add up migration script here
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS chigo (
    "id" SERIAL,
    "blog_name" VARCHAR NOT NULL,
    "blog_details" VARCHAR NOT NULL,
    "blog_description" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
