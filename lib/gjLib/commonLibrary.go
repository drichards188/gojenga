package gjLib

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

////works
////hashHistory should have ledger items it is changing, hash of said ledger, hash of prev hash and current hash, time

////works

type Traffic struct {
	Name               string
	Amount             string
	SourceAccount      string
	DestinationAccount string
	Verb               string
	Role               string
	Port               string
	Payload            string
	Password           string
}

//type User struct {
//	Account  string
//	Password string
//}

//gjUser := User{
//	Account:  "david",
//	Password: "54321",
//}

//ledger := Ledger{
//	Account: "david",
//	Amount:  "200",
//}

//gjQuery := Query{
//	TableName: "hashHistory",
//	Key:       "Iteration",
//	Value:     "1",
//}

//works

//func HashLedger(data string) (results string, err error) {
//	theHash := ""
//	var iteration int
//	hashResult, err := queryMongoAll("hashHistory")
//	if err != nil {
//		log.Println(err)
//		return "gjQuery error in hashLedger", errors.New("gjQuery error in hashLedger")
//	}
//	if hashResult != nil {
//		iteration = len(hashResult)
//
//		mongoHash, err := queryHash(iteration)
//		if err != nil {
//			log.Println(err)
//			return "gjQuery error in hashLedger", errors.New("gjQuery error in hashLedger")
//		}
//
//		//queryMongo(Account)
//
//		//fmt.Print("Your gjQuery result ")
//		hashMap := mongoHash.Map()
//		if _, ok := hashMap["Message"]; ok {
//			//do something here
//
//			if hashMap["Message"].(string) == "No Match" {
//				theHash = "000000"
//			}
//		} else {
//			theHash = hashMap["Hash"].(string)
//			//logger.Info(theHash)
//		}
//	}
//
//	s := data + theHash
//
//	h := sha1.New()
//
//	h.Write([]byte(s))
//
//	bs := h.Sum(nil)
//
//	currentTime := time.Now()
//	hashTime := currentTime.Format("2006-01-02") + currentTime.Format(" 15:04:05")
//
//	//logger.Info(s)
//	prettyHash := fmt.Sprintf("%x", bs)
//	//logger.Info(prettyHash)
//	//logger.Info(hashTime)
//	logger.Debug(fmt.Sprintf("--> ledger hashed: %s", prettyHash))
//
//	_, err = writeToHashHistory("hashHistory", prettyHash, hashTime, iteration+1, theHash, data)
//	if err != nil {
//		log.Println("write error in hashLedger")
//		return "write error in hashLedger", nil
//	}
//	return "hashLedger succesful", nil
//}

func RunDynamoCreateTable(tableName string) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create table Movies
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Account"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Account"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err = svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table in us-east-2")
}

// Create structs to hold info about new item
type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title"`
	Info  ItemInfo `json:"info"`
}

type Hash struct {
	Iteration    int
	Timestamp    string
	Hash         string
	PreviousHash string
	Ledger       string
}

type User struct {
	Password string
	Account  string
	//Info  ItemInfo `json:"info"`
}

type Ledger struct {
	Account string
	Amount  string
}

type Query struct {
	TableName string
	Key       string
	Value     string
}

func RunDynamoCreateItem[T any](tableName string, item T) (resp map[string]string, err error) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	r := make(map[string]string)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//info := ItemInfo{
	//	Plot:   "Nothing happens at all.",
	//	Rating: 0.0,
	//}

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())

		r["msg"] = "-->Could not create: " + tableName
		r["code"] = "1"
		return r, errors.New("error calling PutItem")

	}

	fmt.Println("Successfully added item")

	r["msg"] = "RunDynamoCreateItem finished"

	return r, nil
}

func RunDynamoGetItem(query Query) (resp map[string]string, err error) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	r := make(map[string]string)

	var result *dynamodb.GetItemOutput

	if query.TableName == "hashHistory" {
		result, err = svc.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String(query.TableName),
			Key: map[string]*dynamodb.AttributeValue{
				query.Key: {
					N: aws.String(query.Value),
				},
			},
		})
		if err != nil {
			log.Fatalf("Got error calling GetItem: %s", err)
		}
	} else if query.TableName == "users" || query.TableName == "ledger" {
		result, err = svc.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String(query.TableName),
			Key: map[string]*dynamodb.AttributeValue{
				query.Key: {
					S: aws.String(query.Value),
				},
			},
		})
		if err != nil {
			log.Fatalf("Got error calling GetItem: %s", err)
		}
	}

	//item := Hash{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &r)

	if result.Item == nil {
		r["msg"] = "-->Could not find: " + query.Value
		r["code"] = "1"
		return r, errors.New("-->Could not find: " + query.Value)
	}

	r["msg"] = "RunDynamoGetItem finished"
	r["code"] = "0"

	return r, nil

	//if err != nil {
	//	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	//}
	//
	//if item.Hash == "" {
	//	fmt.Println("Could not find 'The Big New Movie' (2015)")
	//	return
	//}

	//fmt.Println("Found item")
	//fmt.Println("Year:  ", item.Year)
	//fmt.Println("Title: ", item.Title)
	//fmt.Println("Plot:  ", item.Info.Plot)
	//fmt.Println("Rating:", item.Info.Rating)
}

func RunDynamoDeleteItem(tableName string, value string) (resp map[string]string) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	r := make(map[string]string)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//item := User{
	//	Account: "david",
	//}

	m := make(map[string]any)

	if tableName == "users" || tableName == "ledger" {
		m["Account"] = value
	} else if tableName == "hashHistory" {
		m["Iteration"] = value
	}

	av, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		r["msg"] = "-->Could not gjDelete: " + value
		r["code"] = "1"
		return r
	}

	input := &dynamodb.DeleteItemInput{
		Key:       av,
		TableName: aws.String(tableName),
	}

	_, err = svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Delete complete")
	r["msg"] = "-->Completed gjDelete: " + value
	r["code"] = "0"
	return r
}
