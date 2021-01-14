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

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
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
	controllers.NewTermController(v1, client)
	controllers.NewYearController(v1, client)
	controllers.NewSubjectController(v1, client)
	controllers.NewDegreeController(v1, client)
	controllers.NewSubjectsOfferedController(v1, client)

	// Set Departments Data
	departments := Departments{
		Department: []Department{
			{"MANUFACTURING ENGINEERING", "Institute of Engineering"},
			{"AGRICULTURAL ENGINEERING", "Institute of Engineering"},
			{"TRANSPORTATION ENGINEERING", "Institute of Engineering"},
			{"COMPUTER ENGINEERING", "Institute of Engineering"},
			{"CHEMICAL ENGINEERING", "Institute of Engineering"},
			{"MECHANICAL ENGINEERING", "Institute of Engineering"},
			{"CERAMIC ENGINEERING", "Institute of Engineering"},
			{"TELECOMMUNICATION ENGINEERING", "Institute of Engineering"},
			{"POLYMER ENGINEERING", "Institute of Engineering"},
			{"ELECTRICAL ENGINEERING", "Institute of Engineering"},
			{"CIVIL ENGINEERING", "Institute of Engineering"},
			{"METALLURGICAL ENGINEERING", "Institute of Engineering"},
			{"ENVIRONMENTAL ENGINEERING", "Institute of Engineering"},
			{"ELECTRONIC ENGINEERING", "Institute of Engineering"},
			{"INDUSTRIAL ENGINEERING", "Institute of Engineering"},
			{"AUTOMOTIVE ENGINEERING", "Institute of Engineering"},
			{"MECHATRONICS", "Institute of Engineering"},
			{"AERONAUTICAL ENGINEERING", "Institute of Engineering"},
			{"COMMUNICATION ARTS", "Institute of Social Technology"},
			{"GENERAL", "Doctor of Medicine Program"},
			{"NURSING SCIENCE", "Faculty of Nursing"},
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
			{"A101", "Building A"},
			{"A102", "Building A"},
			{"A103", "Building A"},
			{"A104", "Building A"},
			{"A105", "Building A"},
			{"B101", "Building B"},
			{"B102", "Building B"},
			{"B103", "Building B"},
			{"B104", "Building B"},
			{"B105", "Building B"},
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
	subjects := []string{"102105 ORGANIC CHEMISTRY",
	"103102  CALCULUS ",
	"109209 MEDICAL BIOCHEMISTRY",
	"116200  EMBRYOLOGY",
	"104110 CELL BIOLOGY",
	"116203 MEDICAL IMMUNOLOGY",
	"202241 LAW IN DAILY LIFE",
	"214302 COMPUTER SERVICES",
	"109302 NUTRITION AND METABOLISM",
	"116208 NERVOUS SYSTEM",
	"116303 DIGESTIVE SYSTEM",
	"116301 RESPIRATORY SYSTEM",
	"116308 MEDICAL TOXICOLOGY",
	"116307 HEMATOPOIETIC",
	"202108 DIGITAL LITERACY",
	"202201  LIFE SKILLS",
	"523301 COMPUTER STATISTICS",
	"523101 COMPUTER PROGRAMMING",
	"525101  ENGINEERING GRAPHICS",
	"523211 DATABASE SYSTEMS",
	"523332 SOFTWARE ENGINEERING",
	"213101 ENGLISH FOR COMMUNICATION"}
	for _, sub := range subjects {
		client.Subject.
			Create().
			SetSubjectName(sub).
			Save(context.Background())
	}

	// Set Degrees Data
	degrees := []string{"ปริญญาตรี",
	"ปริญญาโท",
	"ปริญญาเอก",
	"ปริญญาตรีเพิ่มขึ้นอีกสาขาหนึ่ง",
	"ผู้ร่วมเรียนปริญญาตรี",
	"ผู้ร่วมเรียนปริญญาโท",
	"นักศึกษาหลักสูตรนานาชาติ"}
	for _, deg := range degrees {
		client.Degree.
			Create().
			SetDegreeName(deg).
			Save(context.Background())
	}

	// Set Year Data
	years := []int{2555,2556,2557,2558,2559,2560,2561,2562,2563,2564,2565,2566,2567,2568,2569,2570,2571,2572,2573,2574,2575}
	for _, y := range years {
		client.Year.
			Create().
			SetYEAR(y).
			Save(context.Background())
	}

	// Set Term Data
	terms := []int{1,2,3}
	for _, te := range terms {
		client.Term.
			Create().
			SetTERM(te).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
