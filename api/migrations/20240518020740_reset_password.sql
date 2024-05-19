CREATE TABLE reset_passwords (
                                id uuid PRIMARY KEY,
                                user_id uuid NOT NULL,
                                created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "reset_passwords"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");