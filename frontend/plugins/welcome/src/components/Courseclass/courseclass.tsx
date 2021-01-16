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
  
  day :    number;
	time    :      number;
	instructor  :  number;
	subject :  number;
	room    :  number;
}

const Courseclass: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

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
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Courseclass;
    const { value } = event.target;
    setCourseclass({ ...courseclass, [name]: value });
    console.log(courseclass);
  };
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
     <Header
       title={`ระบบสร้างแผนการสอนอาจารย์`}
       subtitle=""
     >
       <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"color="primary"
             >
               ออกจากระบบ
       </Button>
     </Header>
     <Content>
       <ContentHeader title="บันทึกข้อมูล"></ContentHeader>
       <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
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
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nAME}
                      </MenuItem>
                    );
                  })}
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