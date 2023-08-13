-- +goose Up
-- +goose StatementBegin
INSERT INTO users (first_name, last_name, email, phone, password, dob, gender, intend, created_at, updated_at)
VALUES
    ('Yoni', 'Saka', 'yonisaka0@gmail.com', '123-456-7890', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1990-01-15', 'man', 'long-term partner', NOW(), NOW()),
    ('John', 'Doe', 'john.doe@example.com', '123-456-7890', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1990-01-15', 'man', 'long-term partner', NOW(), NOW()),
    ('Jane', 'Smith', 'jane.smith@example.com', '234-567-8901', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1985-05-03', 'woman', 'short-term fun', NOW(), NOW()),
    ('Michael', 'Johnson', 'michael.johnson@example.com', '345-678-9012', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1992-09-22', 'man', 'long-term', NOW(), NOW()),
    ('Emily', 'Williams', 'emily.williams@example.com', '456-789-0123', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1994-11-10', 'woman', 'short-term but long-term ok', NOW(), NOW()),
    ('David', 'Brown', 'david.brown@example.com', '567-890-1234', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1988-03-27', 'man', 'long-term partner', NOW(), NOW()),
    ('Sarah', 'Jones', 'sarah.jones@example.com', '678-901-2345', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1998-07-18', 'woman', 'new friends', NOW(), NOW()),
    ('Daniel', 'Miller', 'daniel.miller@example.com', '789-012-3456', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1991-12-08', 'man', 'short-term fun', NOW(), NOW()),
    ('Olivia', 'Davis', 'olivia.davis@example.com', '890-123-4567', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1987-04-02', 'woman', 'long-term', NOW(), NOW()),
    ('James', 'Wilson', 'james.wilson@example.com', '901-234-5678', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1993-06-14', 'man', 'still figuring it out', NOW(), NOW()),
    ('Ava', 'Moore', 'ava.moore@example.com', '012-345-6789', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1995-08-31', 'woman', 'long-term partner', NOW(), NOW()),
    ('William', 'Anderson', 'william.anderson@example.com', '123-456-7890', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1989-02-19', 'man', 'long-term', NOW(), NOW()),
    ('Sophia', 'Taylor', 'sophia.taylor@example.com', '234-567-8901', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1997-10-07', 'woman', 'short-term but long-term ok', NOW(), NOW()),
    ('Robert', 'Martinez', 'robert.martinez@example.com', '345-678-9012', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1996-04-26', 'man', 'new friends', NOW(), NOW()),
    ('Isabella', 'Hernandez', 'isabella.hernandez@example.com', '456-789-0123', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1990-11-29', 'woman', 'short-term fun', NOW(), NOW()),
    ('Ethan', 'Garcia', 'ethan.garcia@example.com', '567-890-1234', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1986-08-12', 'man', 'long-term partner', NOW(), NOW()),
    ('Mia', 'Lee', 'mia.lee@example.com', '678-901-2345', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1993-03-20', 'woman', 'still figuring it out', NOW(), NOW()),
    ('Joseph', 'Lewis', 'joseph.lewis@example.com', '789-012-3456', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1998-09-17', 'man', 'long-term', NOW(), NOW()),
    ('Charlotte', 'Lopez', 'charlotte.lopez@example.com', '890-123-4567', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1991-05-05', 'woman', 'short-term fun', NOW(), NOW()),
    ('David', 'White', 'david.white@example.com', '901-234-5678', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1987-07-23', 'man', 'new friends', NOW(), NOW()),
    ('Avery', 'Clark', 'avery.clark@example.com', '012-345-6789', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1994-12-09', 'woman', 'long-term partner', NOW(), NOW()),
    ('Liam', 'Adams', 'liam.adams@example.com', '123-456-7890', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1996-02-28', 'man', 'short-term but long-term ok', NOW(), NOW()),
    ('Emma', 'Turner', 'emma.turner@example.com', '234-567-8901', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1992-04-05', 'woman', 'short-term', NOW(), NOW()),
    ('Liam', 'Harris', 'liam.harris@example.com', '345-678-9012', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1990-11-15', 'man', 'long-term', NOW(), NOW()),
    ('Olivia', 'Robinson', 'olivia.robinson@example.com', '456-789-0123', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1993-02-25', 'woman', 'short-term fun', NOW(), NOW()),
    ('Noah', 'Walker', 'noah.walker@example.com', '567-890-1234', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1988-09-08', 'man', 'long-term partner', NOW(), NOW()),
    ('Ava', 'Young', 'ava.young@example.com', '678-901-2345', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1997-07-20', 'woman', 'short-term but long-term ok', NOW(), NOW()),
    ('William', 'Hall', 'william.hall@example.com', '789-012-3456', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1994-01-11', 'man', 'new friends', NOW(), NOW()),
    ('Sophia', 'Scott', 'sophia.scott@example.com', '890-123-4567', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1991-06-30', 'woman', 'short-term', NOW(), NOW()),
    ('James', 'King', 'james.king@example.com', '901-234-5678', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1985-08-09', 'man', 'still figuring it out', NOW(), NOW()),
    ('Isabella', 'Wright', 'isabella.wright@example.com', '012-345-6789', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1996-03-22', 'woman', 'new friends', NOW(), NOW()),
    ('Benjamin', 'Turner', 'benjamin.turner@example.com', '123-456-7890', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1998-10-05', 'man', 'short-term but long-term ok', NOW(), NOW()),
    ('Mia', 'Harris', 'mia.harris@example.com', '234-567-8901', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1989-12-15', 'woman', 'long-term partner', NOW(), NOW()),
    ('Ethan', 'Robinson', 'ethan.robinson@example.com', '345-678-9012', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1993-11-25', 'man', 'long-term', NOW(), NOW()),
    ('Charlotte', 'Walker', 'charlotte.walker@example.com', '456-789-0123', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1991-07-08', 'woman', 'short-term fun', NOW(), NOW()),
    ('Oliver', 'Young', 'oliver.young@example.com', '567-890-1234', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1987-03-20', 'man', 'long-term partner', NOW(), NOW()),
    ('Amelia', 'Hall', 'amelia.hall@example.com', '678-901-2345', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1996-01-30', 'woman', 'new friends', NOW(), NOW()),
    ('Henry', 'Scott', 'henry.scott@example.com', '789-012-3456', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1989-06-11', 'man', 'short-term fun', NOW(), NOW()),
    ('Sofia', 'King', 'sofia.king@example.com', '890-123-4567', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1995-09-22', 'woman', 'long-term', NOW(), NOW()),
    ('Jackson', 'Wright', 'jackson.wright@example.com', '901-234-5678', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1984-11-07', 'man', 'still figuring it out', NOW(), NOW()),
    ('Ava', 'Turner', 'ava.turner@example.com', '012-345-6789', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1999-02-18', 'woman', 'short-term but long-term ok', NOW(), NOW()),
    ('Liam', 'Harris', 'liam.harris@example.com', '123-456-7890', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1992-03-25', 'man', 'long-term partner', NOW(), NOW()),
    ('Olivia', 'Robinson', 'olivia.robinson@example.com', '234-567-8901', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1988-08-09', 'woman', 'short-term', NOW(), NOW()),
    ('Noah', 'Walker', 'noah.walker@example.com', '345-678-9012', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1995-01-22', 'man', 'new friends', NOW(), NOW()),
    ('Emma', 'Young', 'emma.young@example.com', '456-789-0123', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1990-10-04', 'woman', 'still figuring it out', NOW(), NOW()),
    ('William', 'Hall', 'william.hall@example.com', '567-890-1234', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1986-07-14', 'man', 'long-term', NOW(), NOW()),
    ('Sophia', 'Scott', 'sophia.scott@example.com', '678-901-2345', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1994-12-18', 'woman', 'short-term fun', NOW(), NOW()),
    ('James', 'King', 'james.king@example.com', '789-012-3456', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1997-09-21', 'man', 'long-term partner', NOW(), NOW()),
    ('Isabella', 'Wright', 'isabella.wright@example.com', '890-123-4567', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1991-05-13', 'woman', 'short-term but long-term ok', NOW(), NOW()),
    ('Benjamin', 'Turner', 'benjamin.turner@example.com', '901-234-5678', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1989-11-06', 'man', 'new friends', NOW(), NOW()),
    ('Mia', 'Harris', 'mia.harris@example.com', '012-345-6789', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8', '1996-02-28', 'woman', 'long-term', NOW(), NOW());


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
-- +goose StatementEnd
