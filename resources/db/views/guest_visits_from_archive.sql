CREATE OR REPLACE VIEW pantry.guest_visits_from_archive
AS SELECT g.status, (g.first_name::text || ' '::text) || g.last_name::text, 
    max(gv.date_visit) AS max
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON gv.guest_id = g.id
  WHERE g.status::text = 'Archive'::text
  GROUP BY g.status, g.first_name, g.last_name
  ORDER BY max(gv.date_visit) DESC;
