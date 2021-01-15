import React, { FC,useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
//import { Link as RouterLink } from 'react-router-dom';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import VpnKeyOutlinedIcon from '@material-ui/icons/VpnKeyOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { DefaultApi } from 'plugin-welcome/src/api/apis';
import { EntInstructorInfo } from 'plugin-welcome/src/api/models/EntInstructorInfo';
//import { ItemCard } from '@backstage/core';
//import { keys } from '@material-ui/core/styles/createBreakpoints';
//import { type, userInfo } from 'os';

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
  const api = new DefaultApi();

  const [login, setLogin] = React.useState<Partial<login>>({});
  const [instructors, setInstructor] = React.useState<EntInstructorInfo[]>([]);


  const getInstructor = async () => {
    const res = await api.listInstructorinfo({offset: 0 });
    setInstructor(res);
   };

  // Lifecycle Hooks
  useEffect(() => {
    getInstructor();
  }, []);

  const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>,) =>{
    const name = event.target.name as keyof typeof Login;
    const { value } = event.target;
    setLogin({ ...login, [name]: value });
    console.log(login);
  }
  
  const Login = async () => {
    instructors.map((item: any) => {
      if (item.eMAIL == login.email && item.pASSWORD == login.password) {
        Linklogin();
        alert('Success')
      }
      else{
        clear();
        alert('Error')
      }
    });
  };

  function clear() {
    setLogin({});
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
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label="Remember me" />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            component={RouterLink}
            to="welcome"
            onClick = {Login}
          >
            Login
          </Button>
          <Grid container>
            <Grid item xs>
              <Link href="#" variant="body2">
                Forgot password?
              </Link>
            </Grid>
            <Grid item>
              <Link
               href="#" 
               variant="body2"
               component={RouterLink}
               to="/Sign-in"
               onClick = {LinkSignIn}
               >
                {"Don't have an account? Sign Up"} 
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
}

export default Login;