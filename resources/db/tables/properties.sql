-- Drop table

DROP TABLE pantry.properties;

CREATE TABLE pantry.properties (
	id serial NOT NULL, -- Unique Identifier
	status varchar(1) NOT NULL,
	prop_category varchar(255) NOT NULL, -- Divides the Properties into distinct Categories
	prop_name varchar(255) NOT NULL, -- Property name
	prop_value text NULL, -- Property value
	prop_description text NULL, -- Property value
	CONSTRAINT properties_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE pantry.properties IS 'Application Essential Parameters';

-- Column comments
COMMENT ON COLUMN pantry.properties.id IS 'Unique Identifier';
COMMENT ON COLUMN pantry.properties.prop_category IS 'Divides the Properties into distinct Categories';
COMMENT ON COLUMN pantry.properties.prop_name IS 'Property name';
COMMENT ON COLUMN pantry.properties.prop_value IS 'Property value';
