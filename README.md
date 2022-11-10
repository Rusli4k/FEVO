FEVO is REST API with two endpoints:

    - POST: 127.0.0.1:8000/transactions   -upload file from request, parse it and save data to DB
        request must contained key "file" and file format -  .csv
        Content-Type: multipart/form-data

    - GET:  127.0.0.1:8000/transactions?filter_name=xxx - get data from DB and sent response with data in JSON format
        allowed filter_name`s:
            ?transaction_id={[0-9]+}   - get row from DB with inputted transaction ID
            ?terminal_id={[0-9]+}      - get row from DB with inputted terminal ID
            ?terminal_id={[0-9]+}&terminal_id={[0-9]+}&terminal_id={[0-9]+}&terminal_id={[0-9]+} 
                - can be used any times in such format - get 
            ?status={accepted|declined}
            ?payment_type={cash|card}
            ?date_post=fromYYYY-MM-DDtoYYYY-MM-DD
            ?date_post=fromYYYY-MM-DD
            ?date_post=toYYYY-MM-DD
            ?payment_narrative={[a-zA-Z]+} 
                -min length is 4 symbols ( can be changed in const`s)
            
environments variables described in .env file

DB struct described in migration\. file


HOW TO RUN 

make migrate up //- creates table in DB

make migrate down  //- drop created table

make startup // - do migrate up and start program

make start // - just start program ( in case of you made migration before)



DO NOT USE DOCKER. IT DON`T WORK!!!

I have no electricity to finish work :( 

At now I don`t know how to upload project to github with out electricity (((
