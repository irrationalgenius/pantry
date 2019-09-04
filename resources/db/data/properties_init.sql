
/*
  Property Categories:
    1. Configuration = Initial Load values for Application Execution
    2. Information = Essential information pertaining to the pantry
    3. Operation = Critial values to Managing the application (backend)
*/

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('configuration', 'A', 'app_name', 'Food Pantry Application', 'Application Name');

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('configuration', 'A', 'app_version', '0.1', 'Application Version');

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('configuration', 'A', 'app_purpose', 'Manage Guests and Visits within the Pantry', 'Application Purpose');

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('operation', 'A', 'app_developer', 'Calvin Hunt', 'Developer Name');

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('operation', 'A', 'app_developer_email', 'hunt.calvin@outlook.com', 'Developer Email');

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('operation', 'A', 'app_owner', 'Laurel Church of Christ MD', 'Application Owner');

INSERT INTO properties(prop_category, status, prop_name, prop_value, prop_description)
  VALUES('configuration', 'A', 'app_visit_interval', '42', 'Current Visit Interval. Default 42 days (about 6 weeks)');
