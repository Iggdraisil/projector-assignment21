# projector-assignment21

## To set up cluser run:
 - `docker-compose up --build -d`
 - `docker cp master.sql assignment21-mysql-master-1:/`
 - `docker cp mydb.sql assignment21-mysql-master-1:/`
 - `docker exec assignment21-mysql-master-1`
 - `mysql -p < master.sql`
 - `mysql -p < mysb.sql`
 #
    Then run in mysql `show master status` and record binlog file and position
    Edit slave1.sql and slave2.sql with recorded values
 
 ### Copy slaves scripts:
 - `docker cp slave1.sql assignment21-mysql-slave1-1:/`
 - `docker cp slave2.sql assignment21-mysql-slave2-1:/`
 
 ### First slave:
 - `docker exec assignment21-mysql-slave1-1`
 - `mysql -p < slave1.sql`
 - `echo 'show slave status\G;' | mysql -p`    #Verify slave status
 
 ### Second slave:
 - `docker exec assignment21-mysql-slave2-1`
 - `mysql -p < slave2.sql`
 - `echo 'show slave status\G;' | mysql -p`    #Verify slave status
 
 ### Run inserter to insert data into master db
 
 
 ### Output from slave status:
 ![image](https://user-images.githubusercontent.com/44341837/236875825-8da3803e-6ba4-4c13-925d-ddc3242f8890.png)

 
 ## Killing slave:
   1. Ensure inserter is running 
   2. To kill slave run (SQL on replica) `STOP REPLICA;`;
   3. Wait some time 
   4 Run `START REPLICA;` and immidiately `show slave status\G`;
   
 ![image](https://user-images.githubusercontent.com/44341837/236877855-18623ada-e248-4719-823a-9eaf288c3ac0.png)

## Removal of column:
   - Query OK, data now is incosisstent


