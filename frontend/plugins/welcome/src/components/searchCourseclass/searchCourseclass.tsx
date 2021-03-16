import React, { useState, useEffect, FC } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
//import DeleteIcon from '@material-ui/icons/Delete';
import SearchIcon from '@material-ui/icons/Search';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntCourseclass } from '../../api/models/EntCourseclass'; // import interface Courseclass
import { Content, Header, Page, pageTheme } from '@backstage/core';
import {  Grid, TextField } from '@material-ui/core';
import Swal from 'sweetalert2';
  const useStyles = makeStyles((theme: Theme) =>

  createStyles({

    table: {
      minWidth: 650,
    },

  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(2),
    fontSize: 20,
  },

  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
   
  textField: {
    width: 2000,
  },

}),
);

const SearchCourseClass: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [searchstext, setSearchstext] = React.useState(String)
  const [courseclass, setCourseclass] = useState<EntCourseclass[]>([]);
  const [courseclasssearch, setCourseClasssearch] = useState<EntCourseclass[]>([]);
  

  useEffect(() => {
    getCourseClass();
  }, []);

  
  const getCourseClass = async () => {
    const res = await api.listCourseclass({ limit: 100 });
    setCourseclass(res); 
  };

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
  var alerts : number;

   const handleChange = (event: React.ChangeEvent<{ value: unknown }>,) =>{
    setSearchstext(event.target.value as string);
    console.log(searchstext);
  }

  const Searchs = async () => {
    if (searchstext != "") {
      const apiUrl = `http://localhost:8080/api/v1/searchcourseclasss?tablecode=${searchstext}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        data.data.map((item : any ) => {        
        if (data.data != "" && item.tablecode == searchstext) {
          setCourseClasssearch(data.data);
          Toast.fire({
            icon: 'success',
            title: 'Search Success',
          })
        }else if(data.data != "" && item.tablecode != searchstext){
          setCourseClasssearch([]);
          Toast.fire({
            icon: 'error',
            title: 'Search Error',
          })
        }})
        if (data.data == ""){
          setCourseClasssearch([]);
          Toast.fire({
            icon: 'error',
            title: 'Search Error',
          })
        }
      });
    }else {
      setCourseClasssearch([]);
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    } 
  }

  return (
    <Page theme={pageTheme.home}>

    <Content>
    <Grid container spacing={1}>  
    <Grid item xs={2}>
                  <div className={classes.paper}> ค้นหารหัสตารางสอน </div>
                </Grid>
                <Grid item xs={3}>
                  <form className={classes.container}>
                    <TextField
                      label="รหัสตารางสอน"
                      name="name"
                      type="string"
                      value={searchstext || ''} 
                      className={classes.textField}
                      onChange={handleChange}
                    />
                  </form>
                </Grid>

                    <Grid item xs={2}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            startIcon={<SearchIcon/>}
                            onClick={() =>{
                                  Searchs();
                                }}
                        >    Search
                        </Button>
                    </Grid>
  </Grid>
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">Table ID</TableCell>
            <TableCell align="center">Instructor</TableCell>
            <TableCell align="center">Subject</TableCell>
            <TableCell align="center">Group</TableCell>
            <TableCell align="center">ClassRoom</TableCell>
            <TableCell align="center">Day</TableCell>
            <TableCell align="center">time</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {courseclasssearch.map((item : any) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.tablecode}</TableCell>
              <TableCell align="center">{item.edges?.InstructorInfo?.NAME}</TableCell>
              <TableCell align="center">{item.edges?.Subject?.Subject_name}</TableCell>
              <TableCell align="center">{item.GroupClass}</TableCell>
              <TableCell align="center">{item.edges?.Classroom?.ROOM}</TableCell>
              <TableCell align="center">{item.edges?.Classdate?.DAY}</TableCell>
              <TableCell align="center">{item.edges?.Classtime?.TIME}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </Content>
   </Page>
  );
}
export default SearchCourseClass;