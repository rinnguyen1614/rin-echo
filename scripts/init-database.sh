#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER admin WITH PASSWORD 'anhnguyen!@0809';
    CREATE DATABASE rin_admin;
    GRANT ALL PRIVILEGES ON DATABASE rin_admin TO admin;
EOSQL

# Add extension
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "rin_admin" <<-EOSQL
   CREATE EXTENSION postgis;
EOSQL
