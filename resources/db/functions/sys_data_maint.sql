CREATE OR REPLACE FUNCTION pantry.sys_data_maint()
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
  PERFORM pantry.sys_guests_to_archive();
  PERFORM pantry.sys_guests_to_inactive();
  PERFORM pantry.sys_guests_to_active();
  PERFORM pantry.sys_guests_set_family();
  PERFORM pantry.sys_guests_info_initcap();
  PERFORM pantry.sys_guest_visits_date_recalc();
END; $function$
;
