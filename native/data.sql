-- UNTUK CRUD OPERATION
CREATE TABLE account (
    account_id              varchar(20),
    first_name              varchar(200),
    last_name               varchar(200),
    email                   varchar(100),
    password                varchar(100),
    account_create_time     time,
    account_update_time     time,
    is_active_account       boolean,
    activation_code         varchar(200),
    constraint pk_account primary key (account_id)
);

-- UNTUK OTENTIFIKASI DAN OTORISASI
-- SKEMA DATABASE BUAT OTENTIFIKASI DAN OTORISASI TIDAK ADA STANDAR YANG BAKU
-- SAYA HANYA MENCOBA MENDESAIN SENDIRI SKEMANYA SESUAI KEBUTUHAN
CREATE TABLE users (
    user_id         int auto_increment,
    username        varchar(100),
    password        varchar(200),
    active          boolean,
    primary key (user_id)
);

CREATE TABLE roles (
    role_id         int auto_increment,
    role_name       varchar(20),
    primary key (role_id)
);

CREATE TABLE permissions (
    permissions_id       int auto_increment,
    permissions_name     varchar(20),
    primary key (permissions_id)
);

CREATE TABLE user_role (
    user_role_id        int auto_increment,
    user_id             int,
    role_id             int,
    primary key (user_role_id),
    constraint fk_uru foreign key (user_id) references users(user_id),
    constraint fk_urr foreign key (role_id) references roles(role_id)
);

CREATE TABLE role_permission (
    user_permission_id  int auto_increment,
    role_id             int,
    permission_id       int,
    primary key (user_permission_id),
    constraint fk_rpr foreign key (role_id) references roles(role_id),
    constraint fk_rpp foreign key (permission_id) references permissions(permissions_id)
);

