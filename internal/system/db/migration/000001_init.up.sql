--
-- PostgreSQL database dum Dumped from database version 14.1 (Debian 14.1-1.pgdg110+1)
-- Dumped by pg_dump version 14.1 (Debian 14.1-1.pgdg110+1)

-- Started on 2022-10-05 16:04:26 +07

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


COMMENT ON EXTENSION postgis IS 'PostGIS geometry and geography spatial types and functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.address_locations (
    id bigint NOT NULL,
    confirm boolean DEFAULT false,
    location public.geometry
);


ALTER TABLE public.address_locations OWNER TO root;

CREATE SEQUENCE public.address_locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.address_locations_id_seq OWNER TO root;


ALTER SEQUENCE public.address_locations_id_seq OWNED BY public.address_locations.id;


CREATE TABLE public.addresses (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    user_id bigint,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    city_id bigint,
    district_id bigint,
    state_id bigint,
    country_id bigint,
    address_line_1 character varying(255) DEFAULT ''::character varying NOT NULL,
    address_line_2 character varying(255),
    location_id bigint
);


ALTER TABLE public.addresses OWNER TO root;

CREATE SEQUENCE public.addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.addresses_id_seq OWNER TO root;


ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


CREATE TABLE public.audit_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    application_name character varying(128),
    user_id bigint,
    username character varying(255),
    "impersonator_user_id " bigint,
    operation_name character varying(128),
    operation_method character varying(128),
    request_method character varying(128),
    request_url character varying(255),
    request_id character varying(255),
    request_body text,
    start_time timestamp with time zone,
    latency bigint,
    location character varying(255),
    ip_address character varying(128),
    device_id character varying(128),
    device_name character varying(128),
    user_agent character varying(128),
    response_body text,
    status_code smallint,
    error character varying(255),
    remark character varying(255)
);


ALTER TABLE public.audit_logs OWNER TO root;

COMMENT ON COLUMN public.audit_logs.latency IS 'Milliseconds';


CREATE SEQUENCE public.audit_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.audit_logs_id_seq OWNER TO root;

ALTER SEQUENCE public.audit_logs_id_seq OWNED BY public.audit_logs.id;


CREATE TABLE public.casbin_rules (
    id bigint NOT NULL,
    ptype character varying(100),
    v0 character varying(100),
    v1 character varying(100),
    v2 character varying(100),
    v3 character varying(100),
    v4 character varying(100),
    v5 character varying(100)
);


ALTER TABLE public.casbin_rules OWNER TO root;

CREATE SEQUENCE public.casbin_rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.casbin_rules_id_seq OWNER TO root;

ALTER SEQUENCE public.casbin_rules_id_seq OWNED BY public.casbin_rules.id;


CREATE TABLE public.menu_roles (
    role_id bigint NOT NULL,
    menu_id bigint NOT NULL
);


ALTER TABLE public.menu_roles OWNER TO root;

CREATE TABLE public.menus (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    slug character varying(128) DEFAULT ''::character varying NOT NULL,
    parent_id bigint,
    path character varying(255),
    hidden boolean,
    component character varying(255),
    sort smallint DEFAULT 0,
    menu_level smallint DEFAULT 0,
    type character varying(10),
    title character varying(255),
    icon character varying(128)
);


ALTER TABLE public.menus OWNER TO root;

CREATE SEQUENCE public.menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menus_id_seq OWNER TO root;

ALTER SEQUENCE public.menus_id_seq OWNED BY public.menus.id;

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    resource_id bigint NOT NULL,
    role_id bigint,
    user_id bigint,
    is_granted boolean
);


ALTER TABLE public.permissions OWNER TO root;

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.permissions_id_seq OWNER TO root;

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


CREATE TABLE public.resources (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    name character varying(100) DEFAULT ''::character varying NOT NULL,
    slug character varying(100) DEFAULT ''::character varying NOT NULL,
    object character varying(100) DEFAULT ''::character varying,
    action character varying(100) DEFAULT ''::character varying,
    description text,
    parent_id bigint
);


ALTER TABLE public.resources OWNER TO root;

CREATE SEQUENCE public.resources_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.resources_id_seq OWNER TO root;

ALTER SEQUENCE public.resources_id_seq OWNED BY public.resources.id;


CREATE TABLE public.roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    name character varying(100) DEFAULT ''::character varying NOT NULL,
    slug character varying(100) DEFAULT ''::character varying NOT NULL,
    is_static boolean,
    is_default boolean
);


ALTER TABLE public.roles OWNER TO root;

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO root;


ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


CREATE TABLE public.security_logs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    username character varying(255),
    location character varying(255),
    ip_address character varying(128),
    device_id character varying(128),
    device_name character varying(128),
    browser character varying(128),
    platform character varying(128),
    os character varying(128),
    user_agent character varying(128),
    "time" timestamp with time zone,
    status_code smallint,
    message character varying(255)
);


ALTER TABLE public.security_logs OWNER TO root;

CREATE SEQUENCE public.security_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.security_logs_id_seq OWNER TO root;


ALTER SEQUENCE public.security_logs_id_seq OWNED BY public.security_logs.id;


CREATE TABLE public.settings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    name character varying(128) DEFAULT ''::character varying NOT NULL,
    value text,
    provider_key character varying(4) DEFAULT ''::character varying NOT NULL,
    provider_name character varying(128) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.settings OWNER TO root;

