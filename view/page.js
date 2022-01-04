const requestUpdate = async() => {
    const response = await fetch('/api');
    const json = await response.json();
    console.log(json)

    console.log('Calling modify page....')
    modifyPage(json)
}
    

function modifyPage(pageData) {
    console.log("Idle: ", pageData.idle)

    const idle = pageData.idle
    if (idle) {
        modifyIdlePage(pageData)
    } else {
        modifyNowPlaying(pageData)
    }

    
}

function modifyNowPlaying(pageData) {
    const playing = document.querySelector('.now_playing')
    //playing.style.display = "block"

    const lastPlayed = document.querySelector('.last_played');
    //lastPlayed.style.display = "none";
    
    const trackElem = document.querySelector("#current_track")
    const pageTrack = trackElem.innerHTML
    const nowPlayingTrack = pageData.nowPlaying.track

    console.log("Track: ", nowPlayingTrack);
    console.log("Page Track", pageTrack)

    switch (pageTrack) {
        case 'Track':
            modifyNowPlayingPage(pageData, trackElem)
            break;
        case nowPlayingTrack:
            console.log('no changes, same track....')
            break;
        default:
            modifyNowPlayingPage(pageData, trackElem)
    }


}

function modifyNowPlayingPage(pageData, trackElem) {
    const artist = pageData.nowPlaying.artist;
    const album = pageData.nowPlaying.album;
    const track = pageData.nowPlaying.track;
    const img = pageData.nowPlaying.art;
    const number = pageData.nowPlaying.totalPlayCount;

    const msg = document.querySelector("#message");
    //msg.innerHTML = artist;
    trackElem.innerHTML = track;

    const albumElem = document.querySelector("#current_album");
    albumElem.innerHTML = album;

    const artistElem = document.querySelector("#current_artist");
    const nowPlayingElem = document.querySelector("#artist_name");

    artistElem.innerHTML = artist;
    if (nowPlayingElem !== null) {
        nowPlayingElem.innerHTML = artist;
    }

    const imgElem = document.querySelector(".art_image");
    imgElem.setAttribute("src", img);

    const artworkElem = document.querySelector(".artwork");
    artworkElem.style.backgroundImage = 'url(' + img + ')';

    const numberElem = document.querySelector("#current_number_plays");
    const nf = new Intl.NumberFormat('en-US');
    const n = nf.format(number);
    numberElem.innerHTML = n + ' Scrobbles';

    const artistStatsElem = document.querySelector("#artist_stats")
    
    const rankingElem = document.querySelector(".ranking")
    rank = pageData.nowPlaying.ArtistRanking
    max = pageData.nowPlaying.MaxRanks
    myPlays = pageData.nowPlaying.MyArtistPlayCount

    artistStatsElem.innerHTML = "(#" + rank + " artist w/" + myPlays + " plays)"
    const fireElem = document.querySelector("#fire")
    albumPlays = pageData.nowPlaying.MyAlbumPlayCount
    var fires = Math.ceil(albumPlays / 50)
    console.log(fires)

    // this didn't seem to work on linux chromium??
    //fireElem.innerHTML = '🔥'.repeat(fires)
    var fireText = ""
    for (let i=0;i<fires;i++) {
        fireText = fireText += '🔥'
    }
    fireElem.innerHTML = fireText
}

function modifyIdlePage(pageData) {
    console.log('The player is idle');
}

(function() {
                
    setInterval(
        requestUpdate,
        5000
    );
    
   //requestUpdate()
})();