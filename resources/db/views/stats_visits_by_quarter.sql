CREATE OR REPLACE VIEW pantry.stats_visits_by_quarter
AS SELECT date_trunc('quarter'::text, guest_visits.date_visit::timestamp with time zone) AS visit_quarter, 
    count(DISTINCT guest_visits.id) AS visits
   FROM pantry.guest_visits
  GROUP BY date_trunc('quarter'::text, guest_visits.date_visit::timestamp with time zone)
  ORDER BY date_trunc('quarter'::text, guest_visits.date_visit::timestamp with time zone);
