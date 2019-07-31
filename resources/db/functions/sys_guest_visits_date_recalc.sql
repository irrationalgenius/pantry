CREATE OR REPLACE FUNCTION pantry.sys_guest_visits_date_recalc()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  UPDATE pantry.guest_visits SET date_visit_next=(date_visit+42)
  WHERE date_visit_next>current_date 
    AND date_visit_next<>(date_visit+42);
end;
$function$
;
