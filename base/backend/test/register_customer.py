import requests
import json
customer_to_register = {
     "customer_name" : "aman",
     "emailid" : "dt@gmail.com",
     "mobile" : ["1234567810"],
     "address" : "Hydrabaad",
     "password" : "aman123"
}

resp = requests.get("http://localhost:7070/registercustomer", data=json.dumps(customer_to_register))
body = resp.text
print(body)
