CREATE OR REPLACE VIEW pantry.metrics_food_totals
AS SELECT metrics_food.food_name, metrics_food.store_code, 
    metrics_food.price * metrics_food.food_count AS total_cost
   FROM pantry.metrics_food;
