CREATE OR REPLACE VIEW pantry.guest_visits_espanol
AS SELECT g.status, max(gv.date_visit) AS "last visit", g.first_name, g.last_name, 
    g.city, g.adult_no, g.child_no, g.tel_no, g.email, g.is_espanol, gv.notes
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE g.is_espanol::text = 'Yes'::text
  GROUP BY g.status, g.first_name, g.last_name, g.city, g.adult_no, g.child_no, g.tel_no, g.email, g.is_espanol, gv.notes;
