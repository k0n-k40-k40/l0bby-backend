ALTER TABLE users ADD CONSTRAINT unique_username UNIQUE (username);
ALTER TABLE users ADD CONSTRAINT username_length CHECK (LENGTH(username) > 3);
ALTER TABLE users ADD CONSTRAINT password_length CHECK (LENGTH(password) > 3);