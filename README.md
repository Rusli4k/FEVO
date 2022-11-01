FEVO is REST API with two endpoints:

    - POST: 127.0.0.1:8000/transactions   -upload file from request, parse it and save data to DB
        request must contained key "file" and file format -  .csv
        Content-Type: multipart/form-data

    - GET:  127.0.0.1:8000/transactions?filter=xxx - get data from DB and sent response with data in JSON format
        allowed filters:
            ?transaction_id={[0-9]+}
            ?terminal_id={[0-9]+}
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
make start // - just start program ( in case of made migration before)

