<?php
require 'flight/Flight.php';

Flight::register('db', 'PDO', array('mysql:host=localhost;dbname=test', 'root', 'root' ));
$db = Flight::db();

Flight::route('GET /get_email', function () {
	$sql = "select * from users a where email='m.rheza69@gmail.com'";
	$res = Flight::db()->query($sql);
	$records = $res->fetchAll(PDO::FETCH_ASSOC);

	echo json_encode($records);
});

Flight::route('GET /get_phone', function () {
	$sql = "select * from users a where phone='08220000008'";
	$res = Flight::db()->query($sql);
	$records = $res->fetchAll(PDO::FETCH_ASSOC);

	echo json_encode($records);
});

Flight::route('GET /get_name', function () {
	$sql = "select email from users a where name like 'rhe%'";
	$res = Flight::db()->query($sql);
	$records = $res->fetchAll(PDO::FETCH_ASSOC);

	echo json_encode($records);
});

Flight::start();
