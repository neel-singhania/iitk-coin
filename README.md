

# SnT summer project
 IITK-coin 

Backend for a centralized coin exchange system, developed under the mentorship of P-club IIT Kanpur

 <h2>endpoins</h2>
<h3>For user auth</h3>                    
 POST  req at  "http://localhost:8080/signin" for signin
<h5>Example req </h5>
{
  "roll_no" : "195380" , <br>
  "password" : "password" <br>
} 
<br>
 POST  req at  "http://localhost:8080/signup"
<h5>Example req </h5>
{
  "roll_no" : "190460" , <br>
  "password" : "pass" <br>
} 
<br>
 GET  req at  "http://localhost:8080/secretpage" <br>
Add JWT token in the header of req...
<h5>Example req </h5>
{ <br>
  "Authorization" : "token_string"  <br>
} <br> 

<br>
 <h3>For currency transactions :-  </h3>                 
 POST  req at  "http://localhost:8080/init"<br>
To award an amount to a roll number if the roll no had an account<br>
Else if no account is present then an account is created in the DB and amount is added
<h5>Example req </h5>
{ <br>
  "roll_no" : "190460" , <br>
  "balance" : 1000 <br>
} <br> 
<br>
 POST  req at  "http://localhost:8080/transfer"<br>
To transfer an amount from one roll number to another if the transaction is valid<br>
<h5>Example req </h5>

{ <br>
    "from_roll_no" : "190981" , <br>
    "to_roll_no" : "190982", <br>
    "amount" : 1000 <br>
} <br>
<br>
 GET  req at  "http://localhost:8080/balance"<br>

<h5>Example req </h5>
{ <br>
  "roll_no" : "190460"  <br>
  
} <br>                  
                  
