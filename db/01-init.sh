#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $POSTGRES_USER;
  \connect $APP_DB_NAME $POSTGRES_USER
  ALTER ROLE $POSTGRES_USER SUPERUSER;
  BEGIN;
    
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS course (
    "id" uuid not null default uuid_generate_v4(),
    "course_name" text,
    "course_description" text,
    "author_fullname" text,   
    "author_firstname" text , 
    "author_lastname" text ,
    "course_no" serial,
    "uts" timestamp default current_timestamp, 
    PRIMARY KEY ("id")
);

  COMMIT;
EOSQL