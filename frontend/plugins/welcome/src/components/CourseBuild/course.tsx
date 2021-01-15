import React, { FC, useEffect } from 'react';


import Swal from 'sweetalert2'; // alert

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import {
 Content,
 Page,
 pageTheme,

} from '@backstage/core';


import PersonIcon from '@material-ui/icons/Person';
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
      marginLeft: theme.spacing(32),
      marginRight: theme.spacing(4),
      marginTop: theme.spacing(3),
      fontSize: 20,
    },

    text1: {  //ชื่อชมรม
        marginLeft: theme.spacing(30),
        marginRight: theme.spacing(4),
        marginTop: theme.spacing(20),
        fontSize: 20,
    },
    text2: {
      marginLeft: theme.spacing(4),
      marginTop: theme.spacing(4),
      fontSize: 20,
  },
  text3: {
    marginLeft: theme.spacing(25),
    marginTop: theme.spacing(4),
    fontSize: 20,
},

text4: {
  marginLeft: theme.spacing(20), 
  fontSize: 25,
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
  marginLeft: theme.spacing(90),
  marginTop: theme.spacing(1),
  marginBottom:  theme.spacing(1),
  fontSize: 14,
  color: "#f8bbd0",
},


text8: {
  marginLeft: theme.spacing(28),
  marginTop: theme.spacing(4),
  marginBottom:  theme.spacing(1),
  fontSize: 21,
  color: "white",
},


  box1: {
    marginLeft: theme.spacing(9.5),
    marginTop: theme.spacing(2),
    width: '60ch',  
},

  box2: {
    marginLeft: theme.spacing(13),
    marginTop: theme.spacing(2),
    width: '60ch',  
  },

  box3: {
    marginLeft: theme.spacing(10),
    marginTop: theme.spacing(20),
    width: '60ch',  
},


  textfill1: {  //ชื่อกิจกรรม
    marginLeft: theme.spacing(13),
    marginTop: theme.spacing(2),
    width: '60ch',
  },
  
  
  
  button: {
    marginTop: theme.spacing(6),
    marginLeft: theme.spacing(60),
    display: 'flex',
    flexWrap: 'wrap',
    
  },

  button1: {
    width: '50ch',
    
    
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

const handleChange = ( event: React.ChangeEvent<{ name?: string; value: unknown }>,) =>{
    const name = event.target.name as keyof typeof WelcomePage;
    const { value } = event.target;
    setCurrency({ ...currency, [name] : value });
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
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
 return (
   <Page theme={pageTheme.home}>
     
     <div>
        <MuiThemeProvider theme={theme}>
          <AppBar color="primary" className="app-header">
           <Toolbar>
           
              <FormControl className={classes.text4} > <text>| สร้างหลักสูตร | </text></FormControl>
              <PersonIcon className={classes.text7} style={{ fontSize: 100}} ></PersonIcon>
              <FormControl className={classes.text6} > <text>chanwit</text> </FormControl>
              <FormControl className={classes.text5} > <text>Logout</text> </FormControl>
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
           <TextField id="outlined-basic"  variant="outlined" 
           name="CourseName"
           defaultValue="พิมพ์ชื่อหลักสูตร"
           value = {currency.CourseName}
           onChange={handleChange}
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