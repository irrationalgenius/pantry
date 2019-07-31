CREATE OR REPLACE FUNCTION pantry.sys_guests_to_inactive()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  ALTER TABLE pantry.guests DISABLE TRIGGER enforce_guests_gender_trgr;
  
  UPDATE pantry.guests g SET status='Inactive'
  WHERE g.status='Active' 
    AND g.id IN(SELECT guest_id 
		FROM pantry.guest_visits v 
		GROUP BY guest_id
		HAVING MAX(date_visit) BETWEEN (current_date-365) AND (current_date-183));
		
  ALTER TABLE pantry.guests ENABLE TRIGGER enforce_guests_gender_trgr;
END;
$function$
;

