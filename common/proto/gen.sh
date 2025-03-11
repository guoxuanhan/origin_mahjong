./protoc --go_out=./common ./commonproto/*.proto
./protoc --go_out=./outer ./outerproto/*.proto
./protoc --go_out=./inner ./innerproto/*.proto
