import React, { Component, FC,useEffect } from 'react';
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
  FormGroup,
  FormControlLabel,
  Checkbox,
  FormLabel,
  FormHelperText,
} from '@material-ui/core';
import { makeStyles, } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { EntYear } from '../../api/models/EntYear'; // import interfa Year
import { EntTerm } from '../../api/models/EntTerm'; // import interfa Term
import { EntSubject } from '../../api/models/EntSubject'; // import interface Subjects
import { EntDegree } from '../../api/models/EntDegree'; // import interface Degree
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

  

interface  EntSubjectsOffered  {
  Subject :    number;
  Year  :    number;
  Degree  :    number;
  Term :       number;
  AMOUNT :      number;
  Remain :     number;
  STATUS :      boolean;
}

const EntSubjectsOffered: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [subjectsoffered, setSubjectsOffered] = React.useState<Partial<EntSubjectsOffered>>({});
  const [subject, setSubject] = React.useState<EntSubject[]>([]);
  const [degree, setDegree] = React.useState<EntDegree[]>([]);
  const [year, setYear] = React.useState<EntYear[]>([]);
  const [term, setTerm] = React.useState<EntTerm[]>([]);
  const [amountError, setAmountError] = React.useState('');
  const [statusError, setStatusError] = React.useState('');
  const [remainError, setRemainError] = React.useState('');
  const [state, setState] = React.useState({
    Open: false,
  });
  const { Open } = state;
  const error = [Open].filter((v) => v).length !== 1;
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
    const res = await http.listSubject({ limit: 100, offset: 0 });
    setSubject(res);
  };

  const getDegree = async () => {
    const res = await http.listDegree({ limit: 100, offset: 0 });
    setDegree(res);
  };

  const getYear = async () => {
    const res = await http.listYear({ limit: 100, offset: 0 });
    setYear(res);
  };
  const getTerm = async () => {
    const res = await http.listTerm({ limit: 100, offset: 0 });
    setTerm(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getSubject();
    getYear();
    getDegree();
    getTerm();
  }, []);

  // set data to object so
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof EntSubjectsOffered;
    const { value } = event.target;
    setSubjectsOffered({ ...subjectsoffered, [name]: value });
    console.log(subjectsoffered);
  };
  
  const handleChangeInt = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
    const name = event.target.name as keyof typeof EntSubjectsOffered;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setSubjectsOffered({ ...subjectsoffered, [name]: Number(value)});
    console.log(subjectsoffered);
  };

  const handleChangeBool = (event: React.ChangeEvent<{ id?: string; value: any; name: string; checked?: boolean; }>) => {
    const id = event.target.id as keyof typeof EntSubjectsOffered;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setState({ ...state, [event.target.name]: event.target.checked });
    setSubjectsOffered({ ...subjectsoffered, [id]: event.target.checked});
    console.log(subjectsoffered);
  };


  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  
  // ฟังก์ชั่นสำหรับ validate Remain
  const validateRemain = (rm: number) => {
    return rm >= 0 ;
  }

  // ฟังก์ชั่นสำหรับ validate Status
  const validateStatus = (val: boolean) => {
    return val == true;}

  // ฟังก์ชั่นสำหรับ validate Amount
  const validateAmount = (amt: number) => {
    return amt >= 1;
  }

  const checkPattern = (name: string, value: any) => {
    switch(name){
      case 'AMOUNT':
        validateAmount(value) ? setAmountError('') : setAmountError("รูปแบบจำนวนที่รับไม่ถูกต้อง ต้องเป็นจำนวนเต็มบวก");
        return;
      case 'STATUS':
        validateStatus(value) ? setStatusError('') : setStatusError("รูปแบบสถานะไม่ถูกต้อง สถานะคือ Open");
        return;
      case 'Remain':
        validateRemain(value) ? setRemainError('') : setRemainError("รูปแบบคงเหลือไม่ถูกต้อง ต้องเป็นจำนวนเต็มบวก");
        return;  
      default:
        return;
    }
  }
   
  const checkCaseSaveError = (field: string) => {
    switch(field){
      case 'AMOUNT':
        alertMessage("error","รูปแบบจำนวนที่รับไม่ถูกต้อง ต้องเป็นจำนวนเต็มบวก");
        return;
      case 'STATUS':
         alertMessage("error","รูปแบบสถานะไม่ถูกต้อง สถานะคือ Open");
         return;
      case 'Remain':
          alertMessage("error","รูปแบบคงเหลือไม่ถูกต้อง ต้องเป็นจำนวนเต็มบวก");
          return;  
       default:
         alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
         return;
    }
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/SubjectsOffereds';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(subjectsoffered),
    };

    console.log(subjectsoffered); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
    if (subjectsoffered.STATUS === true){
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else  {
          checkCaseSaveError(data.error.Name)
        }
      });
    } else {
      Toast.fire({
        icon: 'error',
        title: 'รูปแบบสถานะไม่ถูกต้อง สถานะคือ Open',
      });
  }}
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบบันทึกรายวิชาที่เปิดสอน`}
       subtitle=""
     >
     </Header>
     <Content>
       <ContentHeader title="บันทึกข้อมูล"></ContentHeader>
       <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รายวิชา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกรายวิชา</InputLabel>
        <Select
          name="Subject"
          type="number"
          value={subjectsoffered.Subject || ''}
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
              <div className={classes.paper}>ปีการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกปีการศึกษา</InputLabel>
        <Select
          name="Year"
          type="number"
          value={subjectsoffered.Year || ''}
          onChange={handleChange}
        >
          {year.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.yEAR}
                      </MenuItem>
                    );
                  })}
        </Select>
        
      </FormControl>
      </Grid>

      <Grid item xs={3}>
              <div className={classes.paper}>ภาคการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกภาคการศึกษา</InputLabel>
        <Select
          name="Term"
          type="number"
          value={subjectsoffered.Term || ''}
          onChange={handleChange}
        >
          {term.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tERM}
                      </MenuItem>
                    );
                  })}
        </Select>
        
      </FormControl>
      </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ระดับการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
      <div>
      <FormControl fullWidth  variant="outlined"
        className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">เลือกระดับการศึกษา</InputLabel>
        <Select
          name="Degree"
          type="number"
          value={subjectsoffered.Degree || ''}
          onChange={handleChange}
        >
          {degree.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.degreeName}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </div>
      </Grid>
      <Grid item xs={3}>
              <div className={classes.paper}>จำนวนที่รับ</div>
            </Grid>
            <Grid item xs={9}>
      <form className={classes.root} noValidate autoComplete="off" >
        <TextField 
        error = {amountError ? true : false}
        label="กรอกจำนวนที่รับ"
        name="AMOUNT" 
        type="number" 
        InputProps={{ inputProps: { min: 0 } }}
        helperText= {amountError}
        value={subjectsoffered.AMOUNT || ''}
        variant="outlined"
        onChange={handleChangeInt} 
        className={classes.textField}
                  />
      </form>
      </Grid>
      <Grid item xs={3}>
              <div className={classes.paper}>คงเหลือ</div>
            </Grid>
            <Grid item xs={9}>
      <form className={classes.root} noValidate autoComplete="off" >
        <TextField 
        error = {remainError ? true : false}
        label="กรอกคงเหลือ"
        name="Remain" 
        type="number" 
        InputProps={{ inputProps: { min: 0 } }}
        helperText= {remainError}
        value={subjectsoffered.Remain || ''}
        variant="outlined"
        onChange={handleChangeInt} 
        className={classes.textField}
                  />
      </form>
      </Grid>
      <Grid item xs={3}>
              <div className={classes.paper}>สถานะ</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl required error={error} component="fieldset" className={classes.formControl}>
        <FormGroup>
          <FormControlLabel
            control={<Checkbox 
              id="STATUS"
              checked={Open} 
              value={subjectsoffered.STATUS ||''}
              onChange={handleChangeBool} 
              name="Open" />}
            label="Open"
          />
        </FormGroup>
        <FormHelperText> {statusError}</FormHelperText>
      </FormControl>
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
export default EntSubjectsOffered;