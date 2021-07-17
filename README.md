

# SnT summer project
 IITK-coin 

Backend for a centralized coin exchange system, developed under the mentorship of P-club IIT Kanpur

To run the project install go in your system and add it to your path:
https://golang.org/doc/install

run the command "$ go install" to install the dependencies used in the project

run the command "$ go run main.go" to start the local server at port 8080.


<h3>For user auth</h3>                    
 POST  req at  "http://localhost:8080/signin" for signin

<br>
 POST  req at  "http://localhost:8080/signup"

<br>
 GET  req at  "http://localhost:8080/secretpage" <br>
Add JWT token in the header of req...

<br>
 <h3>For currency transactions :-  </h3>                 
 POST  req at  "http://localhost:8080/init"<br>
To award an amount to a roll number if the roll no had an account<br>
Else if no account is present then an account is created in the DB and amount is added

 POST  req at  "http://localhost:8080/transfer"<br>
To transfer an amount from one roll number to another if the transaction is valid<br>

 GET  req at  "http://localhost:8080/balance"<br>
          
                  
