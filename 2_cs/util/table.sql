create table address_book
(
    id      int auto_increment
        primary key,
    name    varchar(200) not null,
    tel     long          not null,
    address varchar(200) null
);

