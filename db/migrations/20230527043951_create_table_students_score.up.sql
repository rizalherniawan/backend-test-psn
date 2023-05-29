create table student_scores(
    id int(11) not null auto_increment,
    subjects varchar(200) not null,
    student_id int(11) not null,
    score decimal(11,3) not null,
    date_created datetime not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    created_by varchar(200) not null,
    date_modified datetime,
    modified_by varchar(200),
    primary key (id),
    foreign key (student_id) references students(id) on delete cascade
);