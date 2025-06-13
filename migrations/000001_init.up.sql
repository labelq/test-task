CREATE TABLE users (
                       id serial primary key,
                       name varchar not null,
                       password varchar not null,
                       points integer default 0,
                       referrer integer references users(id) on delete set null
);

CREATE TABLE tasks (
                       id serial primary key,
                       name varchar not null,
                       instruction varchar not null,
                       point integer default 50
);

CREATE TABLE users_tasks (
                             user_id integer references users(id) on delete cascade,
                             task_id integer references tasks(id) on delete cascade,
                             primary key (user_id, task_id)
)