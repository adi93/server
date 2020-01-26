sqlFile=$1
dbName="./task.db"
echo "sqlFile : $sqlFile"

if [ -z "$1" ]; then
	echo "Require a sql file"
	exit 1
fi
## ensure that file exists
if ! [ -e $dbName ]
then
	touch $dbName
fi

## ensure that table exists
x=$(sqlite3 ${dbName} ".table migration")
if [ -z $x ]
then
	echo "Migration table not found, adding it..."
	sqlite3 $dbName "create table migration (fileName text, lastLineNumber integer); "
	sqlite3 $dbName "insert into migration values ('$sqlFile', 0)"
	echo "Added migration table"
fi

## run migration
r=$(sqlite3 $dbName "select lastLineNumber from migration where fileName = '$sqlFile'")
echo "lastlineNumber: $r"
$(cat $sqlFile | tail -n +$r | sqlite3 $dbName)
if [ "$?" -gt "0" ]; then
	echo "Error running sql queries"
	echo "aborting..."
	exit 1
fi

lastLine=`wc -l $sqlFile | awk '{print $1}'`
echo "new lastlineNumber: $lastLine"
`sqlite3 ${dbName} "update migration set lastLineNumber = $lastLine where fileName = '$sqlFile'"`
