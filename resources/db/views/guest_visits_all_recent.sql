CREATE OR REPLACE VIEW pantry.guest_visits_all_recent
AS SELECT gv.date_visit, gv.date_visit_next, g.status, 
    (g.first_name::text || ' '::text) || g.last_name::text, g.city, g.tel_no, 
    g.email, g.date_enrolled, g.notes
   FROM pantry.guest_visits gv
   JOIN pantry.guests g ON g.id = gv.guest_id
  ORDER BY gv.id DESC;
