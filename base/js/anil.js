  $(document).ready(function(){
      $(window).scroll(function() { // check if scroll event happened
        if ($(document).scrollTop() > 50) { // check if user scrolled more than 50 from top of the browser window
          $(".ournav").css("background-color", "rgba(0, 0, 0, 0.5)"); // if yes, then change the color of class "navbar-fixed-top" to white (#f8f8f8)
        } else {
          $(".ournav").css("background-color", "transparent"); // if not, change it back to transparent
        }
      });
    });


    var elem = document.querySelector('.ournav');
    var animation = elem.animate([
        {background-color: rgba(0, 0, 0, 0.1), offset: },
        {background-color: rgba(0, 0, 0, 0.2), offset:},
        {background-color: rgba(0, 0, 0, 0.3), offset:},
        {background-color: rgba(0, 0, 0, 0.4), offset:},
    ], {
        // direction: 'alternate',
        // duration: 500,
        // iterations: Infinity
    });