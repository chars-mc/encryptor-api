-- text, file, etc...
CREATE TABLE IF NOT EXISTS data_type (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(10) UNIQUE NOT NULL
);

-- AES, blowfish, etc...
CREATE TABLE IF NOT EXISTS algorithm (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(10) UNIQUE NOT NULL
);

-- 
CREATE TABLE IF NOT EXISTS data (
  id INT PRIMARY KEY AUTO_INCREMENT,
  content VARCHAR(150) NOT NULL,
  id_data_type INT NOT NULL,
  id_algorithm INT NOT NULL,
	CONSTRAINT id_data_type FOREIGN KEY (id_data_type) REFERENCES data_type(id),
	CONSTRAINT id_algorithm FOREIGN KEY (id_algorithm) REFERENCES algorithm(id)
);
