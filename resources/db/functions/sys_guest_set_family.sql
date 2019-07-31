CREATE OR REPLACE FUNCTION pantry.sys_guests_set_family()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  ALTER TABLE pantry.guests DISABLE TRIGGER enforce_guests_gender_trgr;
  
  UPDATE pantry.guests SET is_family='-1' WHERE child_no>0;
  UPDATE pantry.guests SET is_Family='0' WHERE (child_no<1 OR child_no IS NULL);

  ALTER TABLE pantry.guests ENABLE TRIGGER enforce_guests_gender_trgr;
end;
$function$
;
