CREATE OR REPLACE FUNCTION pantry.sys_guests_to_archive()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  ALTER TABLE pantry.guests DISABLE TRIGGER enforce_guests_gender_trgr;
  
  UPDATE pantry.guests g SET status='Archive'
  WHERE g.status IN('Inactive','Waiting','Active') 
    AND g.id IN(SELECT g.id FROM pantry.guests g
		LEFT JOIN pantry.guest_visits v ON g.id=v.guest_id
		GROUP BY g.id
		HAVING max(date_visit) < (current_date-365));

  ALTER TABLE pantry.guests ENABLE TRIGGER enforce_guests_gender_trgr;
END; $function$
;

