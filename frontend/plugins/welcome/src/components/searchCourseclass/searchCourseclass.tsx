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

  function search(){
    courseclass.map((item: any) => {  
      if (item.tablecode == searchstext ) {
    const getCourseClasssearch = async () => {
      const res = await api.getCourseclass({id: item.id})
      setCourseClasssearch(res);
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
     getCourseClasssearch();
    }
    else {
      setCourseClasssearch([]);
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    }
  });
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
                                  search();
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
            <TableCell align="center">ClassRoom</TableCell>
            <TableCell align="center">Day</TableCell>
            <TableCell align="center">time</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {courseclasssearch === undefined
            ? null
            : courseclasssearch.map((item) => (
            <TableRow key={item.id}>
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
}
export default SearchCourseClass;