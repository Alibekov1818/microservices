CREATE TABLE IF NOT EXISTS phones (
                                       id bigserial PRIMARY KEY,
                                       model text NOT NULL,
                                       brand text NOT NULL,
                                       year integer NOT NULL,
                                       price text NOT NULL
);