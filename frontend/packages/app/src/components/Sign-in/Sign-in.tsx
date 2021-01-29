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
    const [to, setTo] = React.useState("");

    const [instructors, setInstructor] = React.useState<Partial<instructor>>({});
    const [phoneError, setPhoneError] = React.useState('');
    const [nameError, setNameError] = React.useState('');
    const [emailError, setEmailError] = React.useState('');
    const [passwordError, setPasswordError] = React.useState('');

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

  // set data to object instructors
  const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof Sign;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setInstructor({ ...instructors, [name]: value });
    console.log(instructors);
    };

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

  // Function for validate email
  const validateEmail = (email: string) => {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
  }

  // Function for validate name
  const validateName = (val: string) => {
    return val.match("[A-Z]");
  }

  // Function for validate phonenumber
  const validatePhone = (val: string) => {
    return val.match("[0]+[9]\\d{8}$|[0]+[8]\\d{8}$|[0]+[6]\\d{8}$"); // \\d[0-9]{6,7}$ for need 9-10 digits
  }

  // Function for validate password
  const validatePassword = (val: string) => {
    return val.match("[^\s]");
  }

  // Check Pattern of Input Data
  const checkPattern  = (name: string, value: string) => {
    switch(name) {
      case 'phonenumber':
        validatePhone(value) ? setPhoneError('') : setPhoneError('Your Phone-Number must begin with 09 08 or 06 and limit 10 digits');
        return;
      case 'name':
        validateName(value) ? setNameError('') : setNameError('Your name must begin with Uppercase');
        return;
      case 'email':
        validateEmail(value) ? setEmailError('') : setEmailError('Invalid Pattern of Email')
        return;
      case 'password':
        validatePassword(value) ? setPasswordError('') : setPasswordError('Your Password is Empty')
        return;
      default:
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'PHONENUMBER':
        alertMessage("error","Your Phone-Number must begin with 09 08 or 06 and limit 10 digits");
        return;
      case 'NAME':
        alertMessage("error","Your name must begin with Uppercase");
        return;
      case 'EMAIL':
        alertMessage("error","Invalid Pattern of Email");
        return;
      case 'PASSWORD':
          alertMessage("error","Your Password is Empty");
          return;
      default:
        alertMessage("error","Save Error");
        return;
    }
  }

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
          LinkIn();
          setTo("/Login")
          Toast.fire({
            icon: 'success',
            title: 'Save Success',
          });
        } else {
          checkCaseSaveError(data.error.Name)
          setTo("/")
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
                      error = {nameError ? true : false}
                      helperText= {nameError}
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
                      error = {phoneError ? true : false}
                      helperText= {phoneError}
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
                      error = {emailError ? true : false}
                      helperText= {emailError}
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
                      error = {passwordError ? true : false}
                      helperText= {passwordError}
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
                to={to}
                startIcon={<PersonIcon />}
                onClick={() =>{
                  save();
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