# Query Optimization

1. Wajib menggunakan index pada database. Contoh pada tabel users kolom email dan phone adalah uniq serta untuk kolom name dipasang index.

2. Hindari penggunaan full dan prefix wildcard. usahakan menggunakan postfix wildcard.

3. Gunakanlah limit pada query.

4. Hindari penggunaan * . Pilih kolom yang akan digunakan saja.

Berikut perbedaan response ketika menggunakan index dan tidak dengan bahasa pemrograman php (flight php):

```
select * from users a where email='m.rheza69@gmail.com';
no indexing -> 1417 ms
indexing -> 16 ms

select * from users a where phone='08220000008'
no indexing -> 1478 ms
indexing -> 42 ms

select email from users a where name like 'rhe%'
no indexing -> 1535 ms
indexing -> 78 ms
```

Berikut perbedaan response ketika menggunakan index dan tidak dengan bahasa pemrograman go :
```
select email from users a where name like 'rhe%'
no indexing -> 1655 ms
indexing -> 32 ms
```

untuk mencoba app silahkan extract users.rar kemudian import file users.json yang berisi 1 juta record.

api php ada didalam folder flight dan api go ada didalam folder rest-api.
