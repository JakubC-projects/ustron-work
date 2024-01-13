CREATE TABLE IF NOT EXISTS "public"."persons" (
	"uid" UUID NOT NULL,
	"display_name" TEXT NOT NULL,
	"team" TEXT NOT NULL,
	"role" TEXT NOT NULL,
	CONSTRAINT "persons_pkey" PRIMARY KEY (uid)
);

CREATE TABLE IF NOT EXISTS "public"."sessions" (
	"uid" UUID NOT NULL,
	"person_uid" UUID NOT NULL,
	"expiry" TIMESTAMP NOT NULL,
	CONSTRAINT "sessions_pkey" PRIMARY KEY (uid),
	CONSTRAINT "session_persons_fk" FOREIGN KEY (person_uid) REFERENCES persons(uid)
);

CREATE TABLE IF NOT EXISTS "public"."registrations" (
	"uid" UUID NOT NULL,
	"person_uid" UUID NOT NULL,
	"team" TEXT NOT NULL,
	"registration_type" TEXT NOT NULL,
	"hourly_wage" INT NULL,
	"hours" INT NULL,
	"paid_sum" INT NULL,
	CONSTRAINT "registrations_pkey" PRIMARY KEY (uid),
	CONSTRAINT "registrations_persons_fk" FOREIGN KEY (person_uid) REFERENCES persons(uid)
);