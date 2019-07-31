CREATE OR REPLACE VIEW pantry.guest_visits_count_last_30_days
AS SELECT guest_visits.date_visit, count(guest_visits.date_visit) AS date_counts
   FROM pantry.guest_visits
  WHERE guest_visits.date_visit > ('now'::text::date - 30)
  GROUP BY guest_visits.date_visit
  ORDER BY guest_visits.date_visit DESC;
