create table report_categories
(
    id   int auto_increment
        primary key,
    name varchar(256) not null
);

create table schools
(
    id      int auto_increment
        primary key,
    name    varchar(256) null,
    display tinyint(1)   not null,
    constraint schools_name_uindex
        unique (name)
);

create table users
(
    id             int auto_increment
        primary key,
    school_id      int                                                                               not null,
    name           varchar(128)                                                                      not null,
    nickname       varchar(32)                                                                       not null,
    code           int(6)                                    default floor(rand() * 900000 + 100000) null,
    code_generated timestamp                                 default current_timestamp()             null,
    token          varchar(32)                                                                       null,
    type           enum ('student', 'teacher', 'headmaster') default 'student'                       not null,
    banned_to      timestamp                                                                         null,
    verified       tinyint(1)                                default 0                               null,
    date           timestamp                                 default current_timestamp()             null,
    constraint users_nickname_uindex
        unique (nickname),
    constraint users_schools_id_fk
        foreign key (school_id) references schools (id)
);

create table proposals
(
    id        int auto_increment
        primary key,
    user_id   int                                   not null,
    title     varchar(512)                          not null,
    text      text                                  not null,
    anonymous tinyint(1)                            not null,
    date      timestamp default current_timestamp() null,
    constraint proposals_users_id_fk
        foreign key (user_id) references users (id)
);

create table comments
(
    id          int auto_increment
        primary key,
    proposal_id int                                   not null,
    user_id     int                                   not null,
    comment     text                                  not null,
    anonymous   tinyint(1)                            not null,
    date        timestamp default current_timestamp() null,
    constraint comments_proposals_id_fk
        foreign key (proposal_id) references proposals (id),
    constraint comments_users_id_fk
        foreign key (user_id) references users (id)
);

create table comment_likes
(
    id         int auto_increment
        primary key,
    comment_id int                                   not null,
    user_id    int                                   not null,
    feedback   enum ('like', 'dislike')              not null,
    date       timestamp default current_timestamp() null,
    constraint comment_likes_comments_id_fk
        foreign key (comment_id) references comments (id),
    constraint comment_likes_users_id_fk
        foreign key (user_id) references users (id)
);

create table comment_reports
(
    id              int auto_increment
        primary key,
    comment_id      int                                   not null,
    user_id         int                                   not null,
    report_category int                                   not null,
    comment         text                                  null,
    date            timestamp default current_timestamp() null,
    constraint comment_reports_comments_id_fk
        foreign key (comment_id) references comments (id),
    constraint comment_reports_report_categories_id_fk
        foreign key (report_category) references report_categories (id),
    constraint comment_reports_users_id_fk
        foreign key (user_id) references users (id)
);

create table proposal_likes
(
    id          int auto_increment
        primary key,
    proposal_id int                                   not null,
    user_id     int                                   not null,
    feedback    enum ('like', 'dislike')              not null,
    date        timestamp default current_timestamp() null,
    constraint proposal_likes_proposals_id_fk
        foreign key (proposal_id) references proposals (id),
    constraint proposal_likes_users_id_fk
        foreign key (user_id) references users (id)
);

create table proposal_reports
(
    id              int auto_increment
        primary key,
    proposal_id     int                                   not null,
    user_id         int                                   not null,
    report_category int                                   not null,
    comment         text                                  null,
    date            timestamp default current_timestamp() null,
    constraint proposal_reports_proposals_id_fk
        foreign key (proposal_id) references proposals (id),
    constraint proposal_reports_report_categories_id_fk
        foreign key (report_category) references report_categories (id),
    constraint proposal_reports_users_id_fk
        foreign key (user_id) references users (id)
);

