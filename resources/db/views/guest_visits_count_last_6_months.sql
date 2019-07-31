CREATE OR REPLACE VIEW pantry.guest_visits_count_last_6_months
AS SELECT guest_visits.date_visit, count(guest_visits.date_visit) AS date_counts
   FROM pantry.guest_visits
  WHERE guest_visits.date_visit > ('now'::text::date - 183)
  GROUP BY guest_visits.date_visit
  ORDER BY guest_visits.date_visit DESC;
