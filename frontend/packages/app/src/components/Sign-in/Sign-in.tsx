import React, { FC, useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import PersonIcon from '@material-ui/icons/Person';
import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';
import { Content, Header, Page, pageTheme, ContentHeader } from '@backstage/core';
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
import Swal from 'sweetalert2';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from 'plugin-welcome/src/api/apis';
import { EntDepartment } from 'plugin-welcome/src/api/models/EntDepartment';
import { EntInstructorRoom } from 'plugin-welcome/src/api/models/EntInstructorRoom';
import { EntTitle } from 'plugin-welcome/src/api/models/EntTitle';

const HeaderCustom = {
    minHeight: '50px',
  };
  
  // css style
  const useStyles = makeStyles(theme => ({
    root: {
      flexGrow: 1,
    },
    margin: {
      margin: theme.spacing(3),
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

  interface instructor {
    name : string;
    phonenumber  : string;
    email: string;
    password: string;
    department: number;
    instructorroom: number;
    title: number;
    
  }

  const Sign: FC = ({ setSession })  => {
    const classes = useStyles();
    const api = new DefaultApi();
    
    const [instructors, setInstructor] = React.useState<Partial<instructor>>({});

    const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
    const [instructorrooms, setInstructorRooms] = React.useState<EntInstructorRoom[]>([]);
    const [titles, setTitles] = React.useState<EntTitle[]>([]);

    const getDepartments = async () => {
        const res = await api.listDepartment({ limit: 100, offset: 0 });
        setDepartments(res);
       };
      
       const getInstructorRooms = async () => {
        const res = await api.listInstructorroom({ limit: 100, offset: 0 });
        setInstructorRooms(res);
       };
      
       const getTitles = async () => {
        const res = await api.listTitle({ limit: 100, offset: 0 });
        setTitles(res);
       };
      
       // Lifecycle Hooks
       useEffect(() => {
        getDepartments();
        getInstructorRooms();
        getTitles();
      }, []);

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
    
    // set data to object instructors
    const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Sign;
    const { value } = event.target;
    setInstructor({ ...instructors, [name]: value });
    console.log(instructors);
    };

    // clear input form
    function clear() {
    setInstructor({});
    }

    function LinkIn(){
      setSession ({
        isSignIn : false,
        isLoggedIn : true
      });
    }

    // function save data
    function save() {
    const apiUrl = 'http://localhost:8080/api/v1/instructorinfos';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(instructors),
      };
    console.log(instructors); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
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
    };

    return (
      <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Create Intructor Account`}>
      </Header>
      <Content>
        <ContentHeader title="Intructor">
        </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
    
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}> Title : </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Title</InputLabel>
                <Select
                  label = "Title"
                  name = "title"
                  value = {instructors.title || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {titles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tITLE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
                  <div className={classes.paper}>Name : </div>
                </Grid>
                <Grid item xs={9}>
                  <form className={classes.container} noValidate>
                    <TextField
                      label="Name Surname"
                      name="name"
                      type="string"
                      value={instructors.name || ''} // (undefined || '') = ''
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
            </Grid>

            <Grid item xs={3}>
                  <div className={classes.paper}>Phone-Number : </div>
                </Grid>
                <Grid item xs={9}>
                  <form className={classes.container} noValidate>
                    <TextField
                      label="PhoneNumber"
                      name="phonenumber"
                      type="string"
                      value={instructors.phonenumber || ''} // (undefined || '') = ''
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
            </Grid>

            <Grid item xs={3}>
                  <div className={classes.paper}>Email : </div>
                </Grid>
                <Grid item xs={9}>
                  <form className={classes.container} noValidate>
                    <TextField
                      label="Email"
                      name="email"
                      type="string"
                      value={instructors.email || ''} // (undefined || '') = ''
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
            </Grid>

            <Grid item xs={3}>
                  <div className={classes.paper}>Password : </div>
                </Grid>
                <Grid item xs={9}>
                  <form className={classes.container} noValidate>
                    <TextField
                      label="Password"
                      name="password"
                      type="password"
                      value={instructors.password || ''} // (undefined || '') = ''
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
            </Grid> 

            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Department : </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Department</InputLabel>
                <Select
                  label = "Department"
                  name="department"
                  value={instructors.department || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {departments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.dEPARTMENT}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
    
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Room : </div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Room</InputLabel>
                <Select
                  label = "Room"
                  name="instructorroom"
                  value={instructors.instructorroom || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {instructorrooms.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.rOOM}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
    
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
    
            <Button
                variant="contained"
                color="secondary"
                size="large"
                startIcon={<ArrowBackIosIcon />}
                component={RouterLink}
                to="/Login"
                onClick={LinkIn}
              >
                Login
              </Button>
    
              <Button
                className={classes.margin}
                variant="contained"
                color="primary"
                size="large"
                component={RouterLink}
                to="/Login"
                startIcon={<PersonIcon />}
                onClick={() =>{
                  save();
                  LinkIn();
                }}>
                Sign-in
              </Button>
              
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
     );
    };

    export default Sign;