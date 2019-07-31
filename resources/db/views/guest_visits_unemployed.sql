CREATE OR REPLACE VIEW pantry.guest_visits_job_searching
AS SELECT g.status, max(gv.date_visit) AS "last visit", g.first_name, g.last_name, 
    g.city, g.adult_no, g.child_no, g.tel_no, g.email, g.is_job_searching, 
    gv.notes
   FROM pantry.guests g
   JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE g.is_job_searching::text = 'Yes'::text
  GROUP BY g.status, g.first_name, g.last_name, g.city, g.adult_no, g.child_no, g.tel_no, g.email, g.is_job_searching, gv.notes;
