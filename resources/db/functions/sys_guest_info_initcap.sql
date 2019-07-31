CREATE OR REPLACE FUNCTION pantry.sys_guests_info_initcap()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  ALTER TABLE pantry.guests DISABLE TRIGGER enforce_guests_gender_trgr;
  
  UPDATE pantry.guests SET first_name=INITCAP(first_name);
  UPDATE pantry.guests SET last_name=INITCAP(last_name);
  UPDATE pantry.guests SET st_address=INITCAP(st_address);
  UPDATE pantry.guests SET city=INITCAP(city);

  ALTER TABLE pantry.guests ENABLE TRIGGER enforce_guests_gender_trgr;
END; $function$
;

