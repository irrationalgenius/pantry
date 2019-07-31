CREATE OR REPLACE VIEW pantry.guests_enrolled_to_last_visit
AS SELECT (g.first_name::text || ' '::text) || g.last_name::text AS fullname, 
    g.date_enrolled, gv.date_visit, gv.date_visit_next
   FROM pantry.guest_visits gv
   JOIN pantry.guests g ON g.id = gv.guest_id
  ORDER BY gv.date_visit DESC, g.date_enrolled DESC;
