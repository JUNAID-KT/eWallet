**Get Ethereum Block related data**
----
    
* **Description**
    
       This API will fetch transaction details based on user address

* **Version**

       /v1.0
  
* **URL**

      /v1.0/eWallet/transaction
       
* **Method:**

       GET
  
* **URL Params**

        None

* **Data Params**

      {
         "from": "0x1a3F275b9Af71D5972198991511*******",
      }

* **Success Response:**

  * **Code:** 200 OK 
    
    **Content:** 

      {
        "status_response": {
            "status_code": 200,
            "description_code": "OK",
            "description": "Transactions fetched"
            },
        "data": [
            {
                "from": "0x1a3F275b9Af71D597219899151140a00***",
                "to": "0xf4f611dCa7DEa0f3aB706B9af03522200aE7***",
                "block_number": 70710**,
                "transaction_hash": "0x195514fea3f4379a87e961b0328e2174e8fca2efe66b6841232ce72370d****"
            },
            {
            "from": "0x1a3F275b9Af71D597219899151140a0049DB***",
            "to": "0xf4f611dCa7DEa0f3aB706B9af03522200aE79***",
            "block_number": 70711**,
            "transaction_hash": "0x1c31beeec70c7b8a432e9ec477c874e17c08cfe680af5cc03db24b6efe***"
            }
            ]
        }
 
* **Error Response:**

  * **Code:** 500 INTERNAL SERVER ERROR 
   
    **Content:** 
        
         {
            "status": {
            "status_code": 500,
            "description_code": "FAILURE",
            "description": "Post http://localhost:9200/ethereum_transactions/_search: dial tcp 127.0.0.1:9200: connect: connection refused"
            }
        }

  OR

  * **Code:** 400 STATUS BAD REQUEST
               
    **Content:** 
    
        {
          "status": {
            "status_code": 400,
            "description_code": "FAILURE",
            "description": "Request binding failed"
          }
        }

