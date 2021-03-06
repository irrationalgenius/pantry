
DROP TABLE guests;

CREATE TABLE guests (
	id serial NOT NULL, -- Unique Identifier
	date_enrolled date NULL DEFAULT 'now'::text::date, -- The date of the first pantry visit
	status varchar(1) NULL, -- Current status [A=Active; I=Inactive; W=Waiting; X=Archive]
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
	-- origin text NULL, -- How did you hear about Us?
	is_member varchar(1) NULL, -- Is a Member of the Laurel Church of Christ
	is_baptized varchar(1) NULL, -- Is Baptized into Jesus Christ
	is_espanol varchar(1) NULL, -- Is Spanish speaking, Hispanic
	is_unemployed varchar(1) NULL, -- Is looking for employment
	is_homeless varchar(1) NULL, -- Is the Guest Homeless? Does not have a Home? Stable roof over the head
	is_family varchar(1) NULL, -- Has at least 2 people that lives with guest
	is_contact_ok varchar(1) NULL, -- Is okay for Contact? [0 = No, -1 = Yes]
	allergies text NULL, -- Note of known allergies pertaining to this guest
	notes text NULL, -- Additional remarks
	last_date_updated date NULL DEFAULT 'now'::text::date, -- Last date that this record of this Guest was updated

	CONSTRAINT guests_pkey PRIMARY KEY (id)
);
COMMENT ON TABLE guests IS 'The Active Collection of Laurel Pantry Guests';

-- Column comments

COMMENT ON COLUMN guests.id IS 'Unique Identifier';
COMMENT ON COLUMN guests.date_enrolled IS 'The date of the first pantry visit';
COMMENT ON COLUMN guests.status IS 'Current status [A=Active; I=Inactive; W=Waiting; X=Archive]';
COMMENT ON COLUMN guests.first_name IS 'First name';
COMMENT ON COLUMN guests.last_name IS 'Last name';
COMMENT ON COLUMN guests.gender IS 'Gender [M=Male; F=Female]';
COMMENT ON COLUMN guests.st_address IS 'Current Street Address';
COMMENT ON COLUMN guests.unit_num IS 'Number of the adress unit: street, apt, etc.';
COMMENT ON COLUMN guests.state IS 'Current State code';
COMMENT ON COLUMN guests.city IS 'Current City';
COMMENT ON COLUMN guests.zip IS '5 digit Zip code';
COMMENT ON COLUMN guests.tel_num IS 'Current Primary contact number';
COMMENT ON COLUMN guests.email IS 'Current Primary email address';
COMMENT ON COLUMN guests.count_children IS 'Total Household child count';
COMMENT ON COLUMN guests.count_adults IS 'Total Household adult count';
COMMENT ON COLUMN guests.worship_place IS 'Primary place of worship';
COMMENT ON COLUMN guests.is_member IS 'Is a Member of the Laurel Church of Christ';
COMMENT ON COLUMN guests.is_baptized IS 'Is Baptized into Jesus Christ';
COMMENT ON COLUMN guests.is_espanol IS 'Is Spanish speaking, Hispanic';
COMMENT ON COLUMN guests.is_unemployed IS 'Is looking for employmentt';
COMMENT ON COLUMN guests.is_homeless IS 'Is the Guest Homeless? Does not have a Home? Stable roof over the head';
COMMENT ON COLUMN guests.is_family IS 'Has at least 2 people that lives with guest';
COMMENT ON COLUMN guests.is_contact_ok IS 'Is okay for Contact? [0 = No, -1 = Yes]';
COMMENT ON COLUMN guests.allergies IS 'Note of known allergies pertaining to this guest';
COMMENT ON COLUMN guests.notes IS 'Additional remarks';
COMMENT ON COLUMN guests.last_date_updated IS 'Last date that this record of this Guest was updated';

-- Constraint comments

COMMENT ON CONSTRAINT guests_pkey ON guests IS 'Primary Key for the Guests table: id';

-- Table Triggers

-- DROP TRIGGER enforce_guests_status_wait_trgr ON guests;

-- create
--     trigger enforce_guests_status_wait_trgr before insert
--         or update
--             on
--             guests for each row execute procedure enforce_guests_status_waiting();
-- DROP TRIGGER last_date_upd_guest_trgr ON guests;

-- create
--     trigger last_date_upd_guest_trgr before update
--         on
--         guests for each row execute procedure update_last_date_current();
-- DROP TRIGGER enforce_guests_gender_trgr ON guests;

-- create
--     trigger enforce_guests_gender_trgr before insert
--         or update
--             on
--             guests for each row execute procedure enforce_guests_gender();
