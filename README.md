# crud-dengan-go-mysql
Ini adalah sebuah aplikasi simple dari golang, dimana berfungsi untuk

Create , Read, Edit, Delete data Mysql dengan bahasa Go


##Instalasi



- Buat database ber-nama **animal** dan sebuah tabel di database mysql dengan nama tabel **animal**

```
CREATE TABLE `animal` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text(50) DEFAULT NULL,
  `class` text(200) NOT NULL,
  `legs` text(4) NOT NULL,
PRIMARY KEY (id));

INSERT INTO animal (id,name,class,legs) VALUES 
(1,'lion','mammal','4');

 ENGINE=InnoDB DEFAULT CHARSET=latin1;

```
Alamat API Database = localhost:8080/v1/animal

```

- lanjut ketikan perintah ini di terminal atau di git command

 ```git clone https://github.com/handi233/crud-golang```

- terakhir, masuk ke dalam folder project go jalankan perintah 

 ```go run main.go ```

Note : sebelum menjalankan project 
yaitu dengan membuat database terlebih dahulu di mysql / maria db dengan mengaktifkan apache ,dan mysql di xampp apps

##Sekian
