-- Drop table

DROP TABLE pantry.properties;

CREATE TABLE pantry.properties (
	id serial NOT NULL, -- Unique Identifier
	prop_name varchar(255) NOT NULL,
	prop_value text NULL,
	CONSTRAINT properties_pkey PRIMARY KEY (id)
);
