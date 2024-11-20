## mysql


### 外部からのアクセスを許可

```sh
mysql -u root -p
```

```sh
CREATE USER isuconp@'%' IDENTIFIED BY 'isuconp';
GRANT ALL on *.* TO isuconp@'%';
CREATE USER isuconp@'localhost' IDENTIFIED BY 'isuconp';
GRANT ALL on *.* TO isuconp@'localhost';
```
