
DROP VIEW pantry.guests_dups_cleanup;

CREATE VIEW pantry.guests_dups_cleanup AS
SELECT g.id, status, first_name, last_name, st_address, zip
       , max(gv.date_visit) AS max
FROM pantry.guests g
LEFT JOIN pantry.guest_visits gv ON g.id = gv.guest_id
WHERE (first_name, last_name, zip) IN(SELECT first_name, last_name, zip
                                      FROM pantry.guests
                                      GROUP BY first_name, last_name, zip
                                      HAVING COUNT(*) > 1)
GROUP BY g.id, status, first_name, last_name, st_address, zip
ORDER BY first_name, last_name, zip;
