select * from history
select count(artist) from history where artist = 'The Smiths'
select artist from history where artist = null

select artist, count(artist) from history where artist = 'Jane''s Addiction' group by artist
select track, count(track) from history where track = 'Stop' group by track