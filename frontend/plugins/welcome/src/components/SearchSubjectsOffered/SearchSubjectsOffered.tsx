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
import { EntSubjectsOffered } from '../../api/models/EntSubjectsOffered'; // import interface SubjectsOffered
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

const SearchSubjectsOffered: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [searchstext, setSearchstext] = React.useState(String)
  const [subjectsoffered, setSubjectsOffered] = React.useState<EntSubjectsOffered[]>([]);
  const [subjectsofferedsearch, setSubjectsOfferedsearch] = React.useState<EntSubjectsOffered[]>([]);
  

  useEffect(() => {
    getSubjectsOffered();
  }, []);

  
  const getSubjectsOffered = async () => {
    const res = await api.listSubjectsOffered({ limit: 10 });
    setSubjectsOffered(res); 
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
  //var alerts : number;

   const handleChange = (event: React.ChangeEvent<{ value: unknown }>,) =>{
    setSearchstext(event.target.value as string);
    console.log(searchstext);
  }

  function search(){
    if (searchstext != "") {
      const apiUrl = `http://localhost:8080/api/v1/searchSubjectsOffereds?name=${searchstext}`;
      const requestOptions = {
        method: 'GET',
      };
      console.log(searchstext);
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data);
        //setSubjectsOfferedsearch(data.data);
        data.data.map((item : any ) => {        
        if (data.data != "" && item.edges?.Subject?.Subject_name == searchstext) {
          Toast.fire({
            icon: 'success',
            title: 'Search Success',
          })
          setSubjectsOfferedsearch(data.data);
        }else if(data.data != "" && item.edges?.Subject?.Subjec_name != searchstext){
          Toast.fire({
            icon: 'error',
            title: 'Search Error',
          })
          setSubjectsOfferedsearch([]);
        }})
        if (data.data == ""){
          setSubjectsOfferedsearch([]);
          Toast.fire({
            icon: 'error',
            title: 'Search Error',
          })
        }
      });
    }else {
      setSubjectsOfferedsearch([]);
      Toast.fire({
        icon: 'error',
        title: 'Search Error',
      })
    } 
  }
  return (
    <Page theme={pageTheme.home}>
    <Header
      title={`ระบบค้นหารายวิชาที่เปิดสอน`}
      subtitle=""
    >
    </Header>
    <Content>
    <Grid container spacing={1}>  
    <Grid item xs={2}>
                  <div className={classes.paper}> ค้นหารายวิชาที่เปิดสอน </div>
                </Grid>
                <Grid item xs={3}>
                  <form className={classes.container}>
                    <TextField
                      label="รายวิชาที่เปิดสอน"
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
            <TableCell align="center">ID</TableCell>
            <TableCell align="center">Subject</TableCell>
            <TableCell align="center">Degree</TableCell>
            <TableCell align="center">Year</TableCell>
            <TableCell align="center">Term</TableCell>
            <TableCell align="center">Amount</TableCell>
            <TableCell align="center">Remain</TableCell>
            
          </TableRow>
        </TableHead>
        <TableBody>
          { subjectsofferedsearch.map((item : any ) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.edges?.Subject?.Subject_name}</TableCell>
              <TableCell align="center">{item.edges?.Degree?.Degree_name}</TableCell>
              <TableCell align="center">{item.edges?.Year?.YEAR}</TableCell>
              <TableCell align="center">{item.edges?.Term?.TERM}</TableCell>
              <TableCell align="center">{item.AMOUNT}</TableCell>
              <TableCell align="center">{item.Remain}</TableCell>
              
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
    </Content>
   </Page>
  );
}
export default SearchSubjectsOffered;