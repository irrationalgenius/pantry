
DROP TABLE pantry.guests_archive;

CREATE TABLE pantry.guests_archive (
  id int4 NOT NULL, -- Unique Identifier
	date_enrolled date NOT NULL, -- The date of the first pantry visit
	status varchar(10) NULL, -- Current status [A=Active; I=Inactive; W=Waiting; X=Archive]
	first_name varchar(40) NOT NULL, -- First name
	last_name varchar(40) NOT NULL, -- Last name
	gender varchar(6) NULL, -- Gender [M=Male; F=Female]
	unit_num varchar(10) NULL, -- Number of the adress unit: street, apt, etc.
	st_address varchar(50) NULL, -- Current Street Address
	state varchar(2) NULL, -- Current State code
	city varchar(35) NULL, -- Current City
	zip varchar(11) NULL, -- 5 digit Zip code
	tel_num varchar(12) NULL, -- Current Primary contact number
	email varchar(75) NULL, -- Current Primary email address
	count_children int4 NULL, -- Total Household child count
	count_adults int4 NULL, -- Total Household adult count
	worship_place varchar(80) NULL, -- Primary place of worship
	is_member varchar(1) NULL, -- Is a Member of the Laurel Church of Christ
	is_baptized varchar(1) NULL, -- Is Baptized into Jesus Christ
	is_espanol varchar(1) NULL, -- Is Spanish speaking, Hispanic
	is_unemployed varchar(1) NULL, -- Is looking for employment
	is_homeless varchar(1) NULL, -- Is the Guest Homeless? Does not have a Home? Stable roof over the head
	is_family varchar(1) NULL, -- Has at least 2 people that lives with guest
	is_contact_ok varchar(1) NULL, -- Is okay for Contact? [0 = No, -1 = Yes]
	allergies text NULL, -- Note of known allergies pertaining to this guest
	notes text NULL, -- Additional remarks
	last_date_updated date NOT NULL, -- Last date that this record of this Guest was updated on the original table

  -- Archival information
  archive_count int NOT NULL,
  archive_last_date_updated date NOT NULL DEFAULT 'now'::text::date,
  archive_method varchar(1) NOT NULL, -- How was this record archived? [D=Deleted, X=Inserted by Procedure]

  CONSTRAINT guests_archive_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE pantry.guests_archive IS 'The Archives of Laurel Pantry Guests';

COMMENT ON COLUMN pantry.guests_archive.id IS 'Unique Identifier';
COMMENT ON COLUMN pantry.guests_archive.date_enrolled IS 'The date of the first pantry visit';
COMMENT ON COLUMN pantry.guests_archive.status IS 'Current status [A=Active; I=Inactive; W=Waiting; X=Archive]';
COMMENT ON COLUMN pantry.guests_archive.first_name IS 'First name';
COMMENT ON COLUMN pantry.guests_archive.last_name IS 'Last name';
COMMENT ON COLUMN pantry.guests_archive.gender IS 'Gender [M=Male; F=Female]';
COMMENT ON COLUMN pantry.guests_archive.st_address IS 'Current Street Address';
COMMENT ON COLUMN pantry.guests_archive.unit_num IS 'Number of the adress unit: street, apt, etc.';
COMMENT ON COLUMN pantry.guests_archive.state IS 'Current State code';
COMMENT ON COLUMN pantry.guests_archive.city IS 'Current City';
COMMENT ON COLUMN pantry.guests_archive.zip IS '5 digit Zip code';
COMMENT ON COLUMN pantry.guests_archive.tel_num IS 'Current Primary contact number';
COMMENT ON COLUMN pantry.guests_archive.email IS 'Current Primary email address';
COMMENT ON COLUMN pantry.guests_archive.count_children IS 'Total Household child count';
COMMENT ON COLUMN pantry.guests_archive.count_adults IS 'Total Household adult count';
COMMENT ON COLUMN pantry.guests_archive.worship_place IS 'Primary place of worship';
COMMENT ON COLUMN pantry.guests_archive.is_member IS 'Is a Member of the Laurel Church of Christ';
COMMENT ON COLUMN pantry.guests_archive.is_baptized IS 'Is Baptized into Jesus Christ';
COMMENT ON COLUMN pantry.guests_archive.is_espanol IS 'Is Spanish speaking, Hispanic';
COMMENT ON COLUMN pantry.guests_archive.is_unemployed IS 'Is looking for employmentt';
COMMENT ON COLUMN pantry.guests_archive.is_homeless IS 'Is the Guest Homeless? Does not have a Home? Stable roof over the head';
COMMENT ON COLUMN pantry.guests_archive.is_family IS 'Has at least 2 people that lives with guest';
COMMENT ON COLUMN pantry.guests_archive.is_contact_ok IS 'Is okay for Contact? [0 = No, -1 = Yes]';
COMMENT ON COLUMN pantry.guests_archive.allergies IS 'Note of known allergies pertaining to this guest';
COMMENT ON COLUMN pantry.guests_archive.notes IS 'Additional remarks';
COMMENT ON COLUMN pantry.guests_archive.last_date_updated IS 'Last date that this record of this Guest was updated';

COMMENT ON CONSTRAINT guests_archive_pkey ON pantry.guests_archive IS 'Primary Key for the Guests table: id';
