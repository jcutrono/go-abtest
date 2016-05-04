package roll

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Selection struct {
	Option   string
	Selected float32
	Offered  float32
}

type Campaign struct {
	Name    string
	Options []Selection
}

type CampaignDTO struct {
	Name    string
	Options []string
}

var (
	mongourl string
	Database string
)

func write(s Campaign) {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("campaigns")
	err := c.Insert(s)
	if err != nil {
		log.Fatal(err)
	}
}

func updateOffered(name string, option string) {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("campaigns")
	err := c.Update(bson.M{"name": name, "options.option": option}, bson.M{"$inc": bson.M{"options.$.offered": 1}})

	if err != nil {
		log.Fatal(err)
	}
}

func updateSelected(name string, option string) {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("campaigns")

	err := c.Update(bson.M{"name": name, "options.option": option}, bson.M{"$inc": bson.M{"options.$.selected": 1}})

	if err != nil {
		log.Fatal(err)
	}
}

func FindCampaign(name string) Campaign {
	session := GetSession()
	defer session.Close()

	c := session.DB(Database).C("campaigns")

	var result Campaign
	c.Find(bson.M{"name": name}).One(&result)

	return result
}

func GetSession() *mgo.Session {
	session, err := mgo.Dial(mongourl)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func init() {

	if mongourl = os.Getenv("MONGO_URL"); mongourl == "" {
		mongourl = "mongodb://localhost"
	}

	Database = "roller"
}
