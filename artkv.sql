CREATE TABLE artkv (
	id BIGINT PRIMARY KEY NOT NULL,
	album VARCHAR NOT NULL,
	arturl VARCHAR NOT NULL
)

CREATE SEQUENCE SQLPro_id_seq;
ALTER TABLE ONLY public.artkv ALTER COLUMN id SET DEFAULT nextval('SQLPro_id_seq');
ALTER SEQUENCE SQLPro_id_seq OWNED BY public.artkv.id;


insert into artkv (album, arturl) values ('2019-08-08 Bridgestone Arena', '2019-08-08 Bridgestone Arena.jpg');
insert into artkv (album, arturl) values ('2019-07-26 CJ Madison Square Garden', '2019-07-26 CJ Madison Square Garden.jpg');
insert into artkv (album, arturl) values ('2019-07-25 Madison Square Garden', '2019-07-25 Madison Square Garden.jpg');
insert into artkv (album, arturl) values ('2019-09-04 Enterprise Center, St Louis, MO', '2019-09-04 Enterprise Center, St Louis, MO.jpg');
insert into artkv (album, arturl) values ('2019-08-11 State Farm Arena Atlanta', '2019-08-11 State Farm Arena Atlanta.jpg');
