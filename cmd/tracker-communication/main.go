package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	trackservice "github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/services"
	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/acls"
	triprepositories "github.com/NicolasDeveloper/tracker-microservices/internal/register/repositories"
	registerrepositories "github.com/NicolasDeveloper/tracker-microservices/internal/trip/repositories"
	registermodels "github.com/NicolasDeveloper/tracker-microservices/internal/trip/models"
	tripservice "github.com/NicolasDeveloper/tracker-microservices/internal/trip/services"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/command"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/report"
	sharedrepositories "github.com/NicolasDeveloper/tracker-microservices/pkg/repositories"
	sharedmodels "github.com/NicolasDeveloper/tracker-microservices/pkg/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
)

func handleConnection(conn net.Conn, ctx dbcontext.DbContext) {
	buf := make([]byte, 1024)
	lenn, err := conn.Read(buf)

	if err != nil {
		log.Println("client left..")
		conn.Close()
		return
	}

	bufferpack := buf[:lenn]
	parser := command.NewParser()
	timeConvert := timeconvert.NewTimeConvert(timeconvert.NewUnixDate())
	mapboxACL := acls.NewMapBoxACL()

	deviceRepository, err := sharedrepositories.NewDeviceRepository(ctx)
	tripRepository, err := triprepositories.NewTripRepository(ctx)
	tripService = tripservice.NewTripService(mapboxACL, tripRepository)
	trackService := trackservice.NewTrackService(deviceRepository)

	deviceID, err := trackService.GetDeviceID(bufferpack)
	device, err := deviceRepository.GetActiveDevice(deviceID)

	if device.ID == "" {
		log.Println("device not found..")
		handleConnection(conn, ctx)
		return
	}

	if parser.IsLogin(bufferpack) == true {
		if resp, err := report.RetriveLogin(bufferpack, timeConvert, parser); err == nil {
			conn.Write(resp)
		}
	} else if parser.IsAlarmReport(bufferpack) == true {
		model, err := trackService.ParseAlarmReport(bufferpack)

		if err == nil {
			size := len(bufferpack) - 1
			buffersliced := bufferpack[27:size]
			model.UserID = device.UserID
			model.DeviceID = device.ID

			if result, err := trackService.IsIgnitionOnAlarm(buffersliced); result == true && err == nil {
				if device.Logged == false {
					queryResp, err := querySearchVinCode(bufferpack)
					if err == nil {
						conn.Write(queryResp)
					}
				}

				err = tripService.Start(model)
			}

			if result, err := trackService.IsIgnitionOffAlarm(buffersliced); result == true && err == nil {
				err = tripService.Close(model)
			}
		}
	} else if parser.IsGpsReport(bufferpack) == true {
		model, err := trackService.ParseGpsReport(bufferpack)
		model.UserID = device.UserID
		model.DeviceID = device.ID

		if err == nil {
			err = tripService.Increment(model)
		}
	} else if parser.IsQueryReport(bufferpack) == true {
		if device.Logged == false {
			tlvDescription, err := trackService.ParseQueryReport(bufferpack)
			if tlvDescription.VinCode != "" && err == nil {
				_, err = trackService.DoLogin(deviceID)
			}
		}
	}

	// recursive func to handle io.EOF for random disconnects
	handleConnection(conn, ctx)
}

func querySearchVinCode(bufferpack []byte) ([]byte, error) {
	lengthHex := "0x2800"
	sequenceHex := "0x0200"
	queryNumberHex := "0x03"
	parametersHex := "0x012001160115"
	return report.QueryCommand(bufferpack, lengthHex, sequenceHex, queryNumberHex, parametersHex, command.NewParser())
}

func main() {
	hostName := os.Getenv("HOSTNAME")

	if hostName == "" {
		hostName = "localhost"
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5005"
	}

	ln, err := net.Listen("tcp", hostName+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port: 5005")

	defer ln.Close()
	rand.Seed(time.Now().Unix())

	ctx := dbcontext.NewContext()
	ctx.Connect()
	createDatabaseIndex(ctx)
	initilizeDatabse(ctx)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Calling handleConnection")
		go handleConnection(conn, ctx)
	}
}

func initilizeDatabse(ctx dbcontext.DbContext) {
	userRepository, err := registerrepositories.NewUserRepository(ctx)
	deviceRepository, err := sharedrepositories.NewDeviceRepository(ctx)

	device, err := deviceRepository.GetActiveDevice("213GDP2018022421")

	if device.ID != "" {
		return
	}

	user, err := registermodels.NewUser()

	if err == nil {
		userRepository.Create(user)
	}

	device, err = sharedmodels.NewDevice("213GDP2018022421", user.ID)

	if err != nil {
		return
	}

	deviceRepository.CreateDevice(device)
}

func createDatabaseIndex(ctx dbcontext.DbContext) {
	deviceCollection, err := ctx.GetCollection(trackmodels.Device{})

	if err != nil {
		return
	}

	deviceCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"pin_code": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
}
