CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/**********************************
 * Create Tables
 **********************************/
CREATE table users (
    user_id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    icon VARCHAR(255)
);

CREATE table user_passwords (
    user_id uuid NOT NULL,
    password VARCHAR(255) NOT NULL,
    algorithm VARCHAR(20) NOT NULL,
    enabled BOOLEAN DEFAULT true NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (user_id, enabled),
    CONSTRAINT foreign_user FOREIGN KEY(user_id) REFERENCES users(user_id) ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE UNIQUE INDEX user_passwords_user_id_enabledx ON user_passwords(user_id, enabled);

/**********************************
 * Auto Update Timestamp
 **********************************/

CREATE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at_update
    BEFORE UPDATE
    ON
        users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

CREATE TRIGGER update_user_passwords_updated_at_update
    BEFORE UPDATE
    ON
        user_passwords
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();
