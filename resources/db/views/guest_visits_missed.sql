CREATE OR REPLACE VIEW pantry.guest_visits_missed
AS SELECT gv.guest_id, max(gv.date_visit) AS last_date_visited, 
    max(gv.date_visit_next) AS missed_date_visit, g.status, g.first_name, 
    g.last_name, g.tel_no, g.email, g.adult_no, g.child_no, g.notes
   FROM pantry.guest_visits gv
   JOIN pantry.guests g ON g.id = gv.guest_id
  WHERE gv.date_visit_next > ('now'::text::date - 183)
  GROUP BY gv.guest_id, g.status, g.first_name, g.last_name, g.tel_no, g.email, g.adult_no, g.child_no, g.notes
 HAVING max(gv.date_visit_next) < 'now'::text::date AND max(gv.date_visit) < 'now'::text::date
  ORDER BY max(gv.date_visit_next) DESC;
