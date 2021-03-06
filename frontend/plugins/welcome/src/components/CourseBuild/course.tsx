import React, { FC, useEffect } from 'react';


import Swal from 'sweetalert2'; // alert

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import {
 Content,
 Page,
 pageTheme,

} from '@backstage/core';


import MenuBookIcon from '@material-ui/icons/MenuBook';
import { EntSubject } from '../../api/models/EntSubject'; //import ข้อมูลจาก backend
import { EntDegree } from '../../api/models/EntDegree'; 
import { EntDepartment } from '../../api/models/EntDepartment'; 
import { DefaultApi} from '../../api/apis';

import {

  
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  AppBar,
  MuiThemeProvider,
  createMuiTheme,
  Toolbar,

} from '@material-ui/core';


const useStyles = makeStyles((theme: Theme) =>

  createStyles({
    root: {
      '& .MuiTextField-root': {
        marginLeft: theme.spacing(9),
        width: '30ch',
      },
    },

    
    text0: {  
      marginLeft: theme.spacing(-130),
      marginRight: theme.spacing(4),
      marginTop: theme.spacing(3),
      fontSize: 20,
    },

    text1: {  //ชื่อชมรม
        marginLeft: theme.spacing(-130),
        marginRight: theme.spacing(4),
        marginTop: theme.spacing(8),
        fontSize: 20,
    },
    text2: {
      marginLeft: theme.spacing(4),
      marginTop: theme.spacing(4),
      fontSize: 20,
  },
  text3: {
    marginLeft: theme.spacing(-130),
    marginTop: theme.spacing(4),
    fontSize: 20,
},

text4: {
  marginLeft: theme.spacing(3), 
  marginRight: theme.spacing(124), 
  fontSize: 30,
  color: "white",
},






text5: {
  marginLeft: theme.spacing(2),
  marginTop: theme.spacing(1),
  marginBottom:  theme.spacing(1),
  fontSize: 14,
  color: "black",
},

text6: {
  marginLeft: theme.spacing(0),
  marginTop: theme.spacing(1),
  marginBottom:  theme.spacing(1),
  fontSize: 14,
  color: "white",
},

text7: {
  marginLeft: theme.spacing(10),
  marginTop: theme.spacing(0),
  color: "white",
},



text8: {
  marginLeft: theme.spacing(-130),
  marginTop: theme.spacing(4),
  marginBottom:  theme.spacing(1),
  fontSize: 21,
  color: "white",
},

text9: {
  marginLeft: theme.spacing(-130),
  marginTop: theme.spacing(4),
  marginBottom:  theme.spacing(1),
  fontSize: 21,
  color: "white",
},

text10: {
  marginLeft: theme.spacing(-130),
  marginTop: theme.spacing(4),
  marginBottom:  theme.spacing(1),
  fontSize: 21,
  color: "white",
},

  box1: {
    marginLeft: theme.spacing(15),
    marginTop: theme.spacing(2),
    width: '60ch',  
},

  box2: {
    marginLeft: theme.spacing(11.5),
    marginTop: theme.spacing(2),
    width: '60ch',  
  },

  box3: {
    marginLeft: theme.spacing(13),
    marginTop: theme.spacing(5),
    width: '60ch',  
},


  textfill1: {  //ชื่อกิจกรรม
    marginLeft: theme.spacing(14),
    marginTop: theme.spacing(2),
    width: '60ch',
  },

  textfill2: {  //ชื่อกิจกรรม
    marginLeft: theme.spacing(14),
  marginTop: theme.spacing(2),
  width: '60ch',
  },
  
  textfill3: {  //ชื่อกิจกรรม
    marginLeft: theme.spacing(16),
    marginTop: theme.spacing(2),
    width: '60ch',
  },
  
  
  button: {
    marginTop: theme.spacing(8),
    marginLeft: theme.spacing(-100),
    display: 'flex',
    flexWrap: 'wrap',
    
  },
  }),
);

const theme = createMuiTheme({
  palette: {
    primary: {
      // light: will be calculated from palette.primary.main,
      main: "#ef6694"
      // dark: will be calculated from palette.primary.main,
      // contrastText: will be calculated to contrast with palette.primary.main
    }
  }
});

//โครงสร้างของ interface
interface currency{
    DepartmentID :      number;
	  SubjectID :         number;
	  DegreeID  :         number;
    CourseName :        string;
    CourseYear :        number;
    TeacherID :        string;
	
}