CREATE SEQUENCE public.settings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.settings_id_seq OWNER TO root;


ALTER SEQUENCE public.settings_id_seq OWNED BY public.settings.id;


CREATE TABLE public.user_roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    user_id bigint,
    role_id bigint
);


ALTER TABLE public.user_roles OWNER TO root;

CREATE SEQUENCE public.user_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_roles_id_seq OWNER TO root;


ALTER SEQUENCE public.user_roles_id_seq OWNED BY public.user_roles.id;


CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    creator_user_id bigint,
    modified_at timestamp with time zone,
    modifier_user_id bigint,
    deleted_at timestamp with time zone,
    deleter_user_id bigint,
    uuid text,
    username text,
    password text,
    full_name text,
    avatar_path text,
    email text,
    email_verified boolean,
    email_verification_code_hashed text,
    date_of_birth timestamp with time zone,
    phone text,
    phone_verified boolean,
    phone_verification_code_hashed text,
    gender smallint DEFAULT 1,
    is_global_admin boolean
);


ALTER TABLE public.users OWNER TO root;

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO root;


ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


ALTER TABLE ONLY public.address_locations ALTER COLUMN id SET DEFAULT nextval('public.address_locations_id_seq'::regclass);


ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


ALTER TABLE ONLY public.audit_logs ALTER COLUMN id SET DEFAULT nextval('public.audit_logs_id_seq'::regclass);


ALTER TABLE ONLY public.casbin_rules ALTER COLUMN id SET DEFAULT nextval('public.casbin_rules_id_seq'::regclass);


ALTER TABLE ONLY public.menus ALTER COLUMN id SET DEFAULT nextval('public.menus_id_seq'::regclass);


ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


ALTER TABLE ONLY public.resources ALTER COLUMN id SET DEFAULT nextval('public.resources_id_seq'::regclass);


ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


ALTER TABLE ONLY public.security_logs ALTER COLUMN id SET DEFAULT nextval('public.security_logs_id_seq'::regclass);


ALTER TABLE ONLY public.settings ALTER COLUMN id SET DEFAULT nextval('public.settings_id_seq'::regclass);


ALTER TABLE ONLY public.user_roles ALTER COLUMN id SET DEFAULT nextval('public.user_roles_id_seq'::regclass);


ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


ALTER TABLE ONLY public.address_locations
    ADD CONSTRAINT address_locations_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT audit_logs_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.casbin_rules
    ADD CONSTRAINT casbin_rules_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.menu_roles
    ADD CONSTRAINT menu_roles_pkey PRIMARY KEY (role_id, menu_id);


ALTER TABLE ONLY public.menus
    ADD CONSTRAINT menus_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.resources
    ADD CONSTRAINT resources_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.security_logs
    ADD CONSTRAINT security_logs_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


CREATE UNIQUE INDEX idx_addresses_user_id_is_primary ON public.addresses USING btree (user_id, user_id);


CREATE UNIQUE INDEX idx_casbin_rules ON public.casbin_rules USING btree (ptype, v0, v1, v2, v3, v4, v5);


CREATE INDEX idx_menus_parent_id ON public.menus USING btree (parent_id);


CREATE INDEX idx_menus_path ON public.menus USING btree (path);


CREATE UNIQUE INDEX idx_menus_slug ON public.menus USING btree (slug);


CREATE INDEX idx_menus_sort ON public.menus USING btree (sort);


CREATE INDEX idx_menus_type ON public.menus USING btree (type);


CREATE INDEX idx_permissions_resource_id ON public.permissions USING btree (resource_id);


CREATE INDEX idx_permissions_role_id ON public.permissions USING btree (role_id);


CREATE INDEX idx_permissions_user_id ON public.permissions USING btree (user_id);


CREATE UNIQUE INDEX idx_resources_object_action ON public.resources USING btree (object, action) WHERE (((object)::text <> ''::text) AND ((action)::text <> ''::text));


CREATE INDEX idx_resources_parent_id ON public.resources USING btree (parent_id);


CREATE UNIQUE INDEX idx_resources_slug ON public.resources USING btree (slug);


CREATE UNIQUE INDEX idx_roles_slug ON public.roles USING btree (slug);


CREATE UNIQUE INDEX "idx_settings_name_providerName_providerKey" ON public.settings USING btree (name, provider_key, provider_name);


ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT fk_addresses_location FOREIGN KEY (location_id) REFERENCES public.address_locations(id);


ALTER TABLE ONLY public.menu_roles
    ADD CONSTRAINT fk_menu_roles_menu FOREIGN KEY (menu_id) REFERENCES public.menus(id);


ALTER TABLE ONLY public.menu_roles
    ADD CONSTRAINT fk_menu_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT fk_permissions_user FOREIGN KEY (user_id) REFERENCES public.users(id);


ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT fk_resources_permissions FOREIGN KEY (resource_id) REFERENCES public.resources(id);


ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT fk_roles_permissions FOREIGN KEY (role_id) REFERENCES public.roles(id);


ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_roles_user_roles FOREIGN KEY (role_id) REFERENCES public.roles(id);


ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT fk_users_addresses FOREIGN KEY (user_id) REFERENCES public.users(id);


ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_users_user_roles FOREIGN KEY (user_id) REFERENCES public.users(id);



