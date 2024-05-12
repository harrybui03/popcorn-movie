CREATE TYPE "role" AS ENUM (
  'CUSTOMER',
  'STAFF',
  'TICKET_MANAGER',
  'ADMIN'
);

CREATE TYPE "status" AS ENUM (
  'UPCOMING',
  'ONGOING',
  'OVER'
);

CREATE TYPE "category" AS ENUM (
  'STANDARD',
  'COUPLE'
);

CREATE TABLE IF NOT EXISTS "users"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "email" varchar not null,
    "displayname" varchar not null,
    "password" varchar not null,
    "is_locked" bool not null,
    "role" role not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "transactions"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "total" float not null,
    "user_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "tickets"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "price" float not null,
    "is_booked" bool not null,
    "transaction_id" uuid,
    "seat_id" uuid not null,
    "show_time_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "movies"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "title" varchar not null,
    "genre" varchar not null,
    "status" status not null,
    "language" varchar not null,
    "director" varchar not null,
    "cast" varchar not null,
    "poster" varchar not null,
    "rated" float not null,
    "duration" varchar not null,
    "trailer" varchar not null,
    "opening_day" timestamp not null,
    "story" varchar not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "comments"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "description" varchar not null,
    "rating" float not null,
    "movie_id" uuid not null,
    "user_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "rooms"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "room_number" integer not null not null,
    "theater_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "theaters"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "address" varchar not null,
    "name" varchar not null,
    "phone_number" varchar not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "seats"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "seat_number" varchar not null,
    "category" category not null,
    "room_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "show_times"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "start_at" timestamp not null,
    "end_at" timestamp not null,
    "movie_id" uuid not null,
    "room_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "foods"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "name" varchar not null,
    "price" float not null,
    "image" varchar not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "food_order_lines"
(
    id
    uuid
    not
    null
    default
    uuid_generate_v4
(
) PRIMARY KEY,
    "quantity" integer not null,
    "food_id" uuid not null,
    "transaction_id" uuid not null,
    created_at timestamp without time zone default now
(
),
    updated_at timestamp
                         without time zone default now
(
)
    );

CREATE TABLE IF NOT EXISTS "sessions"
(
    id
    uuid
    not
    null
    PRIMARY
    KEY,
    "user_id"
    uuid
    NOT
    NULL,
    "refresh_token"
    varchar
    NOT
    NULL,
    "expires_at"
    timestamp
    NOT
    NULL,
    created_at
    timestamp
    without
    time
    zone
    default
    now
(
)
    );

ALTER TABLE "transactions"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tickets"
    ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "tickets"
    ADD FOREIGN KEY ("seat_id") REFERENCES "seats" ("id");

ALTER TABLE "tickets"
    ADD FOREIGN KEY ("show_time_id") REFERENCES "show_times" ("id");

ALTER TABLE "comments"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments"
    ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "rooms"
    ADD FOREIGN KEY ("theater_id") REFERENCES "theaters" ("id");

ALTER TABLE "seats"
    ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "show_times"
    ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "show_times"
    ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "food_order_lines"
    ADD FOREIGN KEY ("food_id") REFERENCES "foods" ("id");

ALTER TABLE "food_order_lines"
    ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");
