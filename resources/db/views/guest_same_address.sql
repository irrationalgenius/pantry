CREATE OR REPLACE VIEW pantry.guests_same_address
AS SELECT guests.status, guests.first_name, guests.last_name, guests.st_address, 
    guests.city, guests.state, guests.zip, guests.date_enrolled
   FROM pantry.guests
  WHERE (guests.st_address::text IN ( SELECT g.st_address
           FROM pantry.guests g
          GROUP BY g.st_address
         HAVING count(g.st_address) > 1)) AND guests.status::text <> 'Archive'::text
  ORDER BY guests.st_address;
