CREATE OR REPLACE FUNCTION pantry.enforce_guests_status_waiting()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
 BEGIN 
   IF TG_OP='INSERT' THEN 
     IF NEW.status='Active' THEN 
        RAISE EXCEPTION 'We are not excepting any new visitors on Active. Please set status to "Waiting"'; 
        RETURN NEW;
     END IF; 
   ELSIF TG_OP='UPDATE' THEN
     IF OLD.status IN('Inactive','Waiting','Archive',NULL) AND NEW.status='Active' THEN
       RAISE EXCEPTION 'We are not excepting any new visitors on Active. Please set status to "Waiting" or '
       'set status back to "%"', OLD.status; 
       -- RETURN NEW;
     END IF;
   END IF;
   RETURN NEW; 
END; $function$
;
