
DROP TABLE pantry.guests_archive;

CREATE TABLE pantry.guests_archive (
  id serial NOT NULL, -- Unique Identifier
	date_enrolled date NULL DEFAULT 'now'::text::date, -- The date of the first pantry visit
	status varchar(10) NULL, -- Current status [A=Active; I=Inactive; W=Waiting; X=Archive]
	first_name varchar(35) NULL, -- First name
	last_name varchar(35) NULL, -- Last name
	gender varchar(6) NULL, -- Gender [M=Male; F=Female]
	unit_num varchar(20) NULL, -- Number of the adress unit: street, apt, etc.
	st_address varchar(50) NULL, -- Current Street Address
	state varchar(2) NULL, -- State code
	city varchar(35) NULL, -- Current City
	zip varchar(11) NULL, -- 5 digit Zip code
	tel_num varchar(12) NULL, -- Current Primary contact number
	email varchar(50) NULL, -- Current Primary email address
	count_children int4 NULL, -- Total Household child count
	count_adults int4 NULL, -- Total Household adult count
	worship_place varchar(80) NULL, -- Primary place of worship
	is_member varchar(3) NULL DEFAULT 0, -- Is a Member of the Laurel Church of Christ
	is_baptized varchar(3) NULL DEFAULT 0, -- Is Baptized into Jesus Christ
	is_espanol varchar(3) NULL DEFAULT 0, -- Is Spanish speaking, Hispanic
	is_unemployed varchar(3) NULL DEFAULT 0, -- Is looking for employment
	is_homeless varchar(3) NULL DEFAULT 0, -- Is the Guest Homeless? Does not have a Home? Stable roof over the head
	is_family varchar(3) NULL DEFAULT 0, -- Has at least 2 people that lives with guest
	is_contact_ok varchar(3) NULL DEFAULT 0, -- Is okay for Contact? [0 = No, -1 = Yes]
	allergies text NULL, -- Note of known allergies pertaining to this guest
	notes text NULL, -- Additional remarks
	last_date_updated date NULL, -- Last date that this record of this Guest was updated

  -- Archiving Data
  archive_count int,
  archive_date_last date,
  archive_method varchar(1) -- How was this record archived? [D=Deleted, X=Inserted by Procedure]
);
