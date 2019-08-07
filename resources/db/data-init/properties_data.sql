
/*
  Property Categories:
    1. Configuration = Initial Load values for Application Execution
    2. Information = Essential information pertaining to the pantry
    3. Operation = Critial values to Managing the application (backend)
*/

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Configuration', 'APP_NAME', 'Food Pantry Application', 'Application Name');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Configuration', 'APP_VERSION', '0.1', 'Application Version');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Configuration', 'APP_PURPOSE', 'Manage Guests and Visits within the Pantry', 'Application Purpose');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Operation', 'APP_DEVELOPER', 'Calvin Hunt', 'Developer Name');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Operation', 'APP_DEVELOPER_EMAIL', 'hunt.calvin@outlook.com', 'Developer Email');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Operation', 'APP_OWNER', 'Laurel Church of Christ MD', 'Application Owner');

INSERT INTO pantry.properties(prop_category, prop_name, prop_value, prop_description)
  VALUES('Configuration', 'APP_VISIT_INTERVAL', '42', 'Current Visit Interval. Default 42 days (about 6 weeks)');
