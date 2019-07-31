CREATE OR REPLACE VIEW pantry.guest_visits_next_7_days
AS SELECT g.status, max(gv.date_visit) AS "last visit", 
    max(gv.date_visit_next) AS "next visit", g.id, g.first_name, g.last_name, 
    g.city, g.adult_no, g.child_no, g.tel_no, g.email, gv.notes
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE gv.date_visit_next >= 'now'::text::date AND gv.date_visit_next <= ('now'::text::date + 7)
  GROUP BY g.status, g.id, g.first_name, g.last_name, g.city, g.adult_no, g.child_no, g.tel_no, g.email, gv.notes
  ORDER BY max(gv.date_visit_next) DESC;
