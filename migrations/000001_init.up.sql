BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS role
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS department
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS team
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "user"
(
    id            UUID                 DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    first_name    TEXT        NOT NULL,
    last_name     TEXT        NOT NULL,
    middle_name   TEXT,
    email         TEXT UNIQUE,
    phone         TEXT UNIQUE,
    role_id       BIGINT      NOT NULL,
    department_id BIGINT      NOT NULL,
    team_id       BIGINT      NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (role_id)
        REFERENCES role (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (department_id)
        REFERENCES department (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (team_id)
        REFERENCES team (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CHECK (email IS NOT NULL || phone IS NOT NULL)
);

COMMIT;

CREATE TABLE IF NOT EXISTS document
(
    id           UUID                 DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id      UUID        NOT NULL,
    passport_nr  VARCHAR(10) NOT NULL UNIQUE,
    snils_nr     VARCHAR(13) NOT NULL UNIQUE,
    inn_nr       VARCHAR(20) NOT NULL UNIQUE,
    insurance_nr VARCHAR     NOT NULL UNIQUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id)
        REFERENCES "user" (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS task
(
    id          UUID                 DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        NOT NULL,
    status      VARCHAR(20),
    user_id     TEXT,
    due_to      TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS education_material
(
    id           UUID                 DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name         TEXT        NOT NULL,
    description  TEXT        NOT NULL,
    category     TEXT        NOT NULL,
    material_url TEXT        NOT NULL,
    task_id      UUID        NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (task_id)
        REFERENCES task (id) ON DELETE CASCADE ON UPDATE CASCADE
);
