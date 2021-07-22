set time_zone = '+03:00';

drop table if exists messages;
drop table if exists user_chat_relation;
drop table if exists chats;
drop table if exists users;

create table users (
    id serial primary key,
    username varchar not null,
    created_at timestamp default current_timestamp
);

create table chats (
    id serial primary key,
    name varchar not null,
    created_at timestamp default current_timestamp
);

create table user_chat_relation (
    chat_id int not null,
    user_id int not null,
    foreign key (chat_id) references chats(id),
    foreign key (user_id) references users(id)
);

create table messages (
    id serial primary key,
    chat int not null,
    author int not null,
    text varchar not null,
    created_at timestamp default current_timestamp,
    foreign key (chat) references chats(id),
    foreign key (author) references users(id)
);

