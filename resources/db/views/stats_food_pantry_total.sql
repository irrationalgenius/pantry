CREATE OR REPLACE VIEW pantry.metrics_food_pantry_total
AS SELECT sum(metrics_food_totals.total_cost) AS pantry_overall_cost
   FROM pantry.metrics_food_totals;
