$(document).ready(function () {
    $(".inline-playable").click(function (e) {
        e.preventDefault();
        $('a').css('font-weight', 'normal');
        e.currentTarget.setAttribute('style', 'font-weight:bold;');
        $("#inline-player-iframe").prop('src', e.currentTarget.id);
        $("#inline-player").show();
    });

    $("#close-inline-player").click(function (e) {
        e.preventDefault();
        $('a').css('font-weight', 'normal');
        $("#inline-player-iframe").prop('src', 'https://upload.wikimedia.org/wikipedia/commons/b/b1/Loading_icon.gif');
        $("#inline-player").hide();
    });
});