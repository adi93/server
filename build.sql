
create table task ( 
	rowid INTEGER primary key AUTOINCREMENT,
	title TEXT not null,
	description TEXT default "",
	dueDate integer not null default CURRENT_TIMESTAMP,
	status TEXT not null default "pending",
	priority TINYINT not null default 0,
	effort TEXT not null default "24h",
	created integer default CURRENT_TIMESTAMP,
	constraint unique_title_name unique (title)
);

insert into task (title, description) values ("Hello world", "My first task");

