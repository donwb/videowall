CREATE TABLE pmrc (
	pmrcid BIGINT PRIMARY KEY NOT NULL,
	lookup VARCHAR NOT NULL,
	safename VARCHAR NOT NULL
)
CREATE SEQUENCE SQLPro_id_seq;
ALTER TABLE ONLY public.pmrc ALTER COLUMN pmrcid SET DEFAULT nextval('SQLPro_id_seq');
ALTER SEQUENCE SQLPro_id_seq OWNED BY public.pmrc.id;

ALTER TABLE public.pmrc RENAME COLUMN sampleid TO pmrcid;
ALTER TABLE public.pmrc ADD COLUMN lookup text NOT NULL;
ALTER TABLE public.pmrc ADD COLUMN safename text NOT NULL;

insert into pmrc (lookup, safename) values ('Revolting Cocks', 'RevCo');
insert into pmrc (lookup, safename) values ('1000 Homo DJs', '1000 **** DJs');
insert into pmrc (lookup, safename) values ('Beers, Steers + Queers (Deluxe Edition)', 'Beers, Steers & (Deluxe Edition)');
insert into pmrc (lookup, safename) values ('Beers, Steers + Queers', 'Beers, Steers & *');

insert into pmrc (lookup, safename) values ('You Goddamned Son of a Bitch', 'You G*******d Son of a B***h');

