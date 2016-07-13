import requests
import json
menu = {
     "items" :[
      {
         "vendor_id" : 1,
         "item_no" :   1 ,
         "item_name" : "Panneer",
         "item_type" : "starter",
         "item-nature" :"v" ,
         "item_description" : "wonderful",
         "price" : "240.25",
         "offer" : "10% off",
         "imageaddress" : "not available",
         "discount" : 10.0,
      },
      {
         "vendor_id" : 1,
         "item_no" :   2 ,
         "item_name" : "Dal",
         "item_type" : "main",
         "item-nature" :"v" ,
         "item_description" : "wonderful",
         "price" : "240.25",
         "offer" : "30% off",
         "imageaddress" : "not available",
         "discount" : 10.0,
      },
    ]
}

resp = requests.post("http://localhost:7070/additems", data=json.dumps(menu))
body = resp.text
print(body)
