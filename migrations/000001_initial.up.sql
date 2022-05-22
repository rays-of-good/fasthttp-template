create type user_role as enum ('user', 'admin', 'super');

create table users (
	id uuid default gen_random_uuid(),
	
	created_at timestamp with time zone not null default current_timestamp,
	updated_at timestamp with time zone not null default current_timestamp,
	deleted_at timestamp with time zone,

	user_role user_role not null default 'user',

	primary key (id)
);
