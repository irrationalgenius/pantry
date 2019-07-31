CREATE OR REPLACE VIEW pantry.stats_visits_by_mnths
AS SELECT date_trunc('month'::text, guest_visits.date_visit::timestamp with time zone) AS date_trunc, 
    count(*) AS count
   FROM pantry.guest_visits
  WHERE guest_visits.date_visit IS NOT NULL
  GROUP BY date_trunc('month'::text, guest_visits.date_visit::timestamp with time zone)
  ORDER BY date_trunc('month'::text, guest_visits.date_visit::timestamp with time zone);
