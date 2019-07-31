CREATE OR REPLACE VIEW pantry.guests_dups_cleanup
AS SELECT g.id, g.status, (g.first_name::text || ' '::text) || g.last_name::text, 
    g.st_address, max(gv.date_visit) AS max
   FROM pantry.guests g
   LEFT JOIN pantry.guest_visits gv ON g.id = gv.guest_id
  WHERE (g.first_name::text IN ( SELECT guests.first_name
      FROM pantry.guests
     GROUP BY guests.first_name
    HAVING count(guests.first_name) > 1)) AND (g.last_name::text IN ( SELECT guests.last_name
      FROM pantry.guests
     GROUP BY guests.last_name
    HAVING count(guests.last_name) > 1))
  GROUP BY g.id, g.status, g.first_name, g.last_name, g.st_address
  ORDER BY g.first_name, g.last_name;
