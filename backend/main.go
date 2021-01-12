package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team19/app/controllers"
	"github.com/team19/app/ent"
)

// Departments defines the struct for the departments
type Departments struct {
	Department []Department
}



// Department defines the struct for the department
type Department struct {
	DEPARTMENT string
	FACULTY    string
}



// InstructorRooms defines the struct for the instructorrooms
type InstructorRooms struct {
	InstructorRoom []InstructorRoom
}

// InstructorRoom defines the struct for the instructorroom
type InstructorRoom struct {
	ROOM     string
	BUILDING string
}

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewInstructorInfoController(v1, client)
	controllers.NewInstructorRoomController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewTitleController(v1, client)

	// Set Departments Data
	departments := Departments{
		Department: []Department{
			Department{"MANUFACTURING ENGINEERING", "Institute of Engineering"},
			Department{"AGRICULTURAL ENGINEERING", "Institute of Engineering"},
			Department{"TRANSPORTATION ENGINEERING", "Institute of Engineering"},
			Department{"COMPUTER ENGINEERING", "Institute of Engineering"},
			Department{"CHEMICAL ENGINEERING", "Institute of Engineering"},
			Department{"MECHANICAL ENGINEERING", "Institute of Engineering"},
			Department{"CERAMIC ENGINEERING", "Institute of Engineering"},
			Department{"TELECOMMUNICATION ENGINEERING", "Institute of Engineering"},
			Department{"POLYMER ENGINEERING", "Institute of Engineering"},
			Department{"ELECTRICAL ENGINEERING", "Institute of Engineering"},
			Department{"CIVIL ENGINEERING", "Institute of Engineering"},
			Department{"METALLURGICAL ENGINEERING", "Institute of Engineering"},
			Department{"ENVIRONMENTAL ENGINEERING", "Institute of Engineering"},
			Department{"ELECTRONIC ENGINEERING", "Institute of Engineering"},
			Department{"INDUSTRIAL ENGINEERING", "Institute of Engineering"},
			Department{"AUTOMOTIVE ENGINEERING", "Institute of Engineering"},
			Department{"MECHATRONICS", "Institute of Engineering"},
			Department{"AERONAUTICAL ENGINEERING", "Institute of Engineering"},
			Department{"COMMUNICATION ARTS", "Institute of Social Technology"},
			Department{"GENERAL", "Doctor of Medicine Program"},
			Department{"NURSING SCIENCE", "Faculty of Nursing"},
		},
	}
	for _, d := range departments.Department {
		client.Department.
			Create().
			SetDEPARTMENT(d.DEPARTMENT).
			SetFACULTY(d.FACULTY).
			Save(context.Background())
	}

	// Set InstructorRooms Data
	instructorrooms := InstructorRooms{
		InstructorRoom: []InstructorRoom{
			InstructorRoom{"A101", "Building A"},
			InstructorRoom{"A102", "Building A"},
			InstructorRoom{"A103", "Building A"},
			InstructorRoom{"A104", "Building A"},
			InstructorRoom{"A105", "Building A"},
			InstructorRoom{"B101", "Building B"},
			InstructorRoom{"B102", "Building B"},
			InstructorRoom{"B103", "Building B"},
			InstructorRoom{"B104", "Building B"},
			InstructorRoom{"B105", "Building B"},
		},
	}
	for _, ir := range instructorrooms.InstructorRoom {
		client.InstructorRoom.
			Create().
			SetROOM(ir.ROOM).
			SetBUILDING(ir.BUILDING).
			Save(context.Background())
	}

	// Set Titles Data
	titles := []string{"Mr.", "Mrs.", "Ms.", "Dr.", "Asst.Prof.", "Assoc.Prof.", "Prof."}
	for _, t := range titles {
		client.Title.
			Create().
			SetTITLE(t).
			Save(context.Background())
	}

	// Set Subjects Data
	subjects := []string{"We love environment", "We love environment",
		"Love the world", "Health and Care"}
	for _, sub := range subjects {
		client.Subject.
			Create().
			SetSubjectName(sub).
			Save(context.Background())
	}

	// Set Degrees Data
	degrees := []string{"We love environment", "We love environment",
		"Love the world", "Health and Care"}
	for _, deg := range degrees {
		client.Degree.
			Create().
			SetDegreeName(deg).
			Save(context.Background())
	}


	



	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
