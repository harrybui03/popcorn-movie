INSERT INTO theaters
VALUES  ('018e9fbf-0809-76a7-b848-ab1651254acf' ,'02 Phạm Ngọc Thạch, Trung Tự, Q. Đống Đa, Tp. Hà Nội', 'Popcorn Movive Phạm Ngọc Thạch', '0123456121'),
        ('018e9fbf-0809-7441-8bd4-d9a2620dd4e1', 'Tầng 8, TTTM Discovery – 302 Cầu Giấy, P.Dịch Vọng, Quận Cầu Giấy, Hà Nội', 'Popcorn Movive Discovery', '0124457121'),
        ( '018e9fbf-0809-7bcc-a913-9abfa0bc2402' ,'Tầng 4, TTTM The Garden, khu đô thị The Manor, đường Mễ Trì, phường Mỹ Đình 1, quận Nam Từ Liêm, Hà Nội', 'Popcorn Movive The Garden', '0223456121');

INSERT INTO rooms (id, room_number,theater_id)
VALUES  ('018e9fc2-0cf4-78cf-93b2-4f75ac4bd9b5',101, '018e9fbf-0809-76a7-b848-ab1651254acf'),
        ('018e9fc2-0cf4-73cb-974d-68629a67ccad',102, '018e9fbf-0809-76a7-b848-ab1651254acf'),
        ('018e9fc2-0cf4-7c2a-af88-3aa8f7ebec11',101, '018e9fbf-0809-7441-8bd4-d9a2620dd4e1'),
        ('018e9fc2-0cf4-7649-bcfe-52109b4e29c9',102, '018e9fbf-0809-7441-8bd4-d9a2620dd4e1'),
        ('018e9fc2-0cf4-74ef-b076-7fa6d2ec5cfe',101, '018e9fbf-0809-7bcc-a913-9abfa0bc2402'),
        ('018e9fc2-0cf4-7587-8111-129f238870e4',102, '018e9fbf-0809-7bcc-a913-9abfa0bc2402'),
        ('018e9fc2-0cf4-7b80-9184-90da25885dd2', 103, '018e9fbf-0809-76a7-b848-ab1651254acf');