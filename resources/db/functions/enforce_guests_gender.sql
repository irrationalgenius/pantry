CREATE OR REPLACE FUNCTION pantry.enforce_guests_gender()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
 BEGIN 
    IF NEW.gender IS NULL THEN 
        RAISE EXCEPTION 'Please select the Guests Gender'; 
    END IF; 
   RETURN NEW; 
END; $function$
;
