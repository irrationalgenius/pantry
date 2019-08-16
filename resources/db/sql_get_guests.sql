select id, date_enrolled, status, first_name, last_name, email, count_children, count_adults, (count_children + count_adults) count_total, is_baptized, last_date_updated
from pantry.guests;
