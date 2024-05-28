-- 1
SELECT
    c.name AS course_name,
    s.name AS student_name,
    g.grade_value
FROM
    course c
JOIN
    student_course sc ON c.id = sc.course_id
JOIN
    grade g ON sc.student_id = g.student_id AND sc.course_id = g.course_id
JOIN
    student s ON s.id = sc.student_id
WHERE
    (g.course_id, g.grade_value) IN (
        SELECT
            g.course_id,
            MAX(g.grade_value)
        FROM
            grade g
        GROUP BY
            g.course_id
    )
ORDER BY
    c.name, g.grade_value DESC;

--2

SELECT
    c.name AS course_name,
    AVG(g.grade_value) AS average_grade
FROM
    course c
JOIN
    student_course sc ON c.id = sc.course_id
JOIN
    grade g ON sc.student_id = g.student_id AND sc.course_id = g.course_id
GROUP BY
    c.name;

--3




-- 4
SELECT
    course_name,
    average_grade
FROM (
    SELECT
        c.name AS course_name,
        AVG(g.grade_value) AS average_grade
    FROM
        course c
    JOIN
        student_course sc ON c.id = sc.course_id
    JOIN
        grade g ON sc.student_id = g.student_id AND sc.course_id = g.course_id
    GROUP BY
        c.name
    ORDER BY
        AVG(g.grade_value) DESC
    LIMIT 1
) AS best_group;


--5
WITH BestGrades AS (
    SELECT
        course_id,
        MAX(grade_value) AS max_grade
    FROM
        grade
    GROUP BY
        course_id
)
SELECT
    c.name AS course_name,
    s.name AS student_name,
    g.grade_value
FROM
    grade g
JOIN
    student s ON g.student_id = s.id
JOIN
    course c ON g.course_id = c.id
JOIN
    BestGrades bg ON g.course_id = bg.course_id AND g.grade_value = bg.max_grade
ORDER BY
    c.name, g.grade_value DESC;

