CREATE OR REPLACE VIEW pantry.guest_visits_last_6_months_nonactive
AS SELECT g.status, g.first_name, g.last_name, max(gv.date_visit) AS last_visit
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE g.status::text <> 'Active'::text
  GROUP BY g.status, g.first_name, g.last_name, gv.date_visit
 HAVING max(gv.date_visit) >= ('now'::text::date - 183) AND max(gv.date_visit) <= 'now'::text::date
  ORDER BY max(gv.date_visit) DESC;
