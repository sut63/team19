import React, { useState, useEffect, FC } from 'react';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
//import DoneOutlineIcon from '@material-ui/icons/DoneOutline';
import SearchIcon from '@material-ui/icons/Search';
import Swal from 'sweetalert2';
import {
    Grid,
    TextField,
} from '@material-ui/core';
import { DefaultApi } from '../../api';
import { EntInstructorInfo } from '../../api/models/EntInstructorInfo'; // import interface InstructorInfo
const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(2),
  },
  textField: {
    width: 300,
  },
  formControl: {
    width: 300,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  }));

const SearchInstructor: FC<{}> = ()  => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [searchstext, setSearchstext] = useState(String)
  const [instructorsearch, setInstructorsearch] = React.useState<EntInstructorInfo[]>([])

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

  //instructorsshow
  const [instructorsshow, setInstructorshow] = React.useState<EntInstructorInfo[]>([])
  const getInstructorshow = async () => {
   const res = await api.listInstructorinfo({limit: 100, offset: 0 });
   setInstructorshow(res);
  };
    
  // Lifecycle Hooks
  useEffect(() => {
    getInstructorshow();
  }, []);

  function clearshow(){
    setInstructorshow([])
  }

  const Searchs = async () => {
    if (searchstext != "") {
      const apiUrl = `http://localhost:8080/api/v1/searchinstructorinfos?name=${searchstext}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        data.data.map((item : any ) => {        
        if (data.data != "" && item.NAME == searchstext) {
          setInstructorsearch(data.data);
          Toast.fire({
            icon: 'success',
            title: 'Search Success',
          })
        }else if(data.data != "" && item.NAME != searchstext){
          setInstructorsearch([]);
          Toast.fire({
            icon: 'error',
            title: 'Search Error',
          })
        }})
        if (data.data == ""){
          setInstructorsearch([]);
          Toast.fire({
            icon: 'error',
            title: 'Search Error',
          })
        }
      });
    }else {
      setInstructorsearch([]);
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    } 
  }

  const handleChange = (event: React.ChangeEvent<{ value: any }>,) =>{
    setSearchstext(event.target.value as string);
  }

  return (
    <Content>
    <Grid container spacing={1}>
                    <Grid item xs={1}>
                  <div className={classes.paper}>Search Intructor : </div>
                </Grid>
                <Grid item xs={2}>
                  <form className={classes.container}>
                    <TextField
                      label="Name Intructor"
                      name="name"
                      type="string"
                      value={searchstext || ''} 
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
                </Grid>

                    <Grid item xs={1}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            startIcon={<SearchIcon/>}
                            onClick={() =>{
                                  Searchs();
                                  clearshow()
                                }}
                        >    Search
                        </Button>
                    </Grid>
      </Grid>
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">Title</TableCell> 
            <TableCell align="center">Name</TableCell>
            <TableCell align="center">Email</TableCell>
            <TableCell align="center">Instructor Room</TableCell>
            <TableCell align="center">Department</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {instructorsshow.map((item : EntInstructorInfo) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.edges?.title?.tITLE}</TableCell>
              <TableCell align="center">{item.nAME}</TableCell>
              <TableCell align="center">{item.eMAIL}</TableCell>
              <TableCell align="center">{item.edges?.instructorroom?.rOOM}</TableCell>
              <TableCell align="center">{item.edges?.department?.dEPARTMENT}</TableCell>
            </TableRow>
          ))}
        </TableBody>
        <TableBody>
          {instructorsearch.map((item : any) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.edges?.Title?.TITLE}</TableCell>
              <TableCell align="center">{item.NAME}</TableCell>
              <TableCell align="center">{item.EMAIL}</TableCell>
              <TableCell align="center">{item.edges?.Instructorroom?.ROOM}</TableCell>
              <TableCell align="center">{item.edges?.Department?.DEPARTMENT}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </Content>
  );
}
export default SearchInstructor;