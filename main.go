package main
// // This program provides a sample application for using MongoDB with
// // the mgo driver.
// package main

// import (
// 	"log"
// 	"sync"
// 	"time"

// 	"labix.org/v2/mgo"
// )

// const (
// 	MongoDBHosts = "127.0.0.1"
// 	AuthDatabase = "crypt"
// 	AuthUserName = ""
// 	AuthPassword = ""
// 	TestDatabase = "crypt"
// )

// // type (
// // 	processedCurrency struct {
// // 		Name           string        `bson:"name"`
// // 		Base           string        `bson:"base"`
// // 		BaseCurrency   string        `bson:"baseCurrency"`
// // 		MarketCurrency string        `bson:"marketCurrency"`
// // 		BuyOrders      []interface{} `bson:"buyOrders"`
// // 		SellOrders     []interface{} `bson:"sellOrders"`
// // 		History        []interface{} `bson:"history"`
// // 		Volume         float32       `bson:"volume"`
// // 		BaseVolume     float32       `bson:"baseVolume"`
// // 		Time           string        `bson:"time"`
// // 		Price          float32       `bson:"price"`
// // 	}
// // )

// // main is the entry point for the application.
// func main() {
// 	// We need this object to establish a session to our MongoDB.
// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:    []string{MongoDBHosts},
// 		Timeout:  60 * time.Second,
// 		Database: AuthDatabase,
// 		Username: AuthUserName,
// 		Password: AuthPassword,
// 	}

// 	// Create a session which maintains a pool of socket connections
// 	// to our MongoDB.
// 	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		log.Fatalf("CreateSession: %s\n", err)
// 	}

// 	// Reads may not be entirely up-to-date, but they will always see the
// 	// history of changes moving forward, the data read will be consistent
// 	// across sequential queries in the same session, and modifications made
// 	// within the session will be observed in following queries (read-your-writes).
// 	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
// 	mongoSession.SetMode(mgo.Monotonic, true)

// 	for i := 0; i < 15; i++ {
// 		// Create a wait group to manage the goroutines.
// 		var waitGroup sync.WaitGroup
// 		for result := range bitrex.GetCurrencies() {
// 			waitGroup.Add(1)
// 			go RunQuery(result, &waitGroup, mongoSession)

// 		}

// 		// Wait for all the queries to complete.
// 		waitGroup.Wait()
// 		//close(out)
// 		log.Println("All Single Queries Completed")
// 		//time.Sleep(5 * time.Second)
// 	}
// 	log.Println("All Queries Completed")

// }

// // RunQuery is a function that is launched as a goroutine to perform
// // the MongoDB work.
// func RunQuery(query bitrex.ProcessedCurrency, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
// 	// Decrement the wait group count so the program knows this
// 	// has been completed once the goroutine exits.
// 	defer waitGroup.Done()

// 	// Request a socket connection from the session to process our query.
// 	// Close the session when the goroutine exits and put the connection back
// 	// into the pool.
// 	sessionCopy := mongoSession.Copy()
// 	defer sessionCopy.Close()

// 	// Get a collection to execute the query against.
// 	collection := sessionCopy.DB(TestDatabase).C("crypt")

// 	if err := collection.Insert(query); err != nil {
// 		print(err)
// 	}

// }
