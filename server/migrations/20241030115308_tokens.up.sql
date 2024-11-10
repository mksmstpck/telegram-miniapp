create table if not exists tokens (
    telegramID varchar(10) primary key,
    token varchar(36) unique
);