CREATE OR REPLACE FUNCTION pantry.app_guest_visits_add(p_gid integer, p_visit date, p_note character varying)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  INSERT INTO pantry.guest_visits(
     guest_id,date_visit,date_visit_next,notes)
  VALUES(
     p_gid,p_visit,(p_visit+42),p_note); 
END;
$function$
;
