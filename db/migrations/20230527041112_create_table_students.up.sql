create table students(
    id int(11) not null auto_increment,
    name varchar(200) not null,
    age int(2) not null,
    date_created datetime not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    created_by varchar(200) not null,
    primary key (id)
);