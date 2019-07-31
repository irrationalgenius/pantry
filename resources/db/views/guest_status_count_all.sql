CREATE OR REPLACE VIEW pantry.guests_status_count_all
AS SELECT guests.status, count(guests.status) AS count
   FROM pantry.guests
  GROUP BY guests.status;
