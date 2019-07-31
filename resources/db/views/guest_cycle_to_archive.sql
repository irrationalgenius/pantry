CREATE OR REPLACE VIEW pantry.guests_cycle_to_archive
AS SELECT g.status, g.first_name, g.last_name, max(v.date_visit) AS max
   FROM pantry.guests g
   JOIN pantry.guest_visits v ON g.id = v.guest_id
  WHERE g.status::text = ANY (ARRAY['Active'::character varying, 'Inactive'::character varying, 'Waiting'::character varying]::text[])
  GROUP BY g.status, g.first_name, g.last_name
 HAVING max(v.date_visit) < ('now'::text::date - 365);
