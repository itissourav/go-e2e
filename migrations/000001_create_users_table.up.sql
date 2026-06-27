-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE IF Not Exists public.users (
	id int8 GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE) NOT NULL,
	first_name text NULL,
	last_name text NULL,
	email text NULL,
	is_active bool NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);