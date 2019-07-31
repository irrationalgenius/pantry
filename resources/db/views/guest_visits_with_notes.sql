CREATE OR REPLACE VIEW pantry.guest_visits_notes_correspondance
AS SELECT (g.first_name::text || ' '::text) || g.last_name::text AS fullname, 
    g.tel_no AS contact_number, 
        CASE
            WHEN g.email::text ~~ '%*%'::text THEN NULL::character varying
            ELSE g.email
        END AS email, 
    gv.date_visit, gv.date_visit_next, g.notes AS guest_note, 
    gv.notes AS visit_note, 
    g.last_date_updated AS last_guest_record_updated_date, 
    gv.last_date_updated AS last_visit_record_update_date
   FROM pantry.guest_visits gv
   JOIN pantry.guests g ON g.id = gv.guest_id
  WHERE gv.notes IS NOT NULL
  ORDER BY gv.date_visit;
