-- Drop table

DROP TABLE users;

CREATE TABLE users (
	id serial NOT NULL, -- Unique Identifier
	status varchar(1) NOT NULL,
	first_name varchar(40) NOT NULL, -- Divides the Properties into distinct Categories
	last_name varchar(40) NOT NULL, -- Property name
	username varchar(40) NOT NULL, -- Property value
	password varchar(40) NOT NULL, -- Property value
  notes text NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
