CREATE OR REPLACE VIEW pantry.guest_visits_cities
AS SELECT g.city, count(gv.date_visit) AS count
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  GROUP BY g.city
  ORDER BY g.city DESC;
