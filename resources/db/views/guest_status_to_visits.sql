CREATE OR REPLACE VIEW pantry.guests_status_to_visits
AS SELECT g.status, (g.first_name::text || ' '::text) || g.last_name::text, 
    gv.date_visit
   FROM pantry.guest_visits gv
   JOIN pantry.guests g ON g.id = gv.guest_id
  GROUP BY g.status, g.first_name, g.last_name, gv.date_visit
 HAVING max(gv.date_visit) > ('now'::text::date - 183)
  ORDER BY g.status DESC;
