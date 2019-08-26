
DROP TABLE visits_archive;

CREATE TABLE visits_archive (
	id serial NOT NULL,
	guest_id int4 NOT NULL, -- Foreign key column to the "id" column on the guest table
	date_visit date NOT NULL, -- Day of pantry visit
	date_visit_next date NOT NULL, -- Calculated Day of next visit day
	notes text NULL, -- Field to store updates when the Guest vistits
	last_date_updated date NULL, -- Captures the date when this record is updated, after it has been inserted

  -- Archival information
  archive_last_date_updated date NOT NULL DEFAULT 'now'::text::date,
  archive_method varchar(1) NOT NULL, -- How was this record archived? [D=Deleted, X=Inserted by Procedure]

	CONSTRAINT visits_archive_pk PRIMARY KEY (id)
);
COMMENT ON TABLE visits IS 'Captures each instance of Guest visits in the Pantry';

-- Column comments

COMMENT ON COLUMN visits_archive.guest_id IS 'Foreign key column to the "id" column on the guest table';
COMMENT ON COLUMN visits_archive.date_visit IS 'Day of pantry visit';
COMMENT ON COLUMN visits_archive.date_visit_next IS 'Calculated Day of next visit day';
COMMENT ON COLUMN visits_archive.notes IS 'Field to store updates when the Guest vistits';
COMMENT ON COLUMN visits_archive.last_date_updated IS 'Captures the date when this record is updated, after it has been inserted';

-- Table Triggers

-- DROP TRIGGER last_date_upd_guest_trgr ON guest_visits;

-- create
--     trigger last_date_upd_guest_trgr before update
--         on
--         guest_visits for each row execute procedure update_last_date_current();
