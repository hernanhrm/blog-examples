CREATE TABLE drinks
(
    id          SERIAL       NOT NULL,
    name        VARCHAR(50)  NOT NULL,
    description VARCHAR(256) NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP,
    CONSTRAINT drinks_id_pk PRIMARY KEY (id)
);

CREATE TABLE meals
(
    id          SERIAL       NOT NULL,
    name        VARCHAR(50)  NOT NULL,
    description VARCHAR(256) NOT NULL,
    drink_id    INTEGER      NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP,
    CONSTRAINT meals_id_pk PRIMARY KEY (id),
    CONSTRAINT meals_drink_id_fk FOREIGN KEY (drink_id) REFERENCES drinks (id)
);