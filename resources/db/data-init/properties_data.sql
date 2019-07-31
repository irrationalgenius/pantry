
/*
  Property Categories:
    1. Configuration = Initial Load values for Application Execution
    2. Information = Essential information pertaining to the pantry
    3. Operation = Critial values to Managing the application (backend)
*/

INSERT INTO pantry.properties(prop_category, prop_name, prop_value)
  VALUES('Configuration', 'APP_NAME', 'Laurel Food Pantry Application');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value)
  VALUES('Configuration', 'APP_VERSION', '0.1');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value)
  VALUES('Configuration', 'APP_PURPOSE', 'Manage Guests and Visits within the Pantry');
