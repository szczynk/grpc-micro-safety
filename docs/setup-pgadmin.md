# Setup pgAdmin Servers List

By default, postgres can only be access using pgAdmin. The pgadmin's servers list is hard to automatically generated using Dockerfile and i give up on that. Use this following instructions to setup pgadmin's servers list

## 1. Open the following [pgAdmin UI](http://localhost:5050) in the browser and insert the credential

```bash
Username: pgadmin4@pgadmin.org
Password: admin
```

## 2. Click `Tools`

![pgAdmin-1.png](/assets/pgadmin/pgAdmin-1.png)

## 3. Click `Import/Export Servers`

![pgAdmin-2.png](/assets/pgadmin/pgAdmin-2.png)

## 4. Select Import and Click the folder icon

![pgAdmin-3.png](/assets/pgadmin/pgAdmin-3.png)

## 5. Click the three dot icon

![pgAdmin-4.png](/assets/pgadmin/pgAdmin-4.png)

## 6. Click `Upload`

![pgAdmin-5.png](/assets/pgadmin/pgAdmin-5.png)

## 7. Drop files or Click to select files and Select **servers.json** on <current_directory>/grpc-micro-safety/pgadmin4

### example: on my machine it's /home/szczynk/GoDev/grpc-micro-safety/pgadmin4

![pgAdmin-6.png](/assets/pgadmin/pgAdmin-6.png)

## 8. Close the upload section

![pgAdmin-7.png](/assets/pgadmin/pgAdmin-7.png)

## 9. Select **servers.json** and Click `Select`

![pgAdmin-8.png](/assets/pgadmin/pgAdmin-8.png)

## 10. Filename should be **/servers.json** and Click `Next`

![pgAdmin-9.png](/assets/pgadmin/pgAdmin-9.png)

## 11. Checklist the **servers**

![pgAdmin-10.png](/assets/pgadmin/pgAdmin-10.png)

## 12. Click `Next`

![pgAdmin-11.png](/assets/pgadmin/pgAdmin-11.png)

## 13. Click `Finish`

![pgAdmin-12.png](/assets/pgadmin/pgAdmin-12.png)

## 14. Click `OK` on the notification

![pgAdmin-13.png](/assets/pgadmin/pgAdmin-13.png)

## 15. Select one of listed database

![pgAdmin-14.png](/assets/pgadmin/pgAdmin-14.png)

## 16. Insert the database's password

### example: if development env with default setting, insert "postgres" as the password

![pgAdmin-15.png](/assets/pgadmin/pgAdmin-15.png)

## 17. Repeat 15 and 16 until all of listed databases is connected

## 18. Done
