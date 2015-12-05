CREATE TABLE medics (
	id serial primary key,
	cpf TEXT,
	name TEXT,
	crm TEXT,
	cpf TEXT,
	avatar BLOB,
	specialty TEXT,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE medications (
	id serial primary key,
	active_agent TEXT,
	label INTEGER,
	dosage TEXT,
	bula TEXT,
	type INTEGER,
	intake_mean INTEGER,
	name TEXT,
	br_register TEXT,
	terapeutic_class TEXT,
	manufacturer TEXT
);
