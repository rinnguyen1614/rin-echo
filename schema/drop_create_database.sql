drop database if exists rin_admin with (force);
create database rin_admin with owner admin;
CREATE EXTENSION postgis;
// SELECT postgis_full_version();
// install postgis in ubuntu, debian... 
// sudo apt install postgis
// sudo apt install postgresql-14-postgis
// sudo apt install postgresql-14-postgis-scripts
// sudo apt install postgresql-14-pgrouting 