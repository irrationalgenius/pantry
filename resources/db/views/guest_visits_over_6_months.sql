CREATE OR REPLACE VIEW pantry.guest_visits_over_6_months
AS SELECT g.status, max(gv.date_visit) AS "last visit", g.first_name, g.last_name, 
    g.city, g.adult_no, g.child_no, g.tel_no, g.email, gv.notes
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE gv.date_visit >= ('now'::text::date - 365) AND gv.date_visit <= ('now'::text::date - 183)
  GROUP BY g.status, g.first_name, g.last_name, g.city, g.adult_no, g.child_no, g.tel_no, g.email, gv.notes;
