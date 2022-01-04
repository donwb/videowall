select * from history limit 1000;

select count(artist) from history where artist = 'The Smiths'
select artist from history where artist = null

select artist, count(artist) from history where artist = 'Jane''s Addiction' group by artist
select track, count(track) from history where track = 'Stop' group by track

-- count by artist
select artist, count(artist) as total
	from history group by artist order by total desc;

select dense_rank() over (order by count(artist) desc) as ranking, artist, count(artist) as total
		from history group by artist order by total desc;

-- total unique artists
select count(distinct artist) from history;

select * from artist_ranking where artist = 'The Cure'
select max(ranking) as total_rank from artist_ranking


select count(*) from history where track = 'Stop';
select track, count(track) from history group by track order by count desc;

select count(*) from history where album = 'Wish';
select album, count(album) from history group by album order by count desc;

