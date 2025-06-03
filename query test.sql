CREATE TABLE query_test(
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	parent_id INT
);

INSERT INTO query_test (name, parent_id) VALUES
('Zaki', 2),
('Ilham', null),
('Irwan', 2),
('Arka', 3);

SELECT a.id AS id, a.name AS name, b.name AS parent_name FROM query_test AS a LEFT JOIN query_test AS b ON a.parent_id = b.id ORDER BY id;