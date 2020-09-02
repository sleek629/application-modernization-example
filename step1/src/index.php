<html>
<head><title>Word Database</title></head>

<body>
<h1>Word Database</h1>
Input something
<form method="POST" action="<?php print($_SERVER['PHP_SELF']) ?>">
<input type="text" name="word">
<input type="submit" value="submit">
</form>

<?php
echo "Your input: " . htmlspecialchars($line['word'], ENT_QUOTES, 'UTF-8') . "\n";
$user = "user";
$password = "Password@123";

$dbh = new PDO("mysql:host=localhost; dbname=word_db; charset=utf8", "$user", "$password");

if ($_POST['word']) {
  $word = $_POST['word'];
  $stmt = $dbh->prepare('SELECT word FROM word_tb WHERE word = :word');
  $stmt->bindValue(':word', $word, PDO::PARAM_STR);
  $stmt->execute();
  if (!$stmt->rowCount()){
    $stmt = $dbh->prepare("INSERT INTO word_tb (word, count) VALUES (:word, :count)");
    $stmt->bindValue(':word', $word, PDO::PARAM_STR);
    $stmt->bindValue(':count', 1, PDO::PARAM_INT);
    $stmt->execute();
  } else {
    $stmt = $dbh->prepare("UPDATE word_tb SET count = count + 1 WHERE word = :word");
    $stmt->bindValue(':word', $word, PDO::PARAM_STR);
    $stmt->execute();
  }
}
?>

<table border='1'>
<tr><th>word</th><th>count</th></tr>
<?php
$stmt = $dbh->prepare('SELECT word, count FROM word_tb');
$stmt->execute();
$lines = $stmt->fetchAll();
foreach($lines as $line) {
  echo '<tr><td>' . htmlspecialchars($line['word'], ENT_QUOTES, 'UTF-8') . '</td><td>' . $line['count'] . "</td></tr>\n";
}
?>
</table>

</body>
</html>
