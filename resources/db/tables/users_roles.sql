-- Drop table

DROP TABLE users_roles;

CREATE TABLE users_roles (
	user_id serial NOT NULL, -- Unique Identifier
	role_id varchar(4) NOT NULL,
	CONSTRAINT users_roles_pkey PRIMARY KEY (user_id, role_id)
);
