-- TODO: answer here
SELECT
    reports.id AS id,
    students.fullname AS fullname,
    students.class AS class,
    students.status AS status,
    reports.study AS study,
    reports.score AS score
FROM reports
JOIN students ON students.id = reports.student_id
WHERE reports.score < 70
ORDER BY reports.score ASC;