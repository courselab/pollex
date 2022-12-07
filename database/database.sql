-- Table: public.user

DROP TABLE IF EXISTS public.travel;
DROP TABLE IF EXISTS public.ride;
DROP TYPE IF EXISTS public.statusEnum;

DROP TABLE IF EXISTS public.locations;

DROP TABLE IF EXISTS public."user";
DROP TABLE IF EXISTS public.car;


-- Table: public.user

CREATE TABLE IF NOT EXISTS public."user"
(
    ID                      character       PRIMARY     KEY,
    name                    character     varying(50) COLLATE pg_catalog."default"  NOT NULL,
    nickname                character     varying(50) COLLATE pg_catalog."default",
    "driver_ratingAVG"      numeric(1,0),
    "driver_ratingCount"    integer,
    "driver_tripCount"    integer,
    "passenger_ratingAVG"   numeric(1,0)  NOT NULL,
    "passenger_ratingCount" integer       NOT NULL,
    "passenger_tripCount" integer       NOT NULL,
    "isDriver"              boolean       NOT NULL,
    "car_licensePlate"      character     varying(50) COLLATE pg_catalog."default",
    CONSTRAINT car_fk FOREIGN KEY (car_licensePlate)
        REFERENCES public.car (car_licensePlate) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION,

)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."user"
    OWNER to postgres;
	
-- Table: public.car
CREATE TABLE IF NO EXISTS public.car
(
    car_licensePlate    character   varying(50) COLLATE pg_catalog."default",
    car_model           character   varying(50) COLLATE pg_catalog."default",
    car_color           character   varying(50) COLLATE pg_catalog."default",
    CONSTRAINT car_pk PRIMARY KEY (car_licensePlate)
)
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
    driver_id character NOT NULL,
    "start" integer NOT NULL,
    "end" integer NOT NULL,
    "when" timestamp without time zone NOT NULL,
	"status" statusEnum,
    "maxSeats" integer NOT NULL,
	CONSTRAINT travel_pk PRIMARY KEY (id),
	CONSTRAINT driver_fk FOREIGN KEY (driver_id)
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
 
 -- Table: public.ride

CREATE TABLE IF NOT EXISTS public.ride
(
    travel_id integer NOT NULL,
    passenger_id character NOT NULL
    "driverRating" integer,
    "passengerRating" integer,
	CONSTRAINT ride_pk PRIMARY KEY (travel_id, passenger_id),
	CONSTRAINT travel_fk FOREIGN KEY (travel_id)
        REFERENCES public.travel (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
	CONSTRAINT passenger_fk FOREIGN KEY (passenger)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.ride
    OWNER to postgres;
