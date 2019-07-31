CREATE OR REPLACE VIEW pantry.guest_visits_top_10
AS SELECT g.status, g.first_name, g.last_name, count(gv.date_visit) AS visits
   FROM pantry.guest_visits gv
   JOIN pantry.guests g ON g.id = gv.guest_id
  GROUP BY g.status, g.first_name, g.last_name
  ORDER BY count(gv.date_visit) DESC
 LIMIT 10;
