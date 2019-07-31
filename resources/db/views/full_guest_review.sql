-- Drop table

-- DROP TABLE pantry.properties

CREATE TABLE pantry.properties (
	"ID" int4 NOT NULL,
	"name" varchar(255) NULL,
	value text NULL,
	CONSTRAINT properties_pkey PRIMARY KEY ("ID")
);
