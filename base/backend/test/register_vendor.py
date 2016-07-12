import requests
import json
vendor_to_register = {
     "owner" : "Rahul Sharma",
     "vendorname" : "Desi Tadka", 
     "email" : "dt@gmail.com",
     "mobile" : ["1234567890"],
     "address" : "Hamirpur (H.P.)",
     "Offer" : "free meal for couples",
     "image" : "not available",
     "password" : "desitadka123"    
}

resp = requests.post("http://localhost:7070/registervendor", data=json.dumps(vendor_to_register))
body = resp.text
print(body)
