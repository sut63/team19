import React, { FC, useEffect } from 'react';
import { EntCourse } from '../../api/models/EntCourse'; 

import Swal from 'sweetalert2'; // alert
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';


import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import {
 Content,
 Page,
 pageTheme,

} from '@backstage/core';


import MenuBookIcon from '@material-ui/icons/MenuBook';
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
  Grid,

} from '@material-ui/core';
import Courseclass from '../Courseclass';
import { EntCourseclass } from '../../api';


const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    table: {
      minWidth: 650,
    },
   
    root: {
      '& .MuiTextField-root': {
        marginLeft: theme.spacing(9),
        width: '30ch',
      },
    },


text0: {  
      marginLeft: theme.spacing(0),
      marginRight: theme.spacing(0),
      marginTop: theme.spacing(1),
      fontSize: 20,
    },

text4: {
  marginLeft: theme.spacing(3), 
  marginRight: theme.spacing(124), 
  fontSize: 30,
  color: "white",
},


text7: {
  marginLeft: theme.spacing(10),
  marginTop: theme.spacing(0),
  color: "white",
},


box1: {
  marginLeft: theme.spacing(0),
  marginTop: theme.spacing(0),
  marginBottom: theme.spacing(2),
  width: '40ch',  
},


  button: {
    marginLeft: theme.spacing(4),
  marginTop: theme.spacing(1),
  marginBottom: theme.spacing(2),
    //display: 'flex',
    //flexWrap: 'wrap',
    
  },

  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(2),
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



const FindCourse: FC<{}> = () => {
const tables = useStyles();



const api = new DefaultApi();
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

const handleChange = ( event: React.ChangeEvent<{ value: any }>,) =>{
    //const name = event.target.name as keyof typeof FindCourse;
    const { value } = event.target;    
    setTableid(value);
    console.log(tableid);
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
       alertMessage("error","ชื่อหลักสูตรต้องเป็น a-z หรือ A-Z หรือ 0-9");
       return; 
    case 'Course_year':
         alertMessage("error","ปีของหลักสูตรต้องเป็นตัวเลข จำนวนเต็มบวก");
         return;  
     default:
       alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
       return;
  }
}



const [tablename, setTableName] = React.useState<Partial<EntCourseclass[]>>([]);
const [tableid, setTableid] = React.useState<number>(0);
const [EntCourseclass, setEntCourseclass] = React.useState<EntCourseclass[]>([]);
const [courseclassshow, setTableShow] = React.useState<EntCourseclass[]>([]);

 
 useEffect(() => { //เรียกใช้
  getTable()  
 }, []
 );



 // clear input form
 function clear() {
    setTableName([]);
}

const getTable = async () => {
    const res = await api.listCourseclass({limit: 100, offset:0});
    setEntCourseclass(res);
   };

   var lenreturn = 0
   const getTableshow = async () => {
    const res = await api.getCourseclass({id: tableid})
    setTableShow(res)
    lenreturn = res.length
    if (lenreturn>0){
      Toast.fire({
        icon: 'success',
        title: 'พบรหัสตารางที่ค้นหา',
      });
    }
    else{
      Toast.fire({
        icon: 'error',
        title: 'ไม่พบรหัสตารางที่ค้นหา',
      });
    }
   };



function serch(){
  getTableshow();
}



 return (
   <Page theme={pageTheme.home}>

    
    

     <Content>
  <Grid container spacing={1}>  
     <Grid item xs={2}>
       <FormControl className={tables.text0} > <text>เลือกรหัสตารางสอน</text> </FormControl>   
      </Grid>
      <Grid item xs={3}>
       <FormControl fullWidth  variant="outlined" 
        className={tables.box1}>
        <InputLabel>เลือก</InputLabel>
        <Select
          name = "TableID"
          value={tableid || ''}
          onChange={handleChange}
        >
        {EntCourseclass.map(item => {
            return (
              <MenuItem key ={item.id} value = {item.id}>
                {item.tablecode }
                </MenuItem>
            );
          }  
            )}
          
        </Select>
        </FormControl>
      
      </Grid> 
    
      
     
      <Grid>
     <Button size="large"  id={tables.button}  variant="contained" color="secondary" onClick = {serch}> ค้นหารหัสตาราง </Button>
     </Grid>

  </Grid>

      <TableContainer component={Paper}>
     <Table id={tables.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">id</TableCell>
           <TableCell align="center">tablecode</TableCell>
           <TableCell align="center">name</TableCell>
           <TableCell align="center">subject</TableCell>
           <TableCell align="center">room</TableCell>
           <TableCell align="center">day</TableCell>
           <TableCell align="center">time</TableCell>


         </TableRow>
       </TableHead>       
       <TableBody>
         {courseclassshow.map((item: EntCourseclass ) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>  
             <TableCell align="center">{item.tablecode}</TableCell>     
             <TableCell align="center">{item.edges?.instructorInfo?.nAME}</TableCell>     
             <TableCell align="center">{item.edges?.subject?.subjectName}</TableCell>     
             <TableCell align="center">{item.edges?.classroom?.rOOM}</TableCell>     
             <TableCell align="center">{item.edges?.classdate?.dAY}</TableCell>     
             <TableCell align="center">{item.edges?.classtime?.tIME}</TableCell>   
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>




     </Content>
   </Page>

 );
};

export default FindCourse;
