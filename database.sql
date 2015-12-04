CREATE TABLE medics (
	id varchar(30) CONSTRAINT medic_pk PRIMARY KEY,
	cpf TEXT,
	name TEXT,
	crm TEXT,
	cpf TEXT,
	avatar BLOB,
	specialty TEXT,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
);