const Course: FC<{}> = () => {
const classes = useStyles();



const api = new DefaultApi();

 const [currency, setCurrency] = React.useState<Partial<currency>>({});
 const [departments, setDepartments ] = React.useState<EntDepartment[]>([]);
 const [subjects, setSubjects ] = React.useState<EntSubject[]>([]);
 const [degrees, setDegrees ] = React.useState<EntDegree[]>([]);
 



 const Toast = Swal.mixin({  //ตั้งค่าการแจ้งเตือน
  toast: true,
  position: 'center',
  showConfirmButton: false,
  timer: 5000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

const handleChange = ( event: React.ChangeEvent<{ name?: string; value: any }>,) =>{
    const name = event.target.name as keyof typeof Course;
    const { value } = event.target;
    setCurrency({ ...currency, [name] : value });
    console.log(currency);
};

const handleChangeInt = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
  const name = event.target.name as keyof typeof Course;
  const { value } = event.target;
  setCurrency({ ...currency, [name]: Number(value)});
  console.log(currency);
};


 const getDepartments = async () => {
  const res = await api.listDepartment({ limit: 100, offset:0});
  setDepartments(res);
 };

 const getSubjects = async () => {//ดึงข้อมูล
  const res = await api.listSubject({ limit: 100, offset:0}); //ดึงจาก listUser
  setSubjects(res); //ส่งค่าไป user ข้างบน
 };

 const getDegrees = async () => {
  const res = await api.listDegree({ limit: 10, offset:0});
  setDegrees(res);
 };

 const alertMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}




 

const checkCaseSaveError = (field: string) => {
  switch(field){
    case 'Teacher_id':
      alertMessage("error","รหัสอาจารย์ขึ้นต้นด้วย T ตามด้วยตัวเลข 7 หลัก");
      return;
    case 'Course_name':
       alertMessage("error","ชื่อหลักสูตรต้องเป็นตัวอักษร A-Z หรือ a-z หรือ 0-9");
       return; 
    case 'Course_year':
         alertMessage("error","ปีของหลักสูตรต้องเป็นตัวเลข จำนวนเต็มบวก");
         return;  
     default:
       alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
       return;
  }
}




 
 useEffect(() => { //เรียกใช้
    getDepartments()
    getSubjects()
    getDegrees()
 }, []
 );

 // clear input form
 function clear() {
  setCurrency({});
}

//6 กดบันทึกที่หน้า ui 
 function save() {
    const apiUrl = 'http://localhost:8080/api/v1/courses'; //url จาก main.go + activitys ตรง crud ใน ctl
    const requestOptions = { //ส่งคำขอที่เป็นแบบ ใส่ค่าข้อมูล
      method: 'POST', // 7 เพิ่มเข้ารายการกิจกรรม (ใส่ค่าข้อมูล)
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(currency),
  };
    console.log(currency); // ดูข้อมูล กด F12 เลือก Tab Console
    fetch(apiUrl, requestOptions) //ส่งคำขอผ่าน apiUrl ให้ ctl ส่งค่ากลับมา
      .then(response => response.json())
      .then(data => {//ส่งค่า data = ac
        console.log(data);
        if (data.status === true) { //เชื่อมกับ controllerหลัก ตรง create
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  }
 return (
   <Page theme={pageTheme.home}>

    <div>
     
     
        <MuiThemeProvider theme={theme}>
          <AppBar position="static" color="primary">
           <Toolbar>
          
              <MenuBookIcon className={classes.text7} style={{ fontSize: 100}} ></MenuBookIcon>
              <FormControl className={classes.text4} > <text>Course</text>
              
              
              </FormControl>
              
            
           </Toolbar>
          </AppBar>
        </MuiThemeProvider>
      
      
    </div>
    

     <Content>

     <div>
       <FormControl className={classes.text1} > <text>สาขาวิชา</text> </FormControl>   
       <FormControl fullWidth  variant="outlined" 
        className={classes.box3}>
        <InputLabel>เลือก</InputLabel>
        <Select
          name="DepartmentID" //เก็บไว้ตัวแปรไหน
          value={currency.DepartmentID || ''}  //เก็บไว้ที่ไหน
          onChange={handleChange}
        >
        {departments.map(item => { //ฟังชันไว้แสดงข้อมูลใน  user combo box
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.dEPARTMENT}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </div> 




     <div>
       <FormControl className={classes.text0} > <text>รายวิชา</text> </FormControl>   
       <FormControl fullWidth  variant="outlined" 
        className={classes.box1}>
        <InputLabel>เลือก</InputLabel>
        <Select
          name="SubjectID"
          value={currency.SubjectID || ''}
          onChange={handleChange}
        >
        {subjects.map(item => {
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.subjectName}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </div> 



       <div>
       <FormControl className={classes.text3} ><text >ระดับการศึกษา</text> </FormControl>
       <FormControl fullWidth  variant="outlined" 
        className={classes.box2}>
        <InputLabel>เลือก</InputLabel>
        <Select
          name="DegreeID"
          value={currency.DegreeID || ''}
          onChange={handleChange}
        >
        {degrees.map(item => {
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.degreeName}
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </div>    

      <div>
      <FormControl className={classes.text8} >  <text >ชื่อหลักสูตร</text> </FormControl>
       <FormControl className={classes.textfill1} >
           <TextField id="Course_name"  variant="outlined"        
           name="CourseName"
           type="string"
           defaultValue="พิมพ์ชื่อหลักสูตรด้วยอักษร A-Z หรือ a-z ต่อด้วยปี พศ. "
           value = {currency.CourseName}
           onChange={handleChange}
           />
      </FormControl>
      </div>

      <div>
      <FormControl className={classes.text9} >  <text >รหัสอาจารย์</text> </FormControl>
       <FormControl className={classes.textfill2} >
           <TextField id="Teacher_id"  variant="outlined" 
           name="TeacherID"
           type="string"
           defaultValue="พิมพ์รหัสอาจารย์ขึ้นต้นด้วย T ตามด้วยตัวเลข 7 หลัก"
           value = {currency.TeacherID}
           onChange={handleChange}
           />
      </FormControl>
      </div>

      <div>
      <FormControl className={classes.text10} >  <text >ปีหลักสูตร</text> </FormControl>
       <FormControl className={classes.textfill3} >
           <TextField id="Course_year"  variant="outlined" 
           name="CourseYear"
           type="number"
           defaultValue="พิมพ์ปีของหลักสูตร เป็นเลขจำนวนเต็มบวก"
           value = {currency.CourseYear}
           onChange={handleChangeInt}
           />
      </FormControl>
      </div>

      <div className={classes.button}>
      <Button size="large" variant="contained" color="secondary" onClick = {save}> บันทึกหลักสูตร </Button>
      </div>
     </Content>
   </Page>
 );
};
 
export default Course;

