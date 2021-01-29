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
  const [searchstext, setSearchstext] = React.useState(String)
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

  //instructors
  const [instructors, setInstructor] = React.useState<EntInstructorInfo[]>([])
  const getInstructor = async () => {
   const res = await api.listInstructorinfo({limit: 100, offset: 0 });
   setInstructor(res);
  };

  //instructorsshow
  const [instructorsshow, setInstructorshow] = React.useState<EntInstructorInfo[]>([])
  const getInstructorshow = async () => {
   const res = await api.listInstructorinfo({limit: 100, offset: 0 });
   setInstructorshow(res);
  };
    
  // Lifecycle Hooks
  useEffect(() => {
    getInstructor();
    getInstructorshow();
  }, []);

  function clearshow(){
    setInstructorshow([])
  }

  var alerts : number

  function Searchs() {
    instructors.map((item: any) => {  
      if (item.nAME == searchstext ) {
    const getInstructorsearch = async () => {
      const res = await api.getInstructorinfo({id: item.id})
      setInstructorsearch(res);
      alerts = res.length
      if (alerts > 0) {
        Toast.fire({
          icon: 'success',
          title: 'Search Success',
        })
      } else {
        Toast.fire({
          icon: 'error',
          title: 'Search Error',
        })
      }
     };
    getInstructorsearch();
    }
    else {
      setInstructorsearch([]);
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    }
  });
  }

  const handleChange = (event: React.ChangeEvent<{ value: unknown }>,) =>{
    setSearchstext(event.target.value as string);
    console.log(searchstext);
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
          {instructorsearch.map((item : EntInstructorInfo) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.edges?.title?.tITLE}</TableCell>
              <TableCell align="center">{item.nAME}</TableCell>
              <TableCell align="center">{item.eMAIL}</TableCell>
              <TableCell align="center">{item.edges?.instructorroom?.rOOM}</TableCell>
              <TableCell align="center">{item.edges?.department?.dEPARTMENT}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </Content>
  );
}
export default SearchInstructor;