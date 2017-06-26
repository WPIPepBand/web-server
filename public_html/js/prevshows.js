var songs = {
    'ACPAGuOcpyw': ['Queen Opener (Bohemian Rhapsody, We Will Rock You, Another One Bites The Dust)', 'Don’t Stop Me Now', 'Bicycle Race'],
    'cU09E_E43qQ': ['James Bond Theme', 'For Your Eyes Only / Goldfinger', 'Die Another Day / Live and Let Die'],
    'xl3n6M3vZIc': [],
    'pZEy0aDUL9U': ['Theme from Generations', 'Inner Light', 'Finale'],
    '2tLxlCJmVBk': ['Overture/Christmas Eve Montage', 'This is Halloween/The Oogie Boogie Song', 'Sally’s Song/Jack’s Obsession/End Title'],
    'Ydo0OxpwgFo': ['Feels So Good/Children of Sanchez', 'Land of Make Believe', 'El Gato Triste'],
    'B4vh_cJQHYU': ['We Didn’t Start the Fire/Piano Man', 'My Life/Only the Good Die Young', 'Air (Dublinesque)', 'Just the Way You Are'],
    'vRgoPOIOPow': ['Everybody’s Everything', 'She’s Not There', 'Novus', '(Da Le) Yaleo'],
    'A7v2BQRHpm0': ['Batman', 'Raider\'s March', 'Iron Man', 'Superman'],
    'ADlD8_s68qA': ['Bridge Over Troubled Waters/You Can Call Me Al', 'Still Crazy After All These Years/Fifty Ways to Leave Your Lover', 'Late in the Evening', 'Sound of Silence/You Can Call Me Al (reprise)'],
    '-9L38iaMvk8': ['Touch Me', 'Come Sail Away', 'Crazy Little Thing Called Love', 'From Me to You'],
    '1hzCBaqpH5I': ['Best Years of our Lives', 'Accidentally in Love', 'Funky Town', 'I\'m a Believer'],
    '96pYhL4e_nU': ['Blister in the Sun', 'House of the Rising Sun', 'Soak up the Sun', 'Walking on the Sun'],
    'qsADgqAgBXM': ['Pretty Fly (For a White Guy)', 'The Impression that I Get', 'Smooth']
};

$(document).ready(function () {
    $(".playable").click(function (e) {
        e.preventDefault();

        $("#show").text($(e.currentTarget.innerHTML).text());
        var songString = '';
        songs[e.currentTarget.id].forEach((song) => {
            songString += '<tr><td>' + song + '</td></tr>';
        });
        $("#songs").html(songString);
        $("#player").prop('src', 'https://www.youtube.com/embed/' + e.currentTarget.id);
    });
});