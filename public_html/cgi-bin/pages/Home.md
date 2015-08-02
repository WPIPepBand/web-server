
Welcome to the home page of the WPI Pep Band!

### Time remaining until band camp starts:


<div id="band-camp-timer"></div>
<div id="band-camp-message" class="alert alert-info" style="display: none;">
    Band camp is currently taking place.  Get off your phone!
</div>

<script type="text/javascript">
    $(function() {
        var bandCamp = new Date(2015, 7, 18, 9);
        var $bandCampSection = $('bandCampSection');

        $(window).on("resize", function () {
            var dif = Math.max($(window).height() - $bandCampSection.height(), 0);
            var padding = Math.floor(dif / 2) + 'px';
            $bandCampSection.css({ 'padding-top': padding, 'padding-bottom': padding });
        }).trigger("resize");

        $('#band-camp-timer').mbComingsoon({
            expiryDate: bandCamp,
            speed: 100,
        });

        setTimeout(function () {
            $(window).resize();
        }, 200);

        if (new Date() > bandCamp) {
            $("#band-camp-message").css("display", "block");
        }
    });
</script>


Our halftime show for fall of 2015 will be ***James Bond***.

Listen to the show [here.](http://www.jwpepper.com/marching-band-the-music-of-james-bond-show.list)

To help prepare the halftime show, we will be holding a band camp from **August 18th - 21st**, which is immediately before New Student Orientation. If you will be joining us as a marching member, you will be required to attend. We will need you to let us know whether you are interested in marching with us by **Wednesday, July 15th**, in order to prepare for your arrival and band camp.

The band is a completely student-run organization and performs at all home WPI football and basketball games, some away games, and we occasionally get invited to play at the Worcester Sharks games. The band welcomes members of all musical ability-- Don't be discouraged if you've never played or marched before, we can help you learn! We also have a color guard which is open to anyone who is willing to learn to spin (click [here](Color%20Guard) to learn more about Color Guard!). The Pep Band also offers physical education credit for those participating in football and/or basketball season. Participation in Pep Band for four seasons will therefore fulfill your physical education requirement. Despite being available for academic credit, members still find the rehearsals and performances to be very casual and loads of fun. The Pep Band is a spirited bunch that has much to offer its members, including friendship, pizza, movie nights, parties, and most importantly, good music!

Currently, we are looking for new members for football season! Please feel free to contact the band officers at [pepoff@wpi.edu](mailto:pepoff@wpi.edu) if you have any questions or you are interested in joining us for football or basketball season. Additionally, we have a Winter Guard so please contact the officers if you would like to participate in that as well.

<div class="alert alert-success">The WPI Pep Band has recently received the honor of being named the Organization of the Year. Way to go, Pep Band!</div>
