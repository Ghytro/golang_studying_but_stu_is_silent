-- IF EXISTS(SELECT * FROM information_schema.tables WHERE table_name = 'phones') THEN
--     TRUNCATE TABLE phones;
-- ELSE
--     CREATE TABLE phones (
--         Id SERIAL PRIMARY KEY,
--         PhoneNumber VARCHAR(30) NOT NULL UNIQUE,
--         SubscriberFirstName VARCHAR(30) NOT NULL,
--         SubscriberLastName VARCHAR(30) NOT NULL
--     );
-- END IF;

TRUNCATE TABLE phones;