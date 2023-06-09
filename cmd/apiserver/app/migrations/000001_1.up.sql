create table "user" (
    id serial primary key,
    name varchar(255) not null,
    password varchar(255) not null,
    email varchar(255) not null unique,
    phone_number varchar(20) not null unique,
    role_id integer,
    branch_id integer not null,
    status varchar(20) not null,
    last_login timestamp,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table "role" (
    id serial primary key,
    name varchar(255) not null,
    status varchar(20) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table page (
    id serial primary key,
    name varchar(255) not null,
    status varchar(20) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table context (
    id serial primary key,
    context_tag varchar(255) not null,
    primary_key integer not null,
    secondary_key integer not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table credential (
    id serial primary key,
    user_id integer not null,
    secret varchar(255) not null,
    status varchar(20) not null,
    expired_at timestamp not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table branch (
    id serial primary key,
    name varchar(255) not null,
    code varchar(5) not null,
    unique_id varchar(255) not null,
    status varchar(20) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table misc (
  id serial primary key,
  name varchar(255) not null,
  value text not null,
  created_at timestamp default now(),
  updated_at timestamp default now()
);

create table country (
  id serial primary key,
  name varchar(255) not null,
  created_at timestamp default now(),
  updated_at timestamp default now()
);