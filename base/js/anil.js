  // $(document).ready(function() {
  //     $(window).scroll(function() { // check if scroll event happened
  //         if ($(document).scrollTop() > 50) { // check if user scrolled more than 50 from top of the browser window
  //             $(".ournav").css("background-color", "rgba(0, 0, 0, 0.7)"); // if yes, then change the color of class "navbar-fixed-top" to white (#f8f8f8)
  //         } else {
  //             $(".ournav").css("background-color", "transparent"); // if not, change it back to transparent
  //         }
  //     });

  // });

// <!-- <li><a href="#">A9</a></li> -->
  window.addEventListener('load',function() {
      //To add hotel-names to  drop drown
      var drop = document.getElementsByClassName("hotel-names-list");
      fetch('/getvendors', {
          method: 'GET',
          headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
          },
          credentials: 'same-origin',
      }).then(function(response) {
          return response.json();
      }).then(function(hotels) {
          for (var i=0;i<hotels.length;i++) {
                    console.log(hotels[i]["vendor_id"],hotels[i]["vendorname"]);
                    var listItem = document.createElement("li");
                    var a = document.createElement("a");
                     // a.SetAttribute("href","#");
                     a.setAttribute("id",hotels[i]["vendor_id"]);
                     a.innerHTML = hotels[i]["vendorname"];
                     listItem.appendChild(a);
                     drop[0].appendChild(listItem);
          }
      }, function(err) {
        console.log(err);
      })
  })

  //Function to register vendor Called from dashboard
  function register_vendor() {
      var name = document.getElementById("vendor_name").value;
      var email = document.getElementById("vendor_email").value;
      var first = document.getElementById("owner_first_name").value;
      var last_name = document.getElementById("owner_last_name").value;
      var mobile = document.getElementById("vendor_mobile").value;
      var addr = document.getElementById("vendor_address1").value;
      var city = document.getElementById("vendor_address_city").value;
      var country = document.getElementById("vendor_address_country").value;
      var pin = document.getElementById("vendor_address_postalcode").value;
      var description = document.getElementById("vendor_description").value;
      var offers = document.getElementById("vendor_offers").value;
      var vendor_name = document.getElementById("vendor_name").value;



      var msg = {
          "owner": first + " " + last_name,
          "vendorname": name,
          "email": email,
          "mobile": [mobile],
          "address": addr + " " + city + " " + country + " " + pin,
          "imageaddress": "not available",
          "description": description,
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

  // var count =0 ;
//Function to add items to Menu 
  function add_items() {

      var name = document.getElementById('name').value;
      var price = document.getElementById('price').value;
      var discount = document.getElementById('discount').value;
      var nature;
      if (document.getElementById('veg').checked){
          nature= true;
      }else if(document.getElementById('n_veg').checked){
          nature = false;
      }
      var Itype;
      if (document.getElementById('starter').checked){
          Itype = 'starter';
      }else if(document.getElementById('main').checked){
          Itype = 'main';
      }else if(document.getElementById('desert').checked){
          Itype = 'desert';
      }
      var description = document.getElementById('description').value;
        
      var tableid = document.getElementById("table-body");   

      var nrow = document.createElement('tr');

      var c1 = document.createElement('td');
      c1.setAttribute('class',"count");
      nrow.appendChild(c1);

      var c2 = document.createElement('td'); 
      c2.setAttribute('spellcheck',false); 
      c2.setAttribute('contenteditable',true);
      c2.innerHTML = name; 
      nrow.appendChild(c2);

      var c3 = document.createElement('td'); 
      c3.setAttribute('spellcheck',false); 
      c3.setAttribute('contenteditable',true);
      c3.innerHTML = price; 
      nrow.appendChild(c3);

      var c4 = document.createElement('td'); 
      c4.setAttribute('spellcheck',false); 
      c4.setAttribute('contenteditable',true);
      c4.innerHTML = Itype; 
      nrow.appendChild(c4);

      var c5 = document.createElement('td'); 
      c5.setAttribute('spellcheck',false); 
      c5.setAttribute('contenteditable',true);
      c5.innerHTML = nature; 
      nrow.appendChild(c5);

      var c6 = document.createElement('td'); 
      c6.setAttribute('spellcheck',false); 
      c6.setAttribute('contenteditable',true);
      c6.innerHTML = discount; 
      nrow.appendChild(c6);

      tableid.appendChild(nrow);

      document.getElementById('item-form').reset();
      // rows += "<tr><td>" + name + "</td><td>" + gender + "</td><td>" + age + "</td><td>" + city + "</td></tr>";
      // $(rows).appendTo("#vendor_menu_items tbody");
  }

function submit_items(){
   var table = document.getElementById('vendor_menu_items');
   var totalrows= document.getElementById('vendor_menu_items').rows.length;


}