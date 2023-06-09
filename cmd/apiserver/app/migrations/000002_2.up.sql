create table sender (
    id serial primary key,
    name varchar(255) not null,
    phone_no varchar(20) not null,
    identification_no varchar(255) not null,
    address varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table receiver (
    id serial primary key,
    name varchar(255) not null,
    phone_no varchar(20) not null,
    identification_no varchar(255) not null,
    country varchar(50) not null,
    address varchar(255) not null,
    postal_code varchar(12) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table resi (
    id serial primary key,
    unique_id varchar(255) not null,
    sender_id integer not null,
    receiver_id integer not null,
    service varchar(30) not null,
    cost bigint not null,
    package_id integer not null,
    user_id integer not null,
    branch_id integer not null,
    status varchar(20) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table package (
    id serial primary key,
    total_weight integer not null,
    total_price_value bigint not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table package_detail (
    id serial primary key,
    package_id integer not null,
    package_type varchar(55) not null,
    name varchar(255) not null,
    pieces integer not null,
    price_value bigint not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table package_type (
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table service_type (
    id serial primary key,
    name varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);