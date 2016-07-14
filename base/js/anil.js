  $(document).ready(function() {
      $(window).scroll(function() { // check if scroll event happened
          if ($(document).scrollTop() > 50) { // check if user scrolled more than 50 from top of the browser window
              $(".ournav").css("background-color", "rgba(0, 0, 0, 0.7)"); // if yes, then change the color of class "navbar-fixed-top" to white (#f8f8f8)
          } else {
              $(".ournav").css("background-color", "transparent"); // if not, change it back to transparent
          }
      });
  });


  function register_vendor() {
      var name = document.getElementById("vendor_name").value();
      var email = document.getElementById("vendor_email").value();
      var first = document.getElementById("owner_first_name").value();
      var last_name = document.getElementById("owner_last_name").value();
      var mobile = document.getElementById("vendor_mobile").value();
      var addr = document.getElementById("vendor_address1").value();
      var city = document.getElementById("vendor_address_city").value();
      var country = document.getElementById("vendor_address_country").value();
      var pin = document.getElementById("vendor_address_postalcode").value();
      var description = document.getElementById("vendor_description").value();
      var offers = document.getElementById("vendor_offers").value();
      var vendor_name = document.getElementById("vendor_name").value();


      var json_obj = {};

      var msg = {
          "owner": first +" "+last_name,
          "vendorname": name,
          "email": email,
          "mobile": [mobile],
          "address": addr +" "+city+" "+country+" "pin ,
          "imageaddress": "not available",
          "description" : description,
          "offer": offers,
          "password": "desitadka123"
      }
      fetch('/registervendor', {
          method: 'POST',
          headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'

          },
          credentials: 'same-origin',
          body: JSON.stringify(msg)

      })


  }
