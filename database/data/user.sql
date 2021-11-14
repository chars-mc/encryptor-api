INSERT INTO user_role
VALUES (1, "Encryptor"), (2, "Decryptor");

INSERT INTO user(id, username, password, id_user_role)
VALUES
	(1, "encryptor_test", "encryptor_test", 1),
	(2, "decryptor_test", "decryptor_test", 2);
