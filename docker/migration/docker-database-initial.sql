CREATE TABLE teachers(
    id serial primary key,
    first_name varchar,
    last_name varchar,
    age int,
    email varchar,
    class int
);

INSERT INTO teachers(first_name, last_name, age, email, class) VALUES
  ('Albert', 'Vignesh', 28, 'teacher1@go-api-class.com', 10),
  ('Ana', 'Kumar', 22, 'teacher2@go-api-class.com', 11),
  ('Eric', '', 25, 'teacher3@go-api-class.com', 12),
  ('Paul', 'Mark', 24, 'teacher4@go-api-class.com', 10)
;