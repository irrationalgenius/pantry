CREATE OR REPLACE VIEW pantry.guest_visits_last_7_days
AS SELECT g.status, max(gv.date_visit) AS "last visit", g.id, g.first_name, 
    g.last_name, g.city, g.adult_no, g.child_no, g.tel_no, g.email, gv.notes
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE gv.date_visit >= ('now'::text::date - 7) AND gv.date_visit <= 'now'::text::date
  GROUP BY g.status, g.id, g.first_name, g.last_name, g.city, g.adult_no, g.child_no, g.tel_no, g.email, gv.notes
  ORDER BY max(gv.date_visit) DESC;
