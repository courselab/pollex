-- Table: public.user

DROP TABLE IF EXISTS public.travel;
DROP TYPE IF EXISTS public.statusEnum;

DROP TABLE IF EXISTS public.locations;

DROP TABLE IF EXISTS public."user";


CREATE TABLE IF NOT EXISTS public."user"
(
    ID                      integer       PRIMARY     KEY,
    name                    character     varying(50) COLLATE pg_catalog."default"  NOT NULL,
    nickname                character     varying(50) COLLATE pg_catalog."default",
    "driver_ratingAVG"      numeric(1,0),
    "driver_ratingCount"    integer,
    "driver_tripCount"    integer,
    "passenger_ratingAVG"   numeric(1,0)  NOT NULL,
    "passenger_ratingCount" integer       NOT NULL,
    "passenger_tripCount" integer       NOT NULL,
    "isDriver"              boolean       NOT NULL,
    car_model               character     varying(50) COLLATE pg_catalog."default",
    car_color               character     varying(50) COLLATE pg_catalog."default",
    "car_licensePlate"      character     varying(50) COLLATE pg_catalog."default"

)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."user"
    OWNER to postgres;
	
-- Table: public.locations

CREATE TABLE IF NOT EXISTS public.locations
(
    id integer PRIMARY KEY,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    coords point NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.locations
    OWNER to postgres;
	
-- Table: public.travel

CREATE TYPE statusEnum AS ENUM ('pending', 'accepted', 'rejected', 'cancelled');
CREATE TABLE IF NOT EXISTS public.travel
(
    id integer NOT NULL,
    seat integer NOT NULL,
    driver integer NOT NULL,
    passenger integer NOT NULL,
    "start" integer NOT NULL,
    "end" integer NOT NULL,
    "when" timestamp without time zone NOT NULL,
	"status" statusEnum,
    "driverRating" integer,
    "passengerRating" integer,
	CONSTRAINT travel_fk PRIMARY KEY (id, seat),
	CONSTRAINT driver_fk FOREIGN KEY (driver)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
	CONSTRAINT passenger_fk FOREIGN KEY (passenger)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
	CONSTRAINT start_fk FOREIGN KEY ("start")
        REFERENCES public."locations" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
	CONSTRAINT end_fk FOREIGN KEY ("end")
        REFERENCES public."locations" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.travel
    OWNER to postgres;
