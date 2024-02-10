-- +goose Up
-- +goose StatementBegin
INSERT INTO USERS (email, hash) VALUES ('rachelle.huel@ethereal.email', 'C6s2S9qe6WrTMB7z3u');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM USERS WHERE EMAIL = 'rachelle.huel@ethereal.email';
-- +goose StatementEnd
