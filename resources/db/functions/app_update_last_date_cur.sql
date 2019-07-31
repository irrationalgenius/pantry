CREATE OR REPLACE FUNCTION pantry.update_last_date_current()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN 
  IF TG_TABLE_NAME='metrics_food' THEN
    NEW.last_date_updated:=('now'::text)::date;
  END IF; 
  IF TG_TABLE_NAME='guests' THEN
    NEW.last_date_updated:=('now'::text)::date;
  END IF;
  IF TG_TABLE_NAME='guest_visits' THEN
    NEW.last_date_updated:=('now'::text)::date;    
    IF (NEW.date_visit<>OLD.date_visit) THEN
      NEW.date_visit_next:=(NEW.date_visit + 42);
    END IF;    
  END IF;
  RETURN NEW;
END; $function$
;

