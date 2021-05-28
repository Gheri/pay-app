package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	GetUser(userId string) *User
	GetMerchant(merchantId string) *Merchant
	SaveUser(u User) error
	UpdateUser(u User) error
	UpdateMerchant(m Merchant) error
	SaveMerchant(m Merchant) error
	QueryToGetUsersReachedLimits() ([]string, error)
	QueryToGetUserswithDues() ([]User, error)
}

type MongoDatabase struct {
	users        *mongo.Collection
	merchants    *mongo.Collection
	transactions *mongo.Collection
}

func NewDatabase() Database {
	clientOptions := options.Client().ApplyURI(GetMongoString())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	databaseName := GetDatabaseName()
	users := client.Database(databaseName).Collection(GetUsersCollectionName())
	merchants := client.Database(databaseName).Collection(GetMerchantsCollectionName())
	trasaction := client.Database(databaseName).Collection(GetTransactionsCollectionName())
	_, err = users.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		panic(err)
	}
	_, err = merchants.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		panic(err)
	}
	return &MongoDatabase{users, merchants, trasaction}
}

func (d *MongoDatabase) GetUser(userId string) *User {
	filter := bson.M{"name": userId}
	var user User
	d.users.FindOne(context.TODO(), filter).Decode(&user)
	return &user
}

func (d MongoDatabase) GetMerchant(merchantId string) *Merchant {
	filter := bson.M{"name": merchantId}
	var merchant Merchant
	d.merchants.FindOne(context.TODO(), filter).Decode(&merchant)
	return &merchant
}

func (d *MongoDatabase) SaveUser(u User) error {
	_, err := d.users.InsertOne(context.TODO(), u)
	return err
}

func (d *MongoDatabase) UpdateUser(u User) error {
	filter := bson.M{"name": u.Name}
	update := bson.M{
		"$set": bson.M{
			"dues":           u.Dues,
			"crrCreditLimit": u.CrrCreditLimit,
		},
	}
	_, err := d.users.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *MongoDatabase) UpdateMerchant(m Merchant) error {
	filter := bson.M{"name": m.Name}
	update := bson.M{
		"$set": bson.M{
			"discountEarned": m.DiscountEarned,
		},
	}
	_, err := d.merchants.UpdateOne(context.TODO(), filter, update)
	return err
}

func (d *MongoDatabase) SaveMerchant(m Merchant) error {
	_, err := d.merchants.InsertOne(context.TODO(), m)
	return err
}

func (d *MongoDatabase) QueryToGetUsersReachedLimits() ([]string, error) {
	filter := bson.M{
		"crrCreditLimit": bson.M{"$eq": 0},
	}
	cursor, err := d.users.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	userNames := []string{}
	for cursor.Next(context.TODO()) {
		var user User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		userNames = append(userNames, user.Name)
	}
	return userNames, nil
}

func (d *MongoDatabase) QueryToGetUserswithDues() ([]User, error) {
	filter := bson.M{
		"dues": bson.M{"$ne": 0},
	}
	cursor, err := d.users.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	users := []User{}
	for cursor.Next(context.TODO()) {
		var user User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
