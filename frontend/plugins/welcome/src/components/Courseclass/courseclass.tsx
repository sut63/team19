import React, { FC,useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';
import { makeStyles, } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { EntClassdate } from '../../api/models/EntClassdate'; // import interface Classdate
import { EntClasstime } from '../../api/models/EntClasstime'; // import interface Classtime
import { EntClassroom } from '../../api/models/EntClassroom'; // import interface Classroom
import { EntSubject } from '../../api/models/EntSubject'; // import interface Subjects

import { EntInstructorInfo } from '../../api/models/EntInstructorInfo'; // import interface InstructorInfo
import Swal from 'sweetalert2'; // alert

const useStyles = makeStyles(theme => ({
   
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

  

interface  Courseclass  {
  tablecode : string;
  day :    number;
	time    :      number;
	instructor  :  number;
	subject :  number;
	room    :  number;
  groupclass : string;
  annotation : string;
}

const Courseclass: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [tablecodeError, setTablecodeError] = React.useState('');
  const [groupclassError, setGroupclassError] = React.useState('');
  const [annotationError, setAnnotationError] = React.useState('');
  const [courseclass, setCourseclass] = React.useState<Partial<Courseclass>>({});
  const [subject, setSubject] = React.useState<EntSubject[]>([]);
  const [instructor, setInstructor] = React.useState<EntInstructorInfo[]>([]);
  const [classdate, setClassdate] = React.useState<EntClassdate[]>([]);
  const [classtime, setClasstime] = React.useState<EntClasstime[]>([]);
  const [classroom, setClassroom] = React.useState<EntClassroom[]>([]);
  
    // alert setting
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });

  const getSubject = async () => {
    const res = await http.listSubject({ limit: 10, offset: 0 });
    setSubject(res);
  };

  const getInstructor = async () => {
    const res = await http.listInstructorinfo({ limit: 10, offset: 0 });
    setInstructor(res);
  };

  const getClassdate = async () => {
    const res = await http.listClassdate({ limit: 10, offset: 0 });
    setClassdate(res);
  };
  const getClasstime = async () => {
    const res = await http.listClasstime({ limit: 10, offset: 0 });
    setClasstime(res);
  };
  const getClassroom = async () => {
    const res = await http.listClassroom({ limit: 10, offset: 0 });
    setClassroom(res);
  };

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  // Lifecycle Hooks
  useEffect(() => {
    getSubject();
    getInstructor();
    getClassdate();
    getClasstime();
    getClassroom();
  }, []);

  // set data to object so
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof Courseclass;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setCourseclass({ ...courseclass, [name]: value });
    console.log(courseclass);
  };

   // Function for validate Tablecode
   const validateTablecode = (val: string) => {
    return val.match("[T]\\d{2}$");
  }

  // Function for validate Groupclass
  const validateGroupclass = (val: string) => {
    return val.match("[G]\\d{1}$"); 
  }

  // Function for validate Annotation
  const validateAnnotation = (val: string) => {
    return val.match("[^\s]");
  }

  const checkPattern  = (name: string, value: string) => {
    switch(name) {
      case 'tablecode':
        validateTablecode(value) ? setTablecodeError('') : setTablecodeError('Tablecode must begin with T and limit 2 digits');
        return;
      case 'groupclass':
        validateGroupclass(value) ? setGroupclassError('') : setGroupclassError('GroupClass must begin with G follow with 1 digit');
        return;
      case 'annotation':
        validateAnnotation(value) ? setAnnotationError('') : setAnnotationError('Your Annotation field is Empty')
        return;
      default:
        return;
    }
  }

  const checkCaseSaveError = (field: string) => {
    switch(field){
      case "tablecode":
        alertMessage("error","รูปแบบรหัสตารางไม่ถูกต้อง ต้องขึ้นต้นด้วย T และตามด้วยเลขสองหลัก");
        return;
      case "GroupClass":
        alertMessage("error","รูปแบบกลุ่มไม่ถูกต้อง ต้องขึ้นต้นด้วย G และตามด้วยเลขหนึ่งหลัก");
        return;
      case "Annotation":
        alertMessage("error","คำอธิบายห้ามมีการว่างไว้");
        return;
       default:
         alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
         return;
    }
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/courseclasss';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(courseclass),
    };

    console.log(courseclass); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        }else {
          checkCaseSaveError(data.error.Name)

        }
      });
  }
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบสร้างแผนการสอนอาจารย์`}
       subtitle=""
     >
     </Header>
     <Content>
       <ContentHeader title="บันทึกข้อมูล"></ContentHeader>
       <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รหัสตาราง</div>
            </Grid>
            <Grid item xs={9}>
            <form className={classes.root} noValidate autoComplete="off">
      <TextField  label="กรอกรหัสตาราง" name ="tablecode" type="string" error = {tablecodeError ? true : false} helperText = {tablecodeError}
      value={courseclass.tablecode } onChange={handleChange} className={classes.textField}/>
    </form>      
      </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รายชื่อ</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกอาจารย์</InputLabel>
        <Select
          name="instructor"
          type="number"
          value={courseclass.instructor || ''}
          onChange={handleChange}
        >
          {instructor.map(item => {
                    if (item.nAME == JSON.parse(String(localStorage.getItem("Name"))) ) {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.nAME}
                        </MenuItem>
                      );
                    }})}
        </Select>
      </FormControl>
      </Grid>
            

            <Grid item xs={3}>
              <div className={classes.paper}>เลือกรายวิชา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกรายวิชา</InputLabel>
        <Select
          name="subject"
          type="number"
          value={courseclass.subject || ''}
          onChange={handleChange}
        >
          {subject.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.subjectName}
                      </MenuItem>
                    );
                  })}
        </Select>
        
      </FormControl>
      </Grid>

      <Grid item xs={3}>
              <div className={classes.paper}>เลือกห้องเรียน</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกห้องเรียน</InputLabel>
        <Select
          name="room"
          type="number"
          value={courseclass.room || ''}
          onChange={handleChange}
        >
          {classroom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.rOOM}
                      </MenuItem>
                    );
                  })}
        </Select>
        
      </FormControl>
      </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลือกวันสอน</div>
            </Grid>
            <Grid item xs={9}>
      <div>
      <FormControl fullWidth  variant="outlined"
        className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">จ.-อา.</InputLabel>
        <Select
          name="day"
          type="number"
          value={courseclass.day || ''}
          onChange={handleChange}
        >
          {classdate.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.dAY}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </div>
      </Grid>

      <Grid item xs={3}>
              <div className={classes.paper}>เลือกเวลาสอน</div>
            </Grid>
            <Grid item xs={9}>
      <div>
      <FormControl fullWidth  variant="outlined"
        className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">เวลาสอน</InputLabel>
        <Select
          name="time"
          type="number"
          value={courseclass.time || ''}
          onChange={handleChange}
        >
          {classtime.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tIME}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </div>
      </Grid>

      <Grid item xs={3}>
              <div className={classes.paper}>กลุ่มเรียน</div>
            </Grid>
            <Grid item xs={9}>
            <form className={classes.root} noValidate autoComplete="off">
      <TextField  label="กลุ่ม" name ="groupclass" type="string" error = {groupclassError ? true : false} helperText = {groupclassError}
      value={courseclass.groupclass } onChange={handleChange} className={classes.textField}/>
    </form>      
      </Grid>

      <Grid item xs={3}>
              <div className={classes.paper}>ใส่คำอธิบายเพิ่มเติม</div>
            </Grid>
            <Grid item xs={9}>
            <form className={classes.root} noValidate autoComplete="off">
      <TextField  label="คำอธิบาย" name ="annotation" type="string" error = {annotationError ? true : false} helperText = {annotationError}
      value={courseclass.annotation } onChange={handleChange} className={classes.textField}/>
    </form>      
      </Grid>
      
          <Grid item xs={3}></Grid>
          <Grid item xs={9}>
             <Button
               onClick={save}
               startIcon={<SaveIcon />}
               variant="contained"
               color="primary"
             >
               บันทึกข้อมูล
             </Button>
           
           </Grid>
          </Grid>
        </Container>
        </Content>
   </Page>
 );
}
export default Courseclass;