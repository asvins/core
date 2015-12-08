CREATE TABLE medics (
	id serial primary key,
	cpf TEXT,
	name TEXT,
	crm TEXT,
	specialty TEXT,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	deleted_at TIMESTAMP WITHOUT TIME ZONE,
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

CREATE TABLE receipts (
	id serial primary key,
	treatment_id INTEGER,
	status int,
	file_path TEXT,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	deleted_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
)

CREATE TABLE patients (
  id serial primary key,
	cpf TEXT,
	name TEXT,
	email TEXT,
	medical_history TEXT,
	Weight TEXT,
	Gender INTEGER,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	deleted_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
)
	
CREATE TABLE pharmacists (
  id serial primary key,
	crf TEXT,
	name TEXT,
	email TEXT,
	created_at TIMESTAMP WITHOUT TIME ZONE,
	deleted_at TIMESTAMP WITHOUT TIME ZONE,
	updated_at TIMESTAMP WITHOUT TIME ZONE
)
