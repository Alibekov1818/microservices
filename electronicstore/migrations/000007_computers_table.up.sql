CREATE TABLE IF NOT EXISTS computers (
                                      id bigserial PRIMARY KEY,
                                      model text NOT NULL,
                                      cpu text NOT NULL,
                                      memory text NOT NULL,
                                      price text NOT NULL
);