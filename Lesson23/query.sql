-- UPDATE jadvaldagi mavjud yozuvlarni o'zgartirish uchun ishlatiladi. Bu malumotlarni tuzatish,
-- ozgarishlarni amalga oshirish va yangi malumotlarga javoban yozuvlarni yangilash uchun foydalidir.

UPDATE students SET stipend = stipend * 1.10 WHERE course = 2;

-- bu yerda 2chi kurs studentlar stippendiyasi 1.10 ga oshadi



-- DELETE operatori jadvaldagi yozuvlarni olib tashlaydi.
-- Bu eskirgan, ahamiyatsiz yoki notogri malumotlarni olib tashlash orqali malumotlar saqlash uchun zarurdir.

DELETE FROM students WHERE course = 1;

-- bu yerda 1 kurs talabalar ochirilmoqda



-- GROUP BY iborasi belgilangan ustunlardagi qiymatlari bir xil bolgan satrlarni jamlangan malumotlarga guruhlaydi
-- va bu guruhlarda COUNT, AVG, SUM, MAX, MIN kabi jamlangan funksiyalarni bajarishga imkon beradi.

SELECT course, AVG(stipend) AS average_stipend FROM students GROUP BY course;

-- bu yerda esa kurslar boyicha ortach stipendiyani chiqarib beradi masalan course 1 deb 1 kurslar ortacha stipendiyasini 2 deb 2ni vaho kazo



-- ORDER BY iborasi natijani bir yoki bir nechta ustunlar boyicha osish (standart) yoki kamayish tartibida tartiblaydi.
-- Bu malumotlarni mazmunli tashkil qilish, tahlil qilish va tushunishni osonlashtirish uchun juda muhimdir.

SELECT * FROM students ORDER BY stipend DESC;

-- bu joyda studentlar stipendiya boyich kamayish tartibida chiqadi



-- JOIN iborasi ikki yoki undan ortiq jadvallar qatorlarini ular orasidagi tegishli ustun asosida birlashtiradi.
-- Bu relyatsion malumotlar bazalari uchun juda muhim bolib, turli jadvallarda saqlanadigan tegishli malumotlarni birlashtirishga imkon beradi.

INNER JOIN courses ON students.course = courses.course_id;

SELECT students.name, students.stipend, courses.course_name FROM students INNER JOIN courses ON students.course = courses.course_id;

-- bu yerda inner join ishlatilyapti va students table bilan courses tableni boglayapti

SELECT students.name, students.stipend, courses.course_name FROM students LEFT JOIN courses ON students.course = courses.course_id;
-- bu yerda ham huddi shunaqa faqat left join bilan

SELECT students.name, students.stipend, courses.course_name FROM students RIGHT JOIN courses ON students.course = courses.course_id;

-- Right join bilan

SELECT students.name, students.stipend, courses.course_name FROM students FULL JOIN courses ON students.course = courses.course_id;

-- Full join bilan
