CREATE TABLE IF NOT EXISTS "public"."rounds" (
	"id" int NOT NULL,
	"start_date" timestamp NOT NULL,
	"end_date" timestamp NOT NULL,
	"freeze_start_date" timestamp NULL,

	CONSTRAINT "rounds_pkey" PRIMARY KEY (id)
);

INSERT INTO rounds(id, start_date, end_date, freeze_start_date) VALUES 
    (1, '2024-01-01', '2024-03-06 21:00:00+01:00', '2024-03-06 18:00:00+01:00'),
    (2, '2024-03-06 21:00:00+01:00', '2024-05-15 21:00:00+01:00', '2024-05-15 18:00:00+01:00');

ALTER TABLE "on_track" 
    ADD COLUMN "round_id" INT NOT NULL DEFAULT 1,
    ADD CONSTRAINT "on_track_round_fk" FOREIGN KEY (round_id) REFERENCES rounds(id);