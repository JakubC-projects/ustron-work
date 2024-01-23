CREATE TABLE IF NOT EXISTS "public"."persons" (
	"person_id" INT NOT NULL,
	"display_name" TEXT NOT NULL,
	"birth_date" DATE NOT NULL,
	"team" TEXT NOT NULL,
	"role" TEXT NOT NULL,
	CONSTRAINT "persons_pkey" PRIMARY KEY (person_id)
);

CREATE TABLE IF NOT EXISTS "public"."sessions" (
	"uid" UUID NOT NULL,
	"person_id" INT NOT NULL,
	"expiry" TIMESTAMP NOT NULL,
	CONSTRAINT "sessions_pkey" PRIMARY KEY (uid)
);

CREATE TABLE IF NOT EXISTS "public"."registrations" (
	"uid" UUID NOT NULL,
	"person_id" INT NOT NULL,
	"team" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"date" DATE NOT NULL,
	"hourly_wage" INT NULL,
	"hours" REAL NULL,
	"paid_sum" INT NULL,
	"description" TEXT NULL,
	CONSTRAINT "registrations_pkey" PRIMARY KEY (uid),
	CONSTRAINT "registrations_persons_fk" FOREIGN KEY (person_id) REFERENCES persons(person_id)
);

CREATE TABLE IF NOT EXISTS "public"."on_track" (
	"team" TEXT NOT NULL,
	"status" INT NOT NULL,
	CONSTRAINT "on_track_pkey" PRIMARY KEY (team)
);