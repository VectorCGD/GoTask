CREATE TABLE students(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(5) NOT NULL,
    age INT CHECK(age > 0 && age < 120),
    grade CHAR(3)
);

INSERT INTO students(name,age,grade) VALUES('张三',20,'三年级');

SELECT * FROM students WHERE age > 18;

UPDATE students SET grade = '四年级' WHERE name = '张三'; 

DELETE FROM students WHERE age < 15;