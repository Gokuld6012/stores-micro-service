package main

import (
	"context"
	"errors"
	"log"
	"net"

	db "github.com/gokuld6012/product-micro/auth-service/database"
	pb "github.com/gokuld6012/product-micro/auth-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Error definition
// TODO: Need to move the error part new file
var (
	errInRegistration    = errors.New("Failed to register the user")
	errInHashPassword    = errors.New("Failed to hash the password")
	errUserAlreadyExists = errors.New("User already exists")
	errUserNotFound      = errors.New("User not found")
	errPasswordIncorrect = errors.New("Password is incorrect")
)

// HashPassword defines to hash the give password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks the hashed password matches with the string password
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type service struct {
	db *mongo.Database
}

func (s *service) GetTestResponse(ctx context.Context, c *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Message: "This is from auth service",
	}, nil
}

// Register new user
func (s *service) RegisterNewUser(ctx context.Context, c *pb.RegisterUserRequest) (*pb.Response, error) {
	hashedPassword, err := HashPassword(c.Password)
	if err != nil {
		return nil, errInHashPassword
	}
	coll := s.db.Collection("user")
	userData := &pb.RegisterUserRequest{}
	err = coll.FindOne(context.Background(), bson.M{"username": c.Username}).Decode(&userData)
	if err == nil {
		return nil, errUserAlreadyExists
	}
	// need to add data sanitisation
	_, err = coll.InsertOne(context.Background(), bson.M{
		"username":  c.Username,
		"firstname": c.Firstname,
		"lastname":  c.Lastname,
		"password":  hashedPassword,
	})
	if err != nil {
		return nil, errInRegistration
	}
	return &pb.Response{
		StatusCode:      200,
		ResponseMessage: "User Created Successfully",
	}, nil
}

func (s *service) Login(ctx context.Context, c *pb.LoginRequest) (*pb.Response, error) {
	userData := &pb.RegisterUserRequest{}
	err := s.db.Collection("user").FindOne(context.Background(), bson.M{
		"username": c.Username,
	}).Decode(userData)
	if err != nil {
		return nil, errUserNotFound
	}
	if !CheckPasswordHash(c.Password, userData.Password) {
		return nil, errPasswordIncorrect
	}
	return &pb.Response{
		StatusCode:      200,
		ResponseMessage: "User Loggedin successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db, err := db.DBConn()

	if err != nil {
		log.Fatalf("Failed to load the connection")
	}

	s := grpc.NewServer()

	pb.RegisterAuthServer(s, &service{
		db: db,
	})

	log.Println("Running on port", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
