CREATE OR REPLACE FUNCTION pantry.sys_guests_contactinfo_refresh()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  UPDATE pantry.guests g SET tel_no=NULL, email=NULL
  WHERE g.status='Active';
END; $function$
;
