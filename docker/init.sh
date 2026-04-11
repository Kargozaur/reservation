#!/bin/bash

until pg_isready -U "$POSTGRES_USER"; do
  sleep 1
done

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d postgres <<'SQL'

CREATE EXTENSION IF NOT EXISTS dblink;

DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = 'user_service') THEN
      RAISE NOTICE 'Database already exists';
   ELSE
      PERFORM dblink_exec('dbname=' || current_database()
                        , 'CREATE DATABASE user_service');
   END IF;
END
$do$;

DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = 'reservation_service') THEN
      RAISE NOTICE 'Database already exists';
   ELSE
      PERFORM dblink_exec('dbname=' || current_database()
                        , 'CREATE DATABASE reservation_service');
   END IF;
END
$do$;

SQL
