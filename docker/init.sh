#!/bin/bash

until pg_isready -U "$POSTGRES_USER"; do
  sleep 1
done

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d postgres <<'SQL'

CREATE DATABASE user_service;
CREATE DATABASE reservation_service;

SQL
