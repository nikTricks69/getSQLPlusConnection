# getSQLPlusConnection
GSD - have a scenario where the DB connection is available in a kubernetes secret (basically a property file) , need to pull it out and create a connectionstring
and corporate windows lappy isnt helping


Example

```
database.host=oracle-gsd.com
database.port=1521
database.service=server1
database.username=userName
database.password=superPassword
```
Usage

```
ps

$load = kubectl get secret $secretName -n $ns -o jsonpath='{.data}' | jq .[] ) >$null
getSqlPlusConnString $load

```
