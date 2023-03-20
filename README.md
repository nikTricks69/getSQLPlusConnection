# getSQLPlusConnection
GSD - have a scenario where the DB connection in available in a kubernetes secret , need to pull them out and create a connectionstring

Example
```
database.host=oracle-gsd.com
database.port=1521
database.service=server1
database.username=userName
database.password=superPassword
```
Usage

'''
ps

$load = kubectl get secret $secretName -n $ns -o jsonpath='{.data}' | jq .[] ) >$null
getSqlPlusConnString $load

'''