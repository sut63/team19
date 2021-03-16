import React, { FC,useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Swal from 'sweetalert2';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import VpnKeyOutlinedIcon from '@material-ui/icons/VpnKeyOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi } from 'plugin-welcome/src/api/apis';
import { EntInstructorInfo } from 'plugin-welcome/src/api/models/EntInstructorInfo';

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: '#1de9b6',
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

interface login {
  email: string;
  password: string;
}

const Login : FC = ({ setSession })  => {
  const classes = useStyles();
  const [to, setTo] = React.useState('');

  const [login, setLogin] = React.useState<Partial<login>>({});
  const [alertse, setAlertse] = useState(0)
  const [alertsp, setAlertsp] = useState(0)

  // Lifecycle Hooks
  useEffect(() => {
    CheckReset();
  }, []);

  function CheckReset(){
    const namereset = JSON.parse(String(localStorage.getItem("Name")));
    if(namereset != null){
      setTo("/welcome")
      Linklogin();
    }
  }

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

  const handleChange = (event: React.ChangeEvent<{ name?: string; value: any }>,) =>{
    const name = event.target.name as keyof typeof Login;
    const { value } = event.target;
    setLogin({ ...login, [name]: value });
    if(event.target.name == 'email'){
      setAlertse(event.target.value.length)
    }else{
      setAlertsp(event.target.value.length)
    }
  }

  const Login = async () => {
    if (alertse != 0 && alertsp != 0) {
      const apiUrl = `http://localhost:8080/api/v1/logins?email=${login.email}&password=${login.password}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => { 
        if (data.data != "") {
          data.data.map((item : any ) => {
          if(item.EMAIL == login.email){
            localStorage.setItem("Name", JSON.stringify(item.NAME))
            localStorage.setItem("Title", JSON.stringify(item.edges?.Title?.TITLE))
            setTo("/welcome")
            Linklogin();
            setAlertse(0)
            setAlertsp(0)
          }else {
            clear();
            setTo("")
            Toast.fire({
              icon: 'error',
              title: 'Incorrect Email or Password.',
            })
          }})
        } else {
          clear();
          setTo("")
          Toast.fire({
            icon: 'error',
            title: 'Incorrect Email or Password.',
          })
        }
      });
    }else if(login.email == undefined && login.password != undefined) {
      Toast.fire({
        icon: 'error',
        title: 'Please fill in Email.',
      })
    }else if(alertse == 0 && alertsp > 0) {
      Toast.fire({
        icon: 'error',
        title: 'Please fill in Email.',
      })
    }else if(login.password == undefined && login.email != undefined){
      Toast.fire({
        icon: 'error',
        title: 'Please fill in Password.',
      }) 
    }else if(alertsp == 0 && alertse > 0){
      Toast.fire({
        icon: 'error',
        title: 'Please fill in Password.',
      }) 
    }else{
      Toast.fire({
        icon: 'error',
        title: 'Please fill out this form.',
      }) 
    }
  }

  function clear() {
    setLogin({});
    setAlertse(0)
    setAlertsp(0)
    }

  function Linklogin(){
    setSession ({
      isLoggedIn : true,
      isSignIn : true
    });
  }

  function LinkSignIn(){
    setSession ({
      isSignIn : true
    });
  }

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <VpnKeyOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Login
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            value = {login.email || ''}
            autoComplete="email"
            autoFocus 
            onChange = {handleChange}
            />
          {<TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            value = {login.password || ''}
            onChange = {handleChange}
          autoComplete="current-password" /> }
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            component={RouterLink}
            to={to}
            onClick ={Login}
          >
            Login
          </Button>
          <Grid container>
            <Grid item>
              <Link 
               align="center"
               href="#" 
               variant="body2"
               component={RouterLink}
               to="/Sign-in"
               onClick = {LinkSignIn}
               >
              {"Don't have an account? Sign-In"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={7}>
        <Copyright />
      </Box>
    </Container>
  );
}

export default Login;