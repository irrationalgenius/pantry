-- Drop table

DROP TABLE roles;

CREATE TABLE roles (
	id serial NOT NULL, -- Unique Identifier
	role_code varchar(4) NOT NULL,
  notes text NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);
