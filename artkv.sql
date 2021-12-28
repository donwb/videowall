CREATE TABLE artkv (
	id BIGINT PRIMARY KEY NOT NULL,
	album VARCHAR NOT NULL,
	arturl VARCHAR NOT NULL
)

CREATE SEQUENCE SQLPro_id_seq;
ALTER TABLE ONLY public.artkv ALTER COLUMN id SET DEFAULT nextval('SQLPro_id_seq');
ALTER SEQUENCE SQLPro_id_seq OWNED BY public.artkv.id;

insert into artkv (album, arturl) values ('', '');

insert into artkv (album, arturl) values ('2019-08-08 Bridgestone Arena', '2019-08-08 Bridgestone Arena.jpg');
insert into artkv (album, arturl) values ('2019-07-26 CJ Madison Square Garden', '2019-07-26 CJ Madison Square Garden.jpg');
insert into artkv (album, arturl) values ('2019-07-25 Madison Square Garden', '2019-07-25 Madison Square Garden.jpg');
insert into artkv (album, arturl) values ('2019-09-04 Enterprise Center, St Louis, MO', '2019-09-04 Enterprise Center, St Louis, MO.jpg');
insert into artkv (album, arturl) values ('2019-08-11 State Farm Arena Atlanta', '2019-08-11 State Farm Arena Atlanta.jpg');
insert into artkv (album, arturl) values ('The Cure in Orange', 'The Cure in Orange.jpg');
insert into artkv (album, arturl) values ('Indigenous Species', 'Indigenous Species.jpg');
insert into artkv (album, arturl) values ('Dreamstorm Collection', 'Dreamstorm Collection.jpg');
insert into artkv (album, arturl) values ('ZOO TV Tour From Sydney', 'ZOO TV Tour From Sydney.jpg');
insert into artkv (album, arturl) values ('The Swing', 'The Swing.jpg');
insert into artkv (album, arturl) values ('Any Given Thursday', 'Any Given Thursday.jpg');
insert into artkv (album, arturl) values ('As Is Volume I', 'As Is Volume I.jpg');
insert into artkv (album, arturl) values ('As/Is (Live @ Houston, TX - 7/24/04)', 'As_Is (Live @ Houston, TX - 7_24_04).jpg');
insert into artkv (album, arturl) values ('As/Is (Live @ Mountain View, CA - 7/16/04)', 'As_Is (Live @ Mountain View, CA - 7_16_04).jpg');
insert into artkv (album, arturl) values ('As/Is (Live @ Philadelphia, PA & Hartford, CT - 8/14/04-8/15/04)', 'As_Is (Live @ Philadelphia, PA & Hartford, CT - 8_14_04-8_15_04).jpg');
insert into artkv (album, arturl) values ('As/Is (Live)', 'As_Is (Live).jpg');
insert into artkv (album, arturl) values ('Bud Light Dive Bar Tour 2017', 'Bud Light Dive Bar Tour 2017.jpg');
insert into artkv (album, arturl) values ('Heavier Things - Bonus Disc', 'Heavier Things - Bonus Disc.jpg');
insert into artkv (album, arturl) values ('Live at Hi Fi Buys', 'Live at Hi Fi Buys.jpg');
insert into artkv (album, arturl) values ('Live at Orlando 1-27-07', 'Live at Orlando 1-27-07.jpg');
insert into artkv (album, arturl) values ('Live at the Mile High Music Fe', 'Live at the Mile High Music Fe.jpg');
insert into artkv (album, arturl) values ('Live at Winooski - 02-25-02', 'Live at Winooski - 02-25-02.jpg');
insert into artkv (album, arturl) values ('Live from the XLounge', 'Live from the XLounge.jpg');
insert into artkv (album, arturl) values ('Live in Atlanta-2010-03-17', 'Live in Atlanta-2010-03-17.jpg');
insert into artkv (album, arturl) values ('Live in Noblesville SBD - 8.10.13', 'Live in Noblesville SBD - 8.10.13.jpg');
insert into artkv (album, arturl) values ('The Lost Tapes - Eddie''s Attic ''05 (Night 2)', 'The Lost Tapes - Eddie''s Attic ''05 (Night 2).jpg');
insert into artkv (album, arturl) values ('2004-06-29 - Live on PBS SoundStage', '2004-06-29 - Live on PBS SoundStage.jpg');
insert into artkv (album, arturl) values ('2004-07-07 - Red Rocks (HT)', '2004-07-07 - Red Rocks (HT).jpg');
insert into artkv (album, arturl) values ('2006-10-14 - Chastain Park Amphitheater REMASTERED', '2006-10-14 - Chastain Park Amphitheater REMASTERED');
insert into artkv (album, arturl) values ('2007-16-07 - Red Rocks Amphitheatre', '2007-16-07 - Red Rocks Amphitheatre.jpg');
insert into artkv (album, arturl) values ('2008-08-02 Woodlands Amphitheater REMASTERED', '2008-08-02 Woodlands Amphitheater REMASTERED.jpg');
insert into artkv (album, arturl) values ('2010-03-17 - Philips Arena', '2010-03-17 - Philips Arena.jpg');
insert into artkv (album, arturl) values ('2010-07-21 - Jones Beach Amphitheater', '2010-07-21 - Jones Beach Amphitheater.jpg');
insert into artkv (album, arturl) values ('2017-08-08 - Bridgestone Arena', '2017-08-08 - Bridgestone Arena.jpg');
insert into artkv (album, arturl) values ('2017-08-10 - Lakewood Amphitheater', '2017-08-10 - Lakewood Amphitheater.jpg');
insert into artkv (album, arturl) values ('2017-08-23 - Jones Beach Amphitheater', '2017-08-23 - Jones Beach Amphitheater.jpg');
insert into artkv (album, arturl) values ('Heard It In A Past Life', 'Heard It In A Past Life.jpg');
insert into artkv (album, arturl) values ('Singles', 'Singles.jpg');
insert into artkv (album, arturl) values ('Substance [Disc 1]', 'Substance [Disc 1].jpg');
insert into artkv (album, arturl) values ('Substance [Disc 2]', 'Substance [Disc 1].jpg');
insert into artkv (album, arturl) values ('Substance', 'Substance.jpg');
insert into artkv (album, arturl) values ('(What''s the Story) Morning Glory? (Deluxe Edition)', '(What''s the Story) Morning Glory_ (Deluxe Edition).jpg');
insert into artkv (album, arturl) values ('Definitely Maybe (Deluxe Edition) [Remastered]', 'Definitely Maybe (Deluxe Edition) [Remastered].jpg');
insert into artkv (album, arturl) values ('Be Here Now (Deluxe Remastered Edition)', 'Be Here Now (Deluxe Remastered Edition).jpg');
insert into artkv (album, arturl) values ('The Joshua Tree (30th Anniversary Super Deluxe Edition)', 'The Joshua Tree (30th Anniversary Super Deluxe Edition).jpg');
insert into artkv (album, arturl) values ('Curaetion-25: From There to Here From Here to There (Live)', 'Curaetion-25_ From There to Here From Here to There (Live).jpg');
insert into artkv (album, arturl) values ('Anniversary: 1978 - 2018 Live In Hyde Park London (Live)', 'Anniversary_ 1978 - 2018 Live In Hyde Park London (Live).jpg');
insert into artkv (album, arturl) values ('The Best of INXS', 'The Best of INXS.jpg');
insert into artkv (album, arturl) values ('Bestival Live - 2016', 'Bestival Live - 2016.jpg');
insert into artkv (album, arturl) values ('Bestival Live 2011', 'Bestival Live 2011.jpg');
insert into artkv (album, arturl) values ('Festival Live', 'Festival Live.jpg');
insert into artkv (album, arturl) values ('Kiss Me, Kiss Me, Kiss Me', 'Kiss Me, Kiss Me, Kiss Me.jpg');
insert into artkv (album, arturl) values ('Live at Lallapalooza', 'Live at Lallapalooza.jpg');
insert into artkv (album, arturl) values ('One Million Virgins', 'One Million Virgins.jpg');
insert into artkv (album, arturl) values ('Paris (Live)', 'Paris (Live).jpg');
insert into artkv (album, arturl) values ('Sessions@AOL - EP', 'Sessions@AOL - EP.jpg');
insert into artkv (album, arturl) values ('Show (Live)', 'Show (Live).jpg');
insert into artkv (album, arturl) values ('Staring at the Sea: The Singles', 'Staring at the Sea: The Singles.jpg');
insert into artkv (album, arturl) values ('Wish', 'Wish.jpg');
insert into artkv (album, arturl) values ('Wished', 'Wished.jpg');
insert into artkv (album, arturl) values ('1992-06-02-Orlando FL', '1992-06-02-Orlando FL.jpg');
insert into artkv (album, arturl) values ('2016-05-20, Sleep Train Amph, Chula Vista, CA', '2016-05-20, Sleep Train Amph, Chula Vista, CA.jpg');
insert into artkv (album, arturl) values ('2016-05-26, Shoreline Amphitheatre, Mountain View, CA', '2016-05-20, Sleep Train Amph, Chula Vista, CA.jpg');
insert into artkv (album, arturl) values ('2016-06-05 Fiddler''s Green Englewood, CO', '2016-05-20, Sleep Train Amph, Chula Vista, CA.jpg');
insert into artkv (album, arturl) values ('2019-06-21 Southside Festival', '2019-06-21 Southside Festival.jpg');


update artkv
set arturl = 'Live at Lallapalooza.png'
where album = 'Live at Lallapalooza.jpg'

update artkv
set arturl = 'Bestival Live - 2016.png'
where album = 'Bestival Live - 2016'

update artkv
set arturl = '2010-07-21 - Jones Beach Amphitheater.png'
where album = '2010-07-21 - Jones Beach Amphitheater'

update artkv
set arturl = 'Bud Light Dive Bar Tour 2017.png'
where album = 'Bud Light Dive Bar Tour 2017'

update artkv
set arturl = '2019-08-08 Bridgestone Arena.png'
where album = '2019-08-08 Bridgestone Arena'


